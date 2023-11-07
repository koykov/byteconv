package fastconv

// B2S is a shorthand alias of BytesToString().
func B2S(b []byte) string {
	return BytesToString(b)
}

// S2B is a shorthand alias of StringToBytes().
func S2B(s string) []byte {
	return StringToBytes(s)
}

// AppendB2R is a shorthand alias of AppendBytesToRunes().
func AppendB2R(dst []rune, b []byte) []rune {
	return AppendBytesToRunes(dst, b)
}

// AppendR2B is a shorthand alias of AppendRunesToBytes().
func AppendR2B(dst []byte, r []rune) []byte {
	return AppendRunesToBytes(dst, r)
}

// AppendS2R is a shorthand alias of AppendStringToRunes().
func AppendS2R(dst []rune, s string) []rune {
	return AppendStringToRunes(dst, s)
}

// AppendR2S is a shorthand alias of AppendRunesToString().
func AppendR2S(buf []byte, r []rune) ([]byte, string) {
	return AppendRunesToString(buf, r)
}
