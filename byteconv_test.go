package byteconv

import (
	"bytes"
	"fmt"
	"testing"
)

var stages = []struct {
	s string
	b []byte
	r []rune
}{
	{s: "", b: nil, r: nil},
	{
		"foobar",
		[]byte("foobar"),
		[]rune("foobar"),
	},
	{
		"Hello Wêreld!",
		[]byte("Hello Wêreld!"),
		[]rune("Hello Wêreld!"),
	},
	{
		"Përshendetje Botë!",
		[]byte("Përshendetje Botë!"),
		[]rune("Përshendetje Botë!"),
	},
	{
		"ሰላም ልዑል!",
		[]byte("ሰላም ልዑል!"),
		[]rune("ሰላም ልዑል!"),
	},
	{
		"مرحبا بالعالم!",
		[]byte("مرحبا بالعالم!"),
		[]rune("مرحبا بالعالم!"),
	},
	{
		"Բարեւ աշխարհ!",
		[]byte("Բարեւ աշխարհ!"),
		[]rune("Բարեւ աշխարհ!"),
	},
	{
		"Kaixo Mundua!",
		[]byte("Kaixo Mundua!"),
		[]rune("Kaixo Mundua!"),
	},
	{
		"Прывітанне Сусвет!",
		[]byte("Прывітанне Сусвет!"),
		[]rune("Прывітанне Сусвет!"),
	},
	{
		"ওহে বিশ্ব!",
		[]byte("ওহে বিশ্ব!"),
		[]rune("ওহে বিশ্ব!"),
	},
	{
		"Здравей свят!",
		[]byte("Здравей свят!"),
		[]rune("Здравей свят!"),
	},
	{
		"Hola món!",
		[]byte("Hola món!"),
		[]rune("Hola món!"),
	},
	{
		"Moni Dziko Lapansi!",
		[]byte("Moni Dziko Lapansi!"),
		[]rune("Moni Dziko Lapansi!"),
	},
	{
		"你好世界！",
		[]byte("你好世界！"),
		[]rune("你好世界！"),
	},
	{
		"Pozdrav svijete!",
		[]byte("Pozdrav svijete!"),
		[]rune("Pozdrav svijete!"),
	},
	{
		"Ahoj světe!",
		[]byte("Ahoj světe!"),
		[]rune("Ahoj světe!"),
	},
	{
		"Hej Verden!",
		[]byte("Hej Verden!"),
		[]rune("Hej Verden!"),
	},
	{
		"Hallo Wereld!",
		[]byte("Hallo Wereld!"),
		[]rune("Hallo Wereld!"),
	},
	{
		"Hello World!",
		[]byte("Hello World!"),
		[]rune("Hello World!"),
	},
	{
		"Tere maailm!",
		[]byte("Tere maailm!"),
		[]rune("Tere maailm!"),
	},
	{
		"Hei maailma!",
		[]byte("Hei maailma!"),
		[]rune("Hei maailma!"),
	},
	{
		"Bonjour monde!",
		[]byte("Bonjour monde!"),
		[]rune("Bonjour monde!"),
	},
	{
		"Hallo wrâld!",
		[]byte("Hallo wrâld!"),
		[]rune("Hallo wrâld!"),
	},
	{
		"გამარჯობა მსოფლიო!",
		[]byte("გამარჯობა მსოფლიო!"),
		[]rune("გამარჯობა მსოფლიო!"),
	},
	{
		"Hallo Welt!",
		[]byte("Hallo Welt!"),
		[]rune("Hallo Welt!"),
	},
	{
		"Γειά σου Κόσμε!",
		[]byte("Γειά σου Κόσμε!"),
		[]rune("Γειά σου Κόσμε!"),
	},
	{
		"Sannu Duniya!",
		[]byte("Sannu Duniya!"),
		[]rune("Sannu Duniya!"),
	},
	{
		"שלום עולם!",
		[]byte("שלום עולם!"),
		[]rune("שלום עולם!"),
	},
	{
		"नमस्ते दुनिया!",
		[]byte("नमस्ते दुनिया!"),
		[]rune("नमस्ते दुनिया!"),
	},
	{
		"Helló Világ!",
		[]byte("Helló Világ!"),
		[]rune("Helló Világ!"),
	},
	{
		"Halló heimur!",
		[]byte("Halló heimur!"),
		[]rune("Halló heimur!"),
	},
	{
		"Ndewo Ụwa!",
		[]byte("Ndewo Ụwa!"),
		[]rune("Ndewo Ụwa!"),
	},
	{
		"Halo Dunia!",
		[]byte("Halo Dunia!"),
		[]rune("Halo Dunia!"),
	},
	{
		"Ciao mondo!",
		[]byte("Ciao mondo!"),
		[]rune("Ciao mondo!"),
	},
	{
		"こんにちは世界！",
		[]byte("こんにちは世界！"),
		[]rune("こんにちは世界！"),
	},
	{
		"Сәлем Әлем!",
		[]byte("Сәлем Әлем!"),
		[]rune("Сәлем Әлем!"),
	},
	{
		"Салам дүйнө!",
		[]byte("Салам дүйнө!"),
		[]rune("Салам дүйнө!"),
	},
	{
		"Sveika pasaule!",
		[]byte("Sveika pasaule!"),
		[]rune("Sveika pasaule!"),
	},
	{
		"Labas pasauli!",
		[]byte("Labas pasauli!"),
		[]rune("Labas pasauli!"),
	},
	{
		"Moien Welt!",
		[]byte("Moien Welt!"),
		[]rune("Moien Welt!"),
	},
	{
		"Здраво свету!",
		[]byte("Здраво свету!"),
		[]rune("Здраво свету!"),
	},
	{
		"Hai dunia!",
		[]byte("Hai dunia!"),
		[]rune("Hai dunia!"),
	},
	{
		"ഹലോ വേൾഡ്!",
		[]byte("ഹലോ വേൾഡ്!"),
		[]rune("ഹലോ വേൾഡ്!"),
	},
	{
		"Сайн уу дэлхий!",
		[]byte("Сайн уу дэлхий!"),
		[]rune("Сайн уу дэлхий!"),
	},
	{
		"မင်္ဂလာပါကမ္ဘာလောက!",
		[]byte("မင်္ဂလာပါကမ္ဘာလောက!"),
		[]rune("မင်္ဂလာပါကမ္ဘာလောက!"),
	},
	{
		"नमस्कार संसार!",
		[]byte("नमस्कार संसार!"),
		[]rune("नमस्कार संसार!"),
	},
	{
		"Hei Verden!",
		[]byte("Hei Verden!"),
		[]rune("Hei Verden!"),
	},
	{
		"سلام نړی!",
		[]byte("سلام نړی!"),
		[]rune("سلام نړی!"),
	},
	{
		"سلام دنیا!",
		[]byte("سلام دنیا!"),
		[]rune("سلام دنیا!"),
	},
	{
		"Witaj świecie!",
		[]byte("Witaj świecie!"),
		[]rune("Witaj świecie!"),
	},
	{
		"Olá Mundo!",
		[]byte("Olá Mundo!"),
		[]rune("Olá Mundo!"),
	},
	{
		"ਸਤਿ ਸ੍ਰੀ ਅਕਾਲ ਦੁਨਿਆ!",
		[]byte("ਸਤਿ ਸ੍ਰੀ ਅਕਾਲ ਦੁਨਿਆ!"),
		[]rune("ਸਤਿ ਸ੍ਰੀ ਅਕਾਲ ਦੁਨਿਆ!"),
	},
	{
		"Salut Lume!",
		[]byte("Salut Lume!"),
		[]rune("Salut Lume!"),
	},
	{
		"Привет мир!",
		[]byte("Привет мир!"),
		[]rune("Привет мир!"),
	},
	{
		"Hàlo a Shaoghail!",
		[]byte("Hàlo a Shaoghail!"),
		[]rune("Hàlo a Shaoghail!"),
	},
	{
		"Здраво Свете!",
		[]byte("Здраво Свете!"),
		[]rune("Здраво Свете!"),
	},
	{
		"Lefatše Lumela!",
		[]byte("Lefatše Lumela!"),
		[]rune("Lefatše Lumela!"),
	},
	{
		"හෙලෝ වර්ල්ඩ්!",
		[]byte("හෙලෝ වර්ල්ඩ්!"),
		[]rune("හෙලෝ වර්ල්ඩ්!"),
	},
	{
		"Pozdravljen svet!",
		[]byte("Pozdravljen svet!"),
		[]rune("Pozdravljen svet!"),
	},
	{
		"¡Hola Mundo!",
		[]byte("¡Hola Mundo!"),
		[]rune("¡Hola Mundo!"),
	},
	{
		"Halo Dunya!",
		[]byte("Halo Dunya!"),
		[]rune("Halo Dunya!"),
	},
	{
		"Salamu Dunia!",
		[]byte("Salamu Dunia!"),
		[]rune("Salamu Dunia!"),
	},
	{
		"Hej världen!",
		[]byte("Hej världen!"),
		[]rune("Hej världen!"),
	},
	{
		"Салом Ҷаҳон!",
		[]byte("Салом Ҷаҳон!"),
		[]rune("Салом Ҷаҳон!"),
	},
	{
		"สวัสดีชาวโลก!",
		[]byte("สวัสดีชาวโลก!"),
		[]rune("สวัสดีชาวโลก!"),
	},
	{
		"Selam Dünya!",
		[]byte("Selam Dünya!"),
		[]rune("Selam Dünya!"),
	},
	{
		"Привіт Світ!",
		[]byte("Привіт Світ!"),
		[]rune("Привіт Світ!"),
	},
	{
		"Salom Dunyo!",
		[]byte("Salom Dunyo!"),
		[]rune("Salom Dunyo!"),
	},
	{
		"Chào thế giới!",
		[]byte("Chào thế giới!"),
		[]rune("Chào thế giới!"),
	},
	{
		"Helo Byd!",
		[]byte("Helo Byd!"),
		[]rune("Helo Byd!"),
	},
	{
		"Molo Lizwe!",
		[]byte("Molo Lizwe!"),
		[]rune("Molo Lizwe!"),
	},
	{
		"העלא וועלט!",
		[]byte("העלא וועלט!"),
		[]rune("העלא וועלט!"),
	},
	{
		"Mo ki O Ile Aiye!",
		[]byte("Mo ki O Ile Aiye!"),
		[]rune("Mo ki O Ile Aiye!"),
	},
	{
		"Sawubona Mhlaba!",
		[]byte("Sawubona Mhlaba!"),
		[]rune("Sawubona Mhlaba!"),
	},
}

