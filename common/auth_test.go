package common

import (
	"fmt"
	"testing"
)

func TestGetQrCode(t *testing.T) {

	a := NewGoogleAuth()

	secret := a.GetSecret()
	//secret="secret"
	fmt.Println(secret)

	qrcode := a.GetQrcode("home-wol", secret)

	qrcodePic, _ := QrcodePictureEncode(qrcode, 200, 200)
	fmt.Println(qrcodePic)
	url := a.GetQrcodeUrl("home-wol", secret)
	fmt.Println(url)
	code, err := a.GetCode(a.GetSecret())

	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(code)
}
