package identicon

import (
	"fmt"
	"io"

	svg "github.com/ajstarks/svgo"
	"github.com/pasca-l/identicon-generator/utils"
)

func drawIdenticon(w io.Writer, fg utils.Array[byte], color utils.Rgb) error {
	rows, cols, err := fg.Shape()
	if err != nil {
		return err
	}

	bgColor, err := utils.NewRgb(0xf0, 0xf0, 0xf0)
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
					fmt.Sprintf(`fill="%s"`, bgColor.ToColorCode()),
				)
			case 1:
				s.Square(
					c*50+25, r*50+25, 50,
					fmt.Sprintf(`fill="%s"`, color.ToColorCode()),
				)
			}
		}
	}
	s.End()

	return nil
}
