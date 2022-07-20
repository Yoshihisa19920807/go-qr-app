package qrgen

import (
	"bytes"
	"image"
	"image/png"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
)

func CreateImage(b []byte) (image.Image, error) {
	return png.Decode(bytes.NewReader(b))
}

// GenQRCode generates QR code for the given URL
func GenQRCode(url string, width, height int) (barcode.Barcode, error) {
	// Encode()の引数は
	// 第１：対象のURL
	// 第２：QRコードの誤り訂正レベル(汚れなどがある場合、どの程度まで読み込めなくてもいいか)
	// 第３：エンコードの形式
	qrCode, err := qr.Encode(url, qr.M, qr.Auto)
	if err != nil {
		return nil, err
	}
	return barcode.Scale(qrCode, width, height)
}
