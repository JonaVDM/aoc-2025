package day02_test

import (
	"testing"

	"github.com/jonavdm/aoc-2025/day02"
)

func TestIsRepeating(t *testing.T) {
	tests := []struct {
		s    string
		want bool
	}{
		{"12", false},
		{"1212", true},
		{"121212", false},
		{"99", true},
		{"9999", true},
		{"9988", false},
		{"999", false},
		{"111", false},
	}
	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			got := day02.IsRepeating(tt.s)
			if got != tt.want {
				t.Errorf("IsRepeating() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsRepeatingALot(t *testing.T) {
	tests := []struct {
		s    string
		want bool
	}{
		{"12", false},
		{"1212", true},
		{"121212", true},
		{"99", true},
		{"9999", true},
		{"9988", false},
		{"999", true},
		{"111", true},
		{"11", true},
		{"22", true},
		{"99", true},
		{"1010", true},
		{"1188511885", true},
		{"222222", true},
		{"824824824", true},
		{"12345", false},
	}
	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			got := day02.IsRepeatingALot(tt.s)
			if got != tt.want {
				t.Errorf("IsRepeatingALot() = %v, want %v", got, tt.want)
			}
		})
	}
}
