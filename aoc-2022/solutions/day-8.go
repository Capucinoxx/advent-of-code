package solutions

import (
	"os"
	"strconv"

	"github.com/Capucinoxx/advent-of-code/aoc-2022/common"
)

type Day8 struct {
	trees [][]Day8Tree
}

type Day8Tree struct {
	height      int
	visible     bool
	scenicScore int
}

func (d *Day8) Init(input string) {
	file, _ := os.Open(input)
	defer file.Close()

	common.ReadLines(file, func(line string) {
		d.trees = append(d.trees, make([]Day8Tree, len(line)))

		for i, ch := range line {
			height, err := strconv.Atoi(string(ch))
			if err != nil {
				continue
			}

			d.trees[len(d.trees)-1][i] = Day8Tree{height: height}
		}
	})

	for row := 0; row < len(d.trees); row++ {
		for col := 0; col < len(d.trees[row]); col++ {
			height := d.trees[row][col].height
			visible := [4]bool{true, true, true, true}
			scenicScores := [4]int{0, 0, 0, 0}

			for x := row - 1; x >= 0; x-- {
				scenicScores[0]++
				if d.trees[x][col].height >= height {
					visible[0] = false
					break
				}
			}

			for x := row + 1; x < len(d.trees); x++ {
				scenicScores[1]++
				if d.trees[x][col].height >= height {
					visible[1] = false
					break
				}
			}

			for x := col - 1; x >= 0; x-- {
				scenicScores[2]++
				if d.trees[row][x].height >= height {
					visible[2] = false
					break
				}
			}

			for x := col + 1; x < len(d.trees[row]); x++ {
				scenicScores[3]++
				if d.trees[row][x].height >= height {
					visible[3] = false
					break
				}
			}

			if visible[0] || visible[1] || visible[2] || visible[3] {
				d.trees[row][col].visible = true
			}
			d.trees[row][col].scenicScore = common.Mul(scenicScores[:])
		}
	}
}

func (d *Day8) Title() string { return "--- Day 8: Treetop Tree House ---" }

func (d *Day8) PartOne() string {
	count := 0

	for _, row := range d.trees {
		for _, tree := range row {
			if tree.visible {
				count++
			}
		}
	}

	return strconv.Itoa(count)
}

func (d *Day8) PartTwo() string {
	max := 0

	for _, row := range d.trees {
		for _, tree := range row {
			if tree.scenicScore > max {
				max = tree.scenicScore
			}
		}
	}

	return strconv.Itoa(max)
}
