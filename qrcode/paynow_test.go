package qrcode

import (
	"net/http"
	"testing"
	"time"

	qrcode "github.com/yeqown/go-qrcode"
)

func TestQRCode(t *testing.T) {
	payee := NewUEN("ACME Pte Ltd", "S99345678ABCD")

	// see https://github.com/yeqown/go-qrcode#options for available ImageOptions
	payee.Options = []qrcode.ImageOption{
		// width of each qr block
		qrcode.WithQRWidth(12),

		// generates png format
		qrcode.WithBuiltinImageEncoder(qrcode.PNG_FORMAT),
	}

	amount := float32(999.12)
	ref := "ABCDEFG"
	editable := false
	expiry := time.Now()

	png, err := payee.QRCodeExpiry(amount, ref, editable, expiry)
	if err != nil {
		t.Fatalf("could not create qrcode: %v", err)
	}

	ctype := http.DetectContentType(png)
	if ctype != "image/png" {
		t.Fatalf("invalid file format %s", ctype)
	}
}

func TestMobileQRCode(t *testing.T) {
	payee := NewMobile("99991234")

	// see https://github.com/yeqown/go-qrcode#options for available ImageOptions
	payee.Options = []qrcode.ImageOption{
		// width of each qr block
		qrcode.WithQRWidth(12),

		// generates png format
		qrcode.WithBuiltinImageEncoder(qrcode.PNG_FORMAT),
	}

	amount := float32(999.12)
	ref := "John"
	editable := false
	expiry := time.Now()

	png, err := payee.QRCodeExpiry(amount, ref, editable, expiry)
	if err != nil {
		t.Fatalf("could not create qrcode: %v", err)
	}

	ctype := http.DetectContentType(png)
	if ctype != "image/png" {
		t.Fatalf("invalid file format %s", ctype)
	}
}
