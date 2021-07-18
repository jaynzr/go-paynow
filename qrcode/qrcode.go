// Package paynow generates QR Code for Singapore PayNow payment system.
package qrcode

import (
	"bytes"
	"time"

	paynow "github.com/jaynzr/go-paynow"

	qrcode "github.com/yeqown/go-qrcode"
)

// QRSpec to set qrcode version
type QRSpec struct {
	Version int
}

// Payee struct, all fields are optional.
type Payee struct {
	paynow.Payee
	Options []qrcode.ImageOption
	Spec    QRSpec // optional qrcode version
}

func NewUEN(merchantName string, uen string) *Payee {
	pp := &Payee{
		Payee: *paynow.NewUEN(merchantName, uen),
	}
	return pp
}

func NewMobile(mobile string) *Payee {
	pp := &Payee{
		Payee: *paynow.NewMobile(mobile),
	}
	return pp
}

// QRCode returns editable qr code image
func (pp *Payee) QRCode(amount float32, ref string) ([]byte, error) {
	return pp.QRCodeExpiry(amount, ref, true, time.Time{})
}

// QRCodeLocked returns uneditable qr code image
func (pp *Payee) QRCodeLocked(amount float32, ref string) ([]byte, error) {
	return pp.QRCodeExpiry(amount, ref, false, time.Time{})
}

// QRCodeExpiry returns QR Code image, in jpeg format by default
// expirySGT should be in Asia/Singapore timezone. If expirySGT is zero, perpetual
func (pp *Payee) QRCodeExpiry(amount float32, ref string, editable bool, expirySGT time.Time) ([]byte, error) {
	var (
		buf bytes.Buffer
		qrc *qrcode.QRCode
		err error
	)

	value := pp.New(amount, ref, editable, expirySGT).String()

	if pp.Spec.Version != 0 {
		qrc, err = qrcode.NewWithSpecV(value, pp.Spec.Version, qrcode.Medium, pp.Options...)
	} else {
		qrc, err = qrcode.New(value, pp.Options...)
	}

	if err != nil {
		return nil, err
	}

	err = qrc.SaveTo(&buf)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
