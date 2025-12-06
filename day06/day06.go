package day06

import (
	"strconv"
	"strings"
)

func Run(data []string) [2]any {
	return [2]any{
		PartOne(data),
		0,
	}
}

func PartOne(data []string) int {
	inst := make([]rune, 0)

	for _, c := range data[len(data)-1] {
		if c == ' ' {
			continue
		}
		inst = append(inst, c)
	}

	nums := make([]int, len(inst))

	for _, l := range data[:len(data)-1] {
		var i int
		for spl := range strings.SplitSeq(l, " ") {
			if spl == "" {
				continue
			}

			n, _ := strconv.Atoi(spl)

			switch inst[i] {
			case '+':
				nums[i] += n
			case '*':
				if nums[i] == 0 {
					nums[i] = n
				} else {
					nums[i] *= n
				}
			default:
				panic("unkown instruction")
			}

			i++
		}
	}

	var total int
	for _, n := range nums {
		total += n
	}

	return total
}
