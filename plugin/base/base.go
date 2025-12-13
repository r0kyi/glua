package base

import (
	"encoding/base32"
	"encoding/base64"

	"github.com/r0kyi/glua/core"
)

type Base struct {
	b32 *base32.Encoding
	b64 *base64.Encoding
}

func (b *Base) base32Encode(raw string) string {
	encoded := b.b32.EncodeToString(core.S2B(raw))

	return encoded
}

func (b *Base) base32Decode(encoded string) string {
	raw, _ := b.b32.DecodeString(encoded)

	return core.B2S(raw)
}

func (b *Base) base64Encode(raw string) string {
	encoded := b.b64.EncodeToString(core.S2B(raw))

	return encoded
}

func (b *Base) base64Decode(encoded string) string {
	raw, _ := b.b64.DecodeString(encoded)
	return core.B2S(raw)
}
