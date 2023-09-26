package qrcode

import (
	"fmt"
	"image/png"
	"os"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	"github.com/tuotoo/qrcode"
)

func QRCEncode(content, fileName string) (err error) {
	qrCode, err := qr.Encode(
		content,
		qr.M,
		qr.Auto,
	)
	if err != nil {
		return
	}
	qrCode, err = barcode.Scale(qrCode, 256, 256)
	if err != nil {
		return
	}
	file, err := os.Create(fileName)
	if err != nil {
		return
	}
	defer file.Close()
	png.Encode(file, qrCode)
	return
}

func QRCDecode(fileName string) (content string, err error) {
	fi, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fi.Close()
	qrMatrix, err := qrcode.Decode(fi)
	if err != nil {
		fmt.Println(err)
		return
	}
	content = qrMatrix.Content
	return
}
