package day04

import (
	"github.com/jonavdm/aoc-2025/utils"
)

func Run(data []string) [2]any {
	grid := make([][]bool, len(data))
	for y, r := range data {
		grid[y] = make([]bool, len(r))
		for x, c := range r {
			grid[y][x] = c == '@'
		}
	}

	remove := CheckRolls(grid)
	first := len(remove)
	second := len(remove)

	for len(remove) > 0 {
		for _, p := range remove {
			grid[p.Y][p.X] = false
		}

		remove = CheckRolls(grid)
		second += len(remove)
	}

	return [2]any{
		first,
		second,
	}
}

func CheckRolls(grid [][]bool) []utils.Point {
	list := make([]utils.Point, 0)
	for y, row := range grid {
		for x, col := range row {
			if !col {
				continue
			}
			count := 0
			for _, p := range utils.GetAdjacend(x, y, len(grid), len(row)) {
				if grid[p.Y][p.X] {
					count++
				}
			}

			if count < 4 {
				list = append(list, utils.Point{X: x, Y: y})
			}
		}
	}

	return list
}
