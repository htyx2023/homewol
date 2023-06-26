package common

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"github.com/makiuchi-d/gozxing"
	"github.com/makiuchi-d/gozxing/qrcode"
	"image"
	"image/jpeg"
)

var qrcodeReader gozxing.Reader
var qrcodeWriter gozxing.Writer

func init() {
	qrcodeReader = qrcode.NewQRCodeReader()
	qrcodeWriter = qrcode.NewQRCodeWriter()
}

// Encode 二维码编码
func QrcodePictureEncode(text string, width int, height int) (string, error) {
	hints := make(map[gozxing.EncodeHintType]interface{})
	hints[gozxing.EncodeHintType_MARGIN] = 0

	bm, e := qrcodeWriter.Encode(text, gozxing.BarcodeFormat_QR_CODE, width, height, hints)
	if e != nil {
		return "", e
	}

	bb := bytes.NewBuffer([]byte{})
	e = jpeg.Encode(bb, bm, nil)
	if e != nil {
		return "", e
	}

	base64Text := base64.StdEncoding.EncodeToString(bb.Bytes())

	return fmt.Sprint("data:image/png;base64,", base64Text), nil
}

// DecodeBytes 二维码解码图片字节
func DecodeBytes(data []byte) (string, error) {
	img, _, e := image.Decode(bytes.NewReader(data))
	if e != nil {
		return "", e
	}

	// prepare BinaryBitmap
	bmp, e := gozxing.NewBinaryBitmapFromImage(img)
	if e != nil {
		return "", e
	}

	// decode image
	result, e := qrcodeReader.Decode(bmp, nil)
	if e != nil {
		return "", e
	}

	return result.GetText(), nil
}
