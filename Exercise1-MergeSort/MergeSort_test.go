package Exercise1_MergeSort

import (
	"testing"
)

func Equal(got, want []int) bool {
	if len(got) != len(want) {
		return false
	}
	for i, g := range got {
		if g != want[i] {
			return false
		}
	}
	return true
}
func TestMerge(t *testing.T) {
	mergeArraysTests := []struct {
		name   string
		left   []int
		right  []int
		merged []int
	}{
		{name: "basic test", left: []int{1, 7, 9}, right: []int{2, 5, 8}, merged: []int{1, 2, 5, 7, 8, 9}},
		{name: "left empty", left: []int{}, right: []int{1, 5, 7}, merged: []int{1, 5, 7}},
		{name: "right empty", left: []int{1, 5, 7}, right: []int{}, merged: []int{1, 5, 7}},
		{name: "duplicate numbers", left: []int{1, 1, 4, 6}, right: []int{2, 4, 5, 6}, merged: []int{1, 1, 2, 4, 4, 5, 6, 6}},
	}
	for _, tt := range mergeArraysTests {
		t.Run(tt.name, func(t *testing.T) {
			got := Merge(tt.left, tt.right)
			if !Equal(got, tt.merged) {
				t.Errorf("failed int test %q", tt.name)
			}
		})
	}
}
