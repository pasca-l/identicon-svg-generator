package identicon

import (
	"crypto/md5"
	"fmt"
)

type Hash struct {
	hash []byte
}

func generateMd5Hash(input string) Hash {
	// get MD5 hash value from input string
	md5Hash := md5.Sum([]byte(input))
	fmt.Printf("Hash value of \"%s\": %x\n", input, md5Hash)

	return Hash{
		// slicing the array
		hash: md5Hash[:],
	}
}

func (h Hash) getForeground() []byte {
	// get foreground binary from byte parity
	parity := make([]byte, 0, 15)
	for _, B := range h.hash {
		// odd = 1, even = 0
		parity = append(parity, B&1)
	}

	// build foreground shape by rearrangement and reflection

	return parity
}

func (h Hash) getColor() (Rgb, error) {
	// get hue from last 7 hexadecimals
	// shift 1 byte to left before casting to 64-bit type
	bytes := append(h.hash[len(h.hash)-7:], 0)
	hue := convertByteToDegrees(bytes)

	hsl, err := newHsl(hue, 0.7, 0.5)
	if err != nil {
		return Rgb{}, err
	}
	rgba, err := hsl.convertHslToRgb()
	if err != nil {
		return Rgb{}, err
	}

	return rgba, nil
}
