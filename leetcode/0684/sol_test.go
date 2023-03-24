package sol

import (
	"reflect"
	"testing"
)

// Input: edges = [[1,2],[1,3],[2,3]]
// Output: [2,3]
func TestFindRedundantConnection(t *testing.T) {
	edges := [][]int{{1, 2}, {1, 3}, {2, 3}}

	got := findRedundantConnection(edges)
	want := []int{2, 3}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestFindRedundantConnection2(t *testing.T) {
	edges := [][]int{{1, 2}, {2, 3}, {3, 4}, {1, 4}, {1, 5}}

	got := findRedundantConnection(edges)
	want := []int{1, 4}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
