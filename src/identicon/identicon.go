package identicon

import (
	"fmt"
	"os"
)

func GenerateIdenticon(accountId string) error {
	hash := generateMd5Hash(accountId)

	foreground, err := getForeground(hash)
	if err != nil {
		return err
	}
	color, err := getColor(hash)
	if err != nil {
		return err
	}

	fmt.Println(foreground, color)
	drawIdenticon(os.Stdout, foreground, color)

	return nil
}

func getForeground(h Hash) (Array[byte], error) {
	// get foreground binary from 4 bit (nibble) parity
	parity := make([]byte, 0, 15)
	for _, nibble := range h.hash {
		// odd (background) = 0, even (foreground) = 1
		parity = append(parity, ^nibble&1)
	}

	// build foreground shape by rearrangement and reflection
	array, err := convertListToArray(parity[:15], 3)
	if err != nil {
		return Array[byte]{}, err
	}
	array, err = rotateArray(array)
	if err != nil {
		return Array[byte]{}, err
	}
	array, err = mirrorOnVerticalAxis(array, 2)
	if err != nil {
		return Array[byte]{}, err
	}

	return array, nil
}

func getColor(h Hash) (Rgb, error) {
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
