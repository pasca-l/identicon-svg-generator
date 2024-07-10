package identicon

import (
	"fmt"
	"math"
)

type Rgb struct {
	r uint8
	g uint8
	b uint8
}

type Hsl struct {
	h float64
	s float64
	l float64
}

func newHsl(h, s, l float64) (Hsl, error) {
	if h < 0 || h >= 360 {
		return Hsl{}, fmt.Errorf(
			"hue out of range, must be within 0~360: %x", h,
		)
	}
	if s < 0 || s > 1 {
		return Hsl{}, fmt.Errorf(
			"saturation out of range, must be within 0~1: %x", s,
		)
	}
	if l < 0 || l > 1 {
		return Hsl{}, fmt.Errorf(
			"lightness out of range, must be within 0~1: %x", l,
		)
	}

	return Hsl{h: h, s: s, l: l}, nil
}

func convertBytesToPercentage(b []byte) float64 {
	var combined uint64 = 0
	for i, nibble := range b {
		combined = combined | uint64(nibble)<<(4*(len(b)-i-1))
	}
	f := float64(combined)
	max := math.Pow(2, float64(4*len(b)))
	return f / max
}

func (hsl Hsl) convertHslToRgb() (Rgb, error) {
	// formula from https://en.wikipedia.org/wiki/HSL_and_HSV#HSL_to_RGB
	chroma := (1 - math.Abs(2*hsl.l-1)) * hsl.s
	h_prime := hsl.h / 60
	x := chroma * (1 - math.Abs(math.Mod(h_prime, 2.0)-1))
	m := hsl.l - chroma/2

	byte_c := uint8((chroma + m) * 255)
	byte_x := uint8((x + m) * 255)
	byte_0 := uint8(m * 255)

	switch {
	case 0 <= h_prime && h_prime < 1:
		return Rgb{byte_c, byte_x, byte_0}, nil
	case 1 <= h_prime && h_prime < 2:
		return Rgb{byte_x, byte_c, byte_0}, nil
	case 2 <= h_prime && h_prime < 3:
		return Rgb{byte_0, byte_c, byte_x}, nil
	case 3 <= h_prime && h_prime < 4:
		return Rgb{byte_0, byte_x, byte_c}, nil
	case 4 <= h_prime && h_prime < 5:
		return Rgb{byte_x, byte_0, byte_c}, nil
	case 5 <= h_prime && h_prime < 6:
		return Rgb{byte_c, byte_0, byte_x}, nil
	default:
		return Rgb{}, fmt.Errorf("unexpected value of h_prime: %x", h_prime)
	}
}
