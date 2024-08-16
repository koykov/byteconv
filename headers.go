package byteconv

// SliceHeader represents deprecated reflect.SliceHeader type.
type SliceHeader struct {
	Data uintptr
	Len  int
	Cap  int
}

// StringHeader represents deprecated reflect.StringHeader type.
type StringHeader struct {
	Data uintptr
	Len  int
}
