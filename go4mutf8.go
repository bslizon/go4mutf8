package go4mutf8

import (
	"bytes"
	"fmt"
)

func Decode(b []byte) (string, error) {
	if len(b) <= 0 {
		return "", fmt.Errorf("len(b) <= 0")
	}

	rr := []rune{}
	index := 0

	for index < len(b) {
		switch {
		case (b[index] & 0b10000000) == 0b00000000:
			rr = append(rr, rune(b[index]))
			index++

		case ((b[index] & 0b11100000) == 0b11000000) &&
			((b[index+1] & 0b11000000) == 0b10000000):
			x := b[index]
			y := b[index+1]
			rr = append(rr, rune((int(x)&0x1f)<<6+int(y)&0x3f))
			index += 2

		case (b[index] == 0b11101101) &&
			((b[index+1] & 0b11110000) == 0b10100000) &&
			((b[index+2] & 0b11000000) == 0b10000000) &&
			(b[index+3] == 0b11101101) &&
			((b[index+4] & 0b11110000) == 0b10110000) &&
			((b[index+5] & 0b11000000) == 0b10000000):
			_ = b[index]
			v := b[index+1]
			w := b[index+2]
			_ = b[index+3]
			y := b[index+4]
			z := b[index+5]
			rr = append(rr, rune(0x10000+(int(v)&0x0f)<<16+(int(w)&0x3f)<<10+(int(y)&0x0f)<<6+int(z)&0x3f))
			index += 6

		case ((b[index] & 0b11110000) == 0b11100000) &&
			((b[index+1] & 0b11000000) == 0b10000000) &&
			((b[index+2] & 0b11000000) == 0b10000000):
			x := b[index]
			y := b[index+1]
			z := b[index+2]
			rr = append(rr, rune(int(x)&0xf<<12+(int(y)&0x3f)<<6+int(z)&0x3f))
			index += 3

		default:
			return "", fmt.Errorf("unexpected byte: %b, index: %v", b[0], index)
		}
	}

	return string(rr), nil

}

func Encode(s string) ([]byte, error) {
	src := []rune(s)

	dst := bytes.Buffer{}

	for _, r := range src {
		if r < 0 {
			return nil, fmt.Errorf("rune:%v not supported", r)
		} else if 0 < r && r <= 0x7F {
			dst.WriteByte(byte(r))
		} else if r <= 0x7FF { //include 0x00
			dst.WriteByte(byte(0b11000000 | (0b00011111 & (r >> 6))))
			dst.WriteByte(byte(0b10000000 | (0b00111111 & r)))
		} else if r <= 0xFFFF {
			dst.WriteByte(byte(0b11100000 | (0b00001111 & (r >> 12))))
			dst.WriteByte(byte(0b10000000 | (0b00111111 & (r >> 6))))
			dst.WriteByte(byte(0b10000000 | (0b00111111 & r)))
		} else if r <= 0x10FFFF {
			r2 := r - 0x10000
			dst.WriteByte(byte(0b11101101))
			dst.WriteByte(byte(0b10100000 | (0b00001111 & (r2 >> 16))))
			dst.WriteByte(byte(0b10000000 | (0b00111111 & (r2 >> 10))))
			dst.WriteByte(byte(0b11101101))
			dst.WriteByte(byte(0b10110000 | (0b00001111 & (r2 >> 6))))
			dst.WriteByte(byte(0b10000000 | (0b00111111 & r2)))
		} else {
			return nil, fmt.Errorf("rune:%v not supported", r)
		}
	}

	return dst.Bytes(), nil
}
