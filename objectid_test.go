package objectid

import (
	"fmt"
	"os"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	objectRand = [5]byte{0x5f, 0x9e, 0x6b, 0x5f, 0x9e}
	counter.Store(0x6b5f9e)
	now = func() time.Time {
		return time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC)
	}
	code := m.Run()
	os.Exit(code)
}

func ExampleNew() {
	counter.Store(0x6b5f9e) // as a user of this package, you don't need to do this, but it's done here to ensure the same output

	g := New()
	fmt.Println(g.Hex())
	// Output: 5c2aad805f9e6b5f9e6b5f9f
}

func ExampleFromHex() {
	g, err := FromHex("5c2aad805f9e6b5f9e6b5fa3")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(g.Hex())
	}
	// Output: 5c2aad805f9e6b5f9e6b5fa3
}

func ExampleGenerateTo() {
	counter.Store(0x6b5f9e) // as a user of this package, you don't need to do this, but it's done here to ensure the same output

	var oid ObjectID
	GenerateTo(&oid)
	fmt.Println(oid.Hex())
	// Output: 5c2aad805f9e6b5f9e6b5f9f
}

func TestGenerate(t *testing.T) {
	t.Run("Equals", func(t *testing.T) {
		g := New()
		g2 := [12]byte{}
		for i := 0; i < len(g); i++ {
			g2[i] = g[i]
		}

		if g != g2 {
			t.Errorf("Generate() = %v, want %v", g, g2)
		}
	})
}

func TestGenerateTo(t *testing.T) {
	t.Run("Equals", func(t *testing.T) {
		var g ObjectID
		GenerateTo(&g)
		g2 := ObjectID{}
		for i := 0; i < len(g); i++ {
			g2[i] = g[i]
		}

		if g != g2 {
			t.Errorf("GenerateTo() = %v, want %v", g, g2)
		}
	})
}

func TestTime(t *testing.T) {
	g := New()
	g2 := g.Time()

	if g2.Unix() != g.Time().Unix() {
		t.Errorf("Time() = %v, want %v", g2, g.Time())
	}
}

func TestFromHex(t *testing.T) {
	g := New()
	g2, err := FromHex(g.Hex())

	if err != nil {
		t.Errorf("FromHex(%v) = %v, want nil", g.Hex(), err)
	}

	if g != g2 {
		t.Errorf("FromHex(%v) = %v, want %v", g.Hex(), g2, g)
	}
}

func TestFromHexError(t *testing.T) {
	shortHex := "5f9e6b"
	_, err := FromHex(shortHex)

	if err == nil {
		t.Errorf("FromHex(%v) = %v, want %v", shortHex, err, ErrInvalidHex)
	}

	nonHex := "gg9e6b5f9e6b5f9e6b5f9e6b"
	_, err = FromHex(nonHex)
	if err == nil {
		t.Errorf("FromHex(%v) = %v, want %v", nonHex, err, ErrInvalidHex)
	}
}

func TestFromBase64(t *testing.T) {
	g := New()
	g2, err := FromBase64(g.Base64())

	if err != nil {
		t.Errorf("FromBase64(%v) = %v, want nil", g.Base64(), err)
	}

	if g != g2 {
		t.Errorf("FromBase64(%v) = %v, want %v", g.Base64(), g2, g)
	}
}

func TestFromBase64Error(t *testing.T) {
	shortBase64 := "5f9e6b"
	_, err := FromBase64(shortBase64)

	if err == nil {
		t.Errorf("FromBase64(%v) = %v, want %v", shortBase64, err, ErrInvalidBase64)
	}

	invalidBase64 := "ZJH6hPBsxutRK2!$"
	if _, err := FromBase64(invalidBase64); err == nil {
		t.Errorf("FromBase64(%v) = %v, want %v", invalidBase64, err, ErrInvalidBase64)
	}
}
func Benchmark_Generate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = New()
	}
}

func Benchmark_GenerateTo(b *testing.B) {
	var oid ObjectID
	for i := 0; i < b.N; i++ {
		GenerateTo(&oid)
	}
}

func Benchmark_FromHex(b *testing.B) {
	g := New()
	for i := 0; i < b.N; i++ {
		_, _ = FromHex(g.Hex())
	}
}