func TestByteconv(t *testing.T) {
	for _, stage := range stages {
		t.Run(fmt.Sprintf("bytes2string/%s", stage.s), func(t *testing.T) {
			s := B2S(stage.b)
			if s != stage.s {
				t.Errorf("BytesToString mismatch, need %s, got %s", stage.s, s)
			}
		})
		t.Run(fmt.Sprintf("string2bytes/%s", stage.s), func(t *testing.T) {
			b := S2B(stage.s)
			if !bytes.Equal(b, stage.b) {
				t.Errorf("StringToBytes mismatch, need %s, got %s", string(stage.b), string(b))
			}
		})
		t.Run(fmt.Sprintf("bytes2runes/%s", stage.s), func(t *testing.T) {
			var buf []rune
			buf = AppendB2R(buf[:0], stage.b)
			if !eqr(buf, stage.r) {
				t.Errorf("BytesToRunes mismatch, need %v, got %v", stage.r, buf)
			}
		})
		t.Run(fmt.Sprintf("runes2bytes/%s", stage.s), func(t *testing.T) {
			var buf []byte
			buf = AppendR2B(buf[:0], stage.r)
			if !bytes.Equal(buf, stage.b) {
				t.Errorf("AppendRunesToBytes mismatch, need %s, got %s", string(stage.b), string(buf))
			}
		})
		t.Run(fmt.Sprintf("string2runes/%s", stage.s), func(t *testing.T) {
			var buf []rune
			buf = AppendS2R(buf[:0], stage.s)
			if !eqr(buf, stage.r) {
				t.Errorf("AppendStringToRunes mismatch, need %v, got %v", stage.r, buf)
			}
		})
		t.Run(fmt.Sprintf("runes2string/%s", stage.s), func(t *testing.T) {
			var (
				buf []byte
				s   string
			)
			buf, s = AppendR2S(buf[:0], stage.r)
			if s != stage.s {
				t.Errorf("AppendRunesToString mismatch, need %s, got %s", stage.s, s)
			}
		})
	}
}

