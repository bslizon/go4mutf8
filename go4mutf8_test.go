package go4mutf8

import (
	"encoding/hex"
	"fmt"
	"math/rand"
	"testing"
)

func TestEncode(t *testing.T) {
	ms, err := Encode("🐒MyApplication")
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(hex.EncodeToString(ms))
}
func TestDecode(t *testing.T) {
	b, err := hex.DecodeString("eda0bdedb0924d794170706c69636174696f6e")
	if err != nil {
		t.Error(err)
		return
	}

	s, err := Decode(b)
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(s)
}

func TestFuzzEncodeDecode(t *testing.T) {
	c := 0
	for {
		rr := []rune{}

		randC := rand.Intn(500) + 1

		for i := 0; i < randC; i++ {
			rr = append(rr, getRandRune())
		}

		s := string(rr)

		ms, err := Encode(s)
		if err != nil {
			fmt.Println(err)
			return
		}

		ss, err := Decode(ms)
		if err != nil {
			fmt.Println(err)
			return
		}

		if s != ss {
			fmt.Printf("%v != %v", s, ss)
			return
		}

		c++
		if c%10000 == 0 {
			fmt.Println(c)
		}
	}
}

func getRandRune() rune {
	return rune(rand.Intn(0x110000))
}
