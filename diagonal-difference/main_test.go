package main

import "testing"

func TestDiagonalDifference(t *testing.T) {
	tests := []struct {
		name string
		arr  [][]int32
		want int32
	}{
		{
			name: "3x3 matrix",
			arr: [][]int32{
				{11, 2, 4},
				{4, 5, 6},
				{10, 8, -12},
			},
			want: 15,
		},
		{
			name: "2x2 matrix",
			arr: [][]int32{
				{1, 2},
				{3, 4},
			},
			want: 0,
		},
		{
			name: "1x1 matrix",
			arr: [][]int32{
				{5},
			},
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := diagonalDifference(tt.arr)
			if got != tt.want {
				t.Errorf("diagonalDifference() = %v, want %v", got, tt.want)
			}
		})
	}
}
