package base

import (
	"encoding/base32"
	"encoding/base64"

	"github.com/r0kyi/glua/core"
)

type Base struct {
	raw     string
	encoded string
}

func (b *Base) base32Encode() {
	b.encoded = base32.StdEncoding.EncodeToString(core.S2B(b.raw))
}

func (b *Base) base32Decode() {
	raw, _ := base32.StdEncoding.DecodeString(b.encoded)
	b.encoded = core.B2S(raw)
}

func (b *Base) base64Encode() {
	b.encoded = base64.StdEncoding.EncodeToString(core.S2B(b.raw))
}

func (b *Base) base64Decode() {
	raw, _ := base64.StdEncoding.DecodeString(b.encoded)
	b.raw = core.B2S(raw)
}
