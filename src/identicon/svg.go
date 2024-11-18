package identicon

import (
	"fmt"
	"io"

	svg "github.com/ajstarks/svgo"
	"github.com/pasca-l/identicon-svg-generator/utils"
)

func DrawIdenticon(w io.Writer, icon Identicon) error {
	rows, cols, err := icon.Foreground.Shape()
	if err != nil {
		return err
	}

	bgColor, err := utils.NewRgb(0xf0, 0xf0, 0xf0)
	if err != nil {
		return err
	}

	s := svg.New(w)
	s.Start(300, 300)
	s.Square(0, 0, 300, fmt.Sprintf(`fill="%s"`, bgColor.ToColorCode()))
	for r := range rows {
		for c := range cols {
			if icon.Foreground[r][c] == 1 {
				s.Square(
					c*50+25, r*50+25, 50,
					fmt.Sprintf(`fill="%s"`, icon.Color.ToColorCode()),
					fmt.Sprintf(`stroke="%s"`, icon.Color.ToColorCode()),
				)
			}
		}
	}
	s.End()

	return nil
}
