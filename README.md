[![Open in Visual Studio Code](https://open.vscode.dev/badges/open-in-vscode.svg)](https://open.vscode.dev/go-paynow/paynow)

# PayNow
A Go package implementing QRCode generator for Singapore PayNow.

The main package implements the method to generate the value string for PayNow.

Optional package `"github.com/jaynzr/go-paynow/qrcode"` uses [yeqown qrcode package](https://github.com/yeqown/go-qrcode) to generate qrcodes. You may use other qrcode generator.

## Generating PayNow Value
```golang

import "github.com/jaynzr/go-paynow"

payee := paynow.NewUEN("ACME Pte Ltd", "S99345678ABCD")
val : payee.New(12.34, "INV1234", true, time.Time{}).String()


```

## QR Code Usage
```golang
import "github.com/jaynzr/go-paynow/qrcode"

payee := qrcode.NewMobile("90991234")

jpg, err := qrcode.QRCode(12.35, "ABCDEFG")
```

To customize image output, append [qrcode.ImageOption](https://github.com/yeqown/go-qrcode#options) to Payee.Options

Using custom logo, image attributes, expiry date, and image options
```golang

import (
    "github.com/jaynzr/go-paynow/qrcode"
	qrcopt "github.com/yeqown/go-qrcode"
)

payee := qrcode.NewUEN("ACME Pte Ltd", "S99345678ABCD")

// see https://github.com/yeqown/go-qrcode#options for available ImageOptions
payee.Options = []qrcopt.ImageOption{
    // logo size should not be > 1/5 of qrcode.
    qrcopt.WithLogoImageFilePNG("logo.png"),

    // width of each qr block
    qrcopt.WithQRWidth(12),

    // generates png format
    qrcopt.WithBuiltinImageEncoder(qrcopt.PNG_FORMAT),
}

amount := 12.35
ref := "ABCDEFG"
editable := false
expiry := time.Now().Add(time.Hour * 48)

png, err := payee.QRCodeExpiry(amount, ref, editable, expiry)

```

# Credits
javascript/node.js

[PaynowQR](https://github.com/ThunderQuoteTeam/PaynowQR)

[QRGenerator.js](https://github.com/jtaych/PayNow-QR-Javascript/blob/master/QRGenerator.js)