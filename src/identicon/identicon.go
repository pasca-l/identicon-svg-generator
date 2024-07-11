package identicon

import (
	"fmt"
	"os"

	"github.com/pasca-l/identicon-generator/utils"
)

func GenerateIdenticon(accountId string) error {
	Hash := utils.GenerateMd5Hash(accountId)

	foreground, err := getForeground(Hash)
	if err != nil {
		return err
	}
	color, err := getColor(Hash)
	if err != nil {
		return err
	}

	fmt.Println(foreground, color)
	drawIdenticon(os.Stdout, foreground, color)

	return nil
}

func getForeground(h utils.Hash) (utils.Array[byte], error) {
	// get foreground binary from 4 bit (nibble) parity
	parity := make([]byte, 0, 15)
	for _, nibble := range h.Hash {
		// odd (background) = 0, even (foreground) = 1
		parity = append(parity, ^nibble&1)
	}

	fg, err := utils.RearrangeForIdenticon(parity)
	if err != nil {
		return utils.Array[byte]{}, err
	}

	return fg, nil
}

func getColor(h utils.Hash) (utils.Rgb, error) {
	// get hsl values from last 7 hexadecimals, mapped as "HHHSSLL"
	hhh := h.Hash[len(h.Hash)-7 : len(h.Hash)-4]
	ss := h.Hash[len(h.Hash)-4 : len(h.Hash)-2]
	ll := h.Hash[len(h.Hash)-2:]

	// remapped to 0 <= hue < 360
	hue := utils.ConvertBytesToPercentage(hhh) * 360
	if hue == 360 {
		hue = 0
	}
	// remapped to 0 <= saturation <= 0.2, and subtracted from max 0.65
	saturation := 0.65 - utils.ConvertBytesToPercentage(ss)*0.2
	// remapped to 0 <= lightness <= 0.2, and subtracted from max 0.75
	lightness := 0.75 - utils.ConvertBytesToPercentage(ll)*0.2

	hsl, err := utils.NewHsl(hue, saturation, lightness)
	if err != nil {
		return utils.Rgb{}, err
	}
	rgba, err := hsl.ConvertHslToRgb()
	if err != nil {
		return utils.Rgb{}, err
	}

	return rgba, nil
}
