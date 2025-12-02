package day02

import (
	"fmt"
	"strconv"
	"strings"
)

func Run(data []string) [2]any {
	pairs := strings.SplitSeq(data[0], ",")

	simple := 0
	multi := 0
	for pair := range pairs {
		spl := strings.Split(pair, "-")
		a, _ := strconv.Atoi(spl[0])
		b, _ := strconv.Atoi(spl[1])

		for i := a; i <= b; i++ {
			if IsRepeating(fmt.Sprint(i)) {
				simple += i
			}
			if IsRepeatingALot(fmt.Sprint(i)) {
				multi += i
			}
		}
	}

	return [2]any{
		simple,
		multi,
	}
}

func IsRepeating(s string) bool {
	n := len(s) / 2
	return s[n:] == s[:n]
}

func IsRepeatingALot(s string) bool {
	n := len(s)
	for l := 1; l <= n/2; l++ {
		if n%l == 0 {
			sub := s[:l]
			repeated := strings.Repeat(sub, n/l)
			if repeated == s {
				return true
			}
		}
	}
	return false
}
