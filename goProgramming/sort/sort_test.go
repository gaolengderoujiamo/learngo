package sort

import (
	"testing"
)

var tests = []struct{ a1, a2, b []int }{
	{[]int{3, 1, 5, 12, 6}, []int{3, 1, 5, 12, 6}, []int{1, 3, 5, 6, 12}},
	{[]int{23, 12, 34, 0, 23}, []int{23, 12, 34, 0, 23}, []int{0, 12, 23, 23, 34}},
	{[]int{12, 12, 12, 12, 12}, []int{12, 12, 12, 12, 12}, []int{12, 12, 12, 12, 12}},
	{[]int{-1, -4, -3, 12, 6}, []int{-1, -4, -3, 12, 6}, []int{-4, -3, -1, 6, 12}},
}

func TestSelectionSortInt(t *testing.T) {
	for _, tt := range tests {
		if SelectionSortInt(tt.a2); !IntSliceEqual(tt.a2, tt.b) {
			t.Errorf("SelectionSortInt(%v) got %v; expected %v", tt.a1, tt.a2, tt.b)
		}
	}
}

func TestInsertionSort(t *testing.T) {
	for _, tt := range tests {
		if InsertionSort(tt.a2); !IntSliceEqual(tt.a2, tt.b) {
			t.Errorf("InsertionSort(%v) got %v; expected %v", tt.a1, tt.a2, tt.b)
		}
	}
}
