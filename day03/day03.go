package day03

import "github.com/jonavdm/aoc-2025/utils"

func Run(data []string) [2]any {
	total := 0
	test := 0
	for _, l := range data {
		nums := make([]int, len(l))
		for i, r := range l {
			nums[i] = int(r - '0')
		}

		total += FindBestNumber(nums, 2)
		test += FindBestNumber(nums, 12)
	}

	return [2]any{
		total,
		test,
	}
}

func FindBestNumber(nums []int, amount int) int {
	if amount == 0 {
		return 0
	}

	high, index := 0, 0

	for i := 0; i < len(nums)-amount+1; i++ {
		n := nums[i]
		if n > high {
			high = n
			index = i
		}
	}

	return high*utils.PowInt(10, amount-1) + FindBestNumber(nums[index+1:], amount-1)
}
