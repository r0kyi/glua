package core

import (
	"unsafe"
)

func S2B(s string) (b []byte) {
	ptr := unsafe.StringData(s)
	return unsafe.Slice(ptr, len(s))
}

func B2S(b []byte) string {
	if len(b) == 0 {
		return ""
	}
	return unsafe.String(unsafe.SliceData(b), len(b))
}
