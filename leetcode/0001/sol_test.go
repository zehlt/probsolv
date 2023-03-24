package sol

import (
	"reflect"
	"testing"
)

func TestTwoSums(t *testing.T) {
	testCases := []struct {
		desc   string
		nums   []int
		target int
		want   []int
	}{
		{
			desc:   "ex1",
			nums:   []int{2, 7, 11, 15},
			target: 9,
			want:   []int{0, 1},
		},
		{
			desc:   "ex2",
			nums:   []int{3, 2, 4},
			target: 6,
			want:   []int{1, 2},
		},
		{
			desc:   "ex3",
			nums:   []int{3, 3},
			target: 6,
			want:   []int{0, 1},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := twoSum(tC.nums, tC.target)

			if !reflect.DeepEqual(got, tC.want) {
				t.Errorf("got %v want %v", got, tC.want)
			}
		})
	}
}
