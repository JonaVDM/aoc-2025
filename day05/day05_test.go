package day05_test

import (
	"github.com/jonavdm/aoc-2025/day05"
	"testing"
)

func TestSolver_IdCount(t *testing.T) {
	tests := []struct {
		name string
		in   [][2]int
		want int
	}{
		{
			"Example",
			[][2]int{
				{3, 5},
				{10, 14},
				{16, 20},
				{12, 18},
			},
			14,
		},
		{
			"Inside",
			[][2]int{
				{10, 20},
				{15, 17},
			},
			11,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := day05.Solver{
				Ranges: tt.in,
			}
			got := s.IdCount()

			if got != tt.want {
				t.Errorf("IdCount() = %v, want %v", got, tt.want)
			}
		})
	}
}
