package day01

import (
	"strconv"
)

func Run(data []string) [2]any {
	return [2]any{
		solveA(data),
		solveB(data),
	}
}

func solveA(data []string) int {
	point := 50
	var count int

	for _, l := range data {
		n, err := strconv.Atoi(l[1:])
		if err != nil {
			panic(err)
		}

		if l[0] == 'R' {
			point += n
		} else {
			point -= n
		}

		point %= 100
		if point < 0 {
			point += 100
		}

		if point == 0 {
			count++
		}
	}

	return count
}

func solveB(data []string) int {
	point := 50
	var count int

	for _, l := range data {
		n, err := strconv.Atoi(l[1:])
		if err != nil {
			panic(err)
		}

		for range n {
			if l[0] == 'R' {
				point++
			} else {
				point--
			}

			point %= 100
			if point < 0 {
				point += 100
			}

			if point == 0 {
				count++
			}
		}
	}

	return count
}
