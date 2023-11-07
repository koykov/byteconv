package fastconv

// Fast alloc-free conversion algorithms between byte sequences and strings.

import (
	"reflect"
	"unicode/utf8"
	"unsafe"
)

// BytesToString makes fast conversion of bytes array to string.
func BytesToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// StringToBytes makes fast conversion of string to bytes sequence.
func StringToBytes(s string) []byte {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	var h reflect.SliceHeader
	h.Data = sh.Data
	h.Len = sh.Len
	h.Cap = sh.Len
	return *(*[]byte)(unsafe.Pointer(&h))
}

// AppendBytesToRunes converts byte array to runes array.
func AppendBytesToRunes(dst []rune, p []byte) []rune {
	l := len(p)
	if l == 0 {
		return dst
	}
	_ = p[l-1]
	for i := 0; i < l; {
		c := p[i]
		if c < utf8.RuneSelf {
			dst = append(dst, rune(c))
			i++
			continue
		}
		x := first[c]
		if x == xx { // invalid.
			i++
			continue
		}
		sz := int(x & 7)
		if i+sz > l { // Short or invalid.
			i++
			continue
		}
		r, _ := utf8.DecodeRune(p[i : i+sz])
		dst = append(dst, r)
		i += sz
	}
	return dst
}

// AppendRunesToBytes converts runes array to bytes array.
func AppendRunesToBytes(dst []byte, r []rune) []byte {
	l := len(r)
	if l == 0 {
		return dst
	}
	_ = r[l-1]
	for i := 0; i < l; i++ {
		dst = utf8.AppendRune(dst, r[i])
	}
	return dst
}

// AppendStringToRunes converts string to runes array.
func AppendStringToRunes(dst []rune, s string) []rune {
	return AppendBytesToRunes(dst, S2B(s))
}

// AppendRunesToString converts runes array to string using byte buffer.
func AppendRunesToString(buf []byte, r []rune) ([]byte, string) {
	off := len(buf)
	buf = AppendRunesToBytes(buf, r)
	return buf, B2S(buf[off:])
}
