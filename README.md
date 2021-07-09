[![Open in Visual Studio Code](https://open.vscode.dev/badges/open-in-vscode.svg)](https://open.vscode.dev/go-paynow/paynow)

# PayNow
A Go package implementing QRCode generator for Singapore PayNow.

Uses [yeqown qrcode package](https://github.com/yeqown/go-qrcode) to generate qrcodes.

## Usage
```golang
import "github.com/jaynzr/go-paynow/paynow"

// default qrcode generated in 1300px with PayNow logo
payee := paynow.Payee{}

jpg, err := payee.QRCode(12.35, "ABCDEFG")
```

To customize image output, append [qrcode.ImageOption](https://github.com/yeqown/go-qrcode#options) to Payee.Options

Using custom logo, image attributes, expiry date, and image options
```golang

import (
    "github.com/jaynzr/go-paynow/paynow"
	qrcode "github.com/yeqown/go-qrcode"
)

payee := paynow.Payee{
	MerchantName: "ACME Pte Ltd",
	UEN:          "S99345678ABCD",
}

// see https://github.com/yeqown/go-qrcode#options for available ImageOptions
payee.Options = []qrcode.ImageOption{
    // logo size should not be > 1/5 of qrcode.
    qrcode.WithLogoImageFilePNG("logo.png"),

    // width of each qr block
    qrcode.WithQRWidth(12),

    // generates png format
    qrcode.WithBuiltinImageEncoder(qrcode.PNG_FORMAT),
}

amount := 12.35
ref := "ABCDEFG"
editable := false
expiry := time.Now().Add(time.Hour * 48)

jpg, err := payee.QRCodeExpiry(amount, ref, editable, expiry)

```

# Credits
javascript/node.js

[PaynowQR](https://github.com/ThunderQuoteTeam/PaynowQR)

[QRGenerator.js](https://github.com/jtaych/PayNow-QR-Javascript/blob/master/QRGenerator.js)