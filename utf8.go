package fastconv

const (
	xx = 0xF1
	as = 0xF0
	s1 = 0x02
	s2 = 0x13
	s3 = 0x03
	s4 = 0x23
	s5 = 0x34
	s6 = 0x04
	s7 = 0x44
)

var (
	first = [256]uint8{
		as, as, as, as, as, as, as, as, as, as, as, as, as, as, as, as,
		as, as, as, as, as, as, as, as, as, as, as, as, as, as, as, as,
		as, as, as, as, as, as, as, as, as, as, as, as, as, as, as, as,
		as, as, as, as, as, as, as, as, as, as, as, as, as, as, as, as,
		as, as, as, as, as, as, as, as, as, as, as, as, as, as, as, as,
		as, as, as, as, as, as, as, as, as, as, as, as, as, as, as, as,
		as, as, as, as, as, as, as, as, as, as, as, as, as, as, as, as,
		as, as, as, as, as, as, as, as, as, as, as, as, as, as, as, as,
		xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx,
		xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx,
		xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx,
		xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx,
		xx, xx, s1, s1, s1, s1, s1, s1, s1, s1, s1, s1, s1, s1, s1, s1,
		s1, s1, s1, s1, s1, s1, s1, s1, s1, s1, s1, s1, s1, s1, s1, s1,
		s2, s3, s3, s3, s3, s3, s3, s3, s3, s3, s3, s3, s3, s4, s3, s3,
		s5, s6, s6, s6, s7, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx,
	}
)
