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
		{name: "negative numbers included", left: []int{-7, 1, 9}, right: []int{-2, 5, 8}, merged: []int{-7, -2, 1, 5, 8, 9}},
	}
	for _, tt := range mergeArraysTests {
		t.Run(tt.name, func(t *testing.T) {
			got := Merge(tt.left, tt.right)
			if !Equal(got, tt.merged) {
				t.Errorf("failed in test %q", tt.name)
			}
		})
	}
}

func TestMergeSort(t *testing.T) {
	mergeSortArraysTests := []struct {
		name     string
		source   []int
		expected []int
	}{
		{name: "basic test", source: []int{8, 4, 6, 2, 1, 6, 9}, expected: []int{1, 2, 4, 6, 6, 8, 9}},
		{name: "empty", source: []int{}, expected: []int{}},
		{name: "negative numbers", source: []int{-8, -4, -6, -2, -1, -6, -9}, expected: []int{-9, -8, -6, -6, -4, -2, -1}},
		{name: "both positive and negative", source: []int{-8, 4, -6, 2, -1, 6, 9}, expected: []int{-8, -6, -1, 2, 4, 6, 9}},
	}
	for _, tt := range mergeSortArraysTests {
		t.Run(tt.name, func(t *testing.T) {
			got := MergeSort(tt.source)
			if !Equal(got, tt.expected) {
				t.Errorf("failed in test %q", tt.name)
			}
		})
	}
}
