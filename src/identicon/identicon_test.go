package identicon_test

import (
	"reflect"
	"testing"

	"github.com/pasca-l/identicon-svg-generator/identicon"
	"github.com/pasca-l/identicon-svg-generator/utils"
)

func TestGenerateIdenticon(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want identicon.Identicon
	}{
		{
			name: "test for user 'pasca-l'",
			arg:  "pasca-l",
			want: identicon.Identicon{
				Foreground: utils.Array[byte]{
					[]byte{1, 1, 1, 1, 1},
					[]byte{1, 0, 0, 0, 1},
					[]byte{1, 1, 0, 1, 1},
					[]byte{0, 1, 1, 1, 0},
					[]byte{0, 1, 0, 1, 0},
				},
				Color: utils.Rgb{
					R: 214,
					G: 182,
					B: 130,
				},
			},
		},
		{
			name: "test for user 'github'",
			arg:  "github",
			want: identicon.Identicon{
				Foreground: utils.Array[byte]{
					[]byte{1, 1, 1, 1, 1},
					[]byte{0, 1, 1, 1, 0},
					[]byte{1, 1, 1, 1, 1},
					[]byte{0, 1, 0, 1, 0},
					[]byte{1, 1, 1, 1, 1},
				},
				Color: utils.Rgb{
					R: 218,
					G: 189,
					B: 145,
				},
			},
		},
	}

	for _, test := range tests {
		// shadowing variable to resolve "loop variable captured by func literal",
		// necessary for parallel testing
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			got, err := identicon.GenerateIdenticon(test.arg)
			if err != nil {
				t.Errorf("error in generating identicon: %v", err)
			}
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("test failed, got: %+v, want: %+v", got, test.want)
			}
		})
	}
}
