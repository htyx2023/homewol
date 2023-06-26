package service

import (
	"encoding/hex"
	"fmt"
	"github.com/labstack/echo"
	"home-wol/common"
	"log"
	"net"
	"net/http"
	"strings"
)

var BroadcastIP net.IP

func Wol(c echo.Context) error {
	mac := c.QueryParam("mac")
	authCode := c.QueryParam("auth_code")
	//userName := c.QueryParam("user_name")

	if authCode == "" {
		return c.String(http.StatusInternalServerError, "验证失败")
	}

	log.Println("mac:", mac)
	log.Println("auth_code:", authCode)
	log.Println("secret:", common.Secret)

	auth := common.NewGoogleAuth()

	code, err := auth.GetCode(common.Secret)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	if authCode != code {
		log.Println("验证失败:")
		return c.String(http.StatusInternalServerError, "验证失败")
	}

	if mac == "" {
		return c.String(http.StatusInternalServerError, "mac为空")
	}

	//处理mac地址
	mac = strings.ReplaceAll(mac, ":", "")
	mac = strings.ReplaceAll(mac, "-", "")
	mac = strings.ToUpper(mac)

	hexStr := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "A", "B", "C", "D", "E", "F"}

	for _, str := range mac {
		flag := ArrayIndex(string(str), hexStr)
		if flag == -1 {
			return c.String(http.StatusInternalServerError, "mac格式无效")
		}
	}

	if len(mac) != 12 {
		return c.String(http.StatusInternalServerError, "mac格式无效")
	}

	data := "FFFFFFFFFFFF"

	for i := 0; i < 16; i++ {
		data = fmt.Sprintf("%s%s", data, mac)
	}

	// 将 16进制的字符串 转换 byte
	hexData, _ := hex.DecodeString(data)

	go common.SendWol(BroadcastIP, 9, hexData)
	log.Println("指令发送成功:")
	return c.String(http.StatusOK, "指令发送成功")
}

func ArrayIndex(v string, a []string) int {
	for ai, av := range a {
		if v == av {
			return ai
		}
	}

	return -1
}
