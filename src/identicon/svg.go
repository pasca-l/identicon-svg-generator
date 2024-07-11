package identicon

import (
	"fmt"
	"io"

	svg "github.com/ajstarks/svgo"
)

func drawIdenticon(w io.Writer, fg Array[byte], color Rgb) error {
	rows, cols, err := fg.shape()
	if err != nil {
		return err
	}

	bgColor, err := newRgb(0xf0, 0xf0, 0xf0)
	if err != nil {
		return err
	}

	s := svg.New(w)
	s.Start(300, 300)
	for r := range rows {
		for c := range cols {
			switch fg[r][c] {
			case 0:
				s.Square(
					c*50+25, r*50+25, 50,
					fmt.Sprintf(`fill="%s"`, bgColor.toColorCode()),
				)
			case 1:
				s.Square(
					c*50+25, r*50+25, 50,
					fmt.Sprintf(`fill="%s"`, color.toColorCode()),
				)
			}
		}
	}
	s.End()

	return nil
}