func BenchmarkByteconv(b *testing.B) {
	b.Run("bytes2string", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			stage := stages[i%len(stages)]
			s := B2S(stage.b)
			if s != stage.s {
				b.Errorf("BytesToString mismatch, need %s, got %s", stage.s, s)
			}
		}
	})
	b.Run("string2bytes", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			stage := stages[i%len(stages)]
			x := S2B(stage.s)
			if !bytes.Equal(x, stage.b) {
				b.Errorf("StringToBytes mismatch, need %s, got %s", string(stage.b), string(x))
			}
		}
	})
	b.Run("bytes2runes", func(b *testing.B) {
		var buf []rune
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			stage := stages[i%len(stages)]
			buf = AppendB2R(buf[:0], stage.b)
			if !eqr(buf, stage.r) {
				b.Errorf("BytesToRunes mismatch, need %v, got %v", stage.r, buf)
			}
		}
	})
	b.Run("runes2bytes", func(b *testing.B) {
		var buf []byte
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			stage := stages[i%len(stages)]
			buf = AppendR2B(buf[:0], stage.r)
			if !bytes.Equal(buf, stage.b) {
				b.Errorf("AppendRunesToBytes mismatch, need %s, got %s", string(stage.b), string(buf))
			}
		}
	})
	b.Run("string2runes", func(b *testing.B) {
		var buf []rune
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			stage := stages[i%len(stages)]
			buf = AppendS2R(buf[:0], stage.s)
			if !eqr(buf, stage.r) {
				b.Errorf("AppendStringToRunes mismatch, need %v, got %v", stage.r, buf)
			}
		}
	})
	b.Run("runes2string", func(b *testing.B) {
		var (
			buf []byte
			s   string
		)
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			stage := stages[i%len(stages)]
			buf, s = AppendR2S(buf[:0], stage.r)
			if s != stage.s {
				b.Errorf("AppendRunesToString mismatch, need %s, got %s", stage.s, s)
			}
		}
	})
}

func eqr(a, b []rune) bool {
	la, lb := len(a), len(b)
	if la != lb {
		return false
	}
	for i := 0; i < la; i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
