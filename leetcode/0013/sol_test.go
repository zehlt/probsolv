package sol

import (
	"reflect"
	"testing"
)

func TestRomanToInt(t *testing.T) {
	testCases := []struct {
		desc   string
		input  string
		output int
	}{
		{
			desc:   "ex1",
			input:  "III",
			output: 3,
		},
		{
			desc:   "ex2",
			input:  "LVIII",
			output: 58,
		},
		{
			desc:   "ex3",
			input:  "MCMXCIV",
			output: 1994,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := romanToInt(tC.input)

			if !reflect.DeepEqual(got, tC.output) {
				t.Errorf("got %v want %v", got, tC.output)
			}
		})
	}
}
