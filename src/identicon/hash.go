package identicon

import (
	"crypto/md5"
	"fmt"
)

type Hash struct {
	hash []byte // 4-bit data array
}

func generateMd5Hash(input string) Hash {
	// get MD5 hash value from input string
	md5Hash := md5.Sum([]byte(input))
	fmt.Printf("Hash value of \"%s\": %x\n", input, md5Hash)

	nibbles := make([]byte, 0, 4*len(md5Hash))
	for _, B := range md5Hash {
		nibbles = append(nibbles, (B>>4)&0x0f, B&0x0f)
	}

	return Hash{
		hash: nibbles,
	}
}

func (h Hash) getForeground() []byte {
	// get foreground binary from 4 bit (nibble) parity
	parity := make([]byte, 0, 15)
	for _, nibble := range h.hash {
		// odd (background) = 0, even (foreground) = 1
		parity = append(parity, ^nibble&1)
	}

	// build foreground shape by rearrangement and reflection

	return parity
}

func (h Hash) getColor() (Rgb, error) {
	// get hsl values from last 7 hexadecimals, mapped as "HHHSSLL"
	hhh := h.hash[len(h.hash)-7 : len(h.hash)-4]
	ss := h.hash[len(h.hash)-4 : len(h.hash)-2]
	ll := h.hash[len(h.hash)-2:]

	// remapped to 0 <= hue < 360
	hue := convertBytesToPercentage(hhh) * 360
	if hue == 360 {
		hue = 0
	}
	// remapped to 0 <= saturation <= 0.2, and subtracted from max 0.65
	saturation := 0.65 - convertBytesToPercentage(ss)*0.2
	// remapped to 0 <= lightness <= 0.2, and subtracted from max 0.75
	lightness := 0.75 - convertBytesToPercentage(ll)*0.2

	hsl, err := newHsl(hue, saturation, lightness)
	if err != nil {
		return Rgb{}, err
	}
	rgba, err := hsl.convertHslToRgb()
	if err != nil {
		return Rgb{}, err
	}

	return rgba, nil
}
