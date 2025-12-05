package day05

import (
	"slices"
	"strconv"
	"strings"
)

func Run(data []string) [2]any {
	ranges := make([][2]int, 0)
	var ids []string

	for i, l := range data {
		if l == "" {
			ids = data[i+1:]
			break
		}

		spl := strings.Split(l, "-")
		st, _ := strconv.Atoi(spl[0])
		en, _ := strconv.Atoi(spl[1])

		ranges = append(ranges, [2]int{st, en})
	}

	solver := Solver{
		Ranges: ranges,
	}

	count := 0

	for _, id := range ids {
		if solver.ValidId(id) {
			count++
		}
	}

	return [2]any{
		count,
		solver.IdCount(),
	}
}

type Solver struct {
	Ranges [][2]int
}

func (s *Solver) ValidId(id string) bool {
	n, err := strconv.Atoi(id)
	if err != nil {
		panic(err)
	}

	for _, r := range s.Ranges {
		if r[0] <= n && n <= r[1] {
			return true
		}
	}

	return false
}

func (s *Solver) IdCount() int {
	sorted := make([][2]int, len(s.Ranges))
	copy(sorted, s.Ranges)
	slices.SortFunc(sorted, func(a, b [2]int) int {
		return a[0] - b[0]
	})

	index := 0
	for index < len(sorted)-1 {
		if sorted[index][1] >= sorted[index+1][0] {
			if sorted[index][1] < sorted[index+1][1] {
				sorted[index][1] = sorted[index+1][1]
			}
			sorted = append(sorted[:index+1], sorted[index+2:]...)
		} else {
			index++
		}
	}

	count := 0
	for _, r := range sorted {
		count += r[1] - r[0] + 1
	}

	return count
}
