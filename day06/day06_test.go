package day06_test

import (
	"testing"

	"github.com/jonavdm/aoc-2025/day06"
)

func TestPartOne(t *testing.T) {
	tests := []struct {
		name string

		data []string
		want int
	}{
		{
			"Example",
			[]string{
				"123 328  51 64 ",
				" 45 64  387 23 ",
				"  6 98  215 314",
				"*   +   *   +  ",
			},
			4277556,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := day06.PartOne(tt.data)
			if got != tt.want {
				t.Errorf("PartOne() = %v, want %v", got, tt.want)
			}
		})
	}
}
