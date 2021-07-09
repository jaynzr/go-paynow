package paynow

import (
	"net/http"
	"os"
	"testing"
	"time"

	qrcode "github.com/yeqown/go-qrcode"
)

func init() {
	os.Remove("qrcode.png")
}

func TestPayNow(t *testing.T) {
	payee := Payee{
		MerchantName: "ACME Pte Ltd",
		UEN:          "S99345678ABCD",
	}

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

	jpg, err := payee.QRCodeExpiry(amount, ref, editable, expiry)
	if err != nil {
		t.Fatalf("could not create qrcode: %v", err)
	}

	ctype := http.DetectContentType(jpg)
	if ctype != "image/png" {
		t.Fatalf("invalid file format %s", ctype)
	}
}
