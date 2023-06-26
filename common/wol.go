package common

import (
	"encoding/hex"
	"fmt"
	"log"
	"net"
)

func SendWol(ip net.IP, port int, context []byte) {

	socket, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   ip,
		Port: port,
	})
	if err != nil {
		log.Println("连接UDP失败，err: ", err)
		return
	}
	defer socket.Close()
	log.Println("ip:", ip, "port:", port, ";send:", hex.EncodeToString(context))
	_, err = socket.Write(context) // 发送数据
	if err != nil {
		log.Println("发送数据失败，err: ", err)
		return
	}
	data := make([]byte, 4096)
	n, remoteAddr, err := socket.ReadFromUDP(data) // 接收数据
	if err != nil {
		log.Println("接收数据失败, err: ", err)
		return
	}
	fmt.Printf("recv:%v addr:%v count:%v\n", string(data[:n]), remoteAddr, n)

}
