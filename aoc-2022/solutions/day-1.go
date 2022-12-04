package solutions

import (
	"os"
	"sort"
	"strconv"

	"github.com/Capucinoxx/advent-of-code/aoc-2022/common"
)

type Day1 struct {
	calories []int
}

func (d *Day1) Init(input string) {
	file, _ := os.Open(input)
	defer file.Close()

	d.calories = make([]int, 1)

	common.ReadLines(file, func(line string) {
		if line == "" {
			d.calories = append(d.calories, 0)
			return
		}

		v, _ := strconv.Atoi(line)
		d.calories[len(d.calories)-1] += v
	})

	sort.Ints(d.calories)
}

func (d *Day1) Title() string { return "--- Day 1: Calorie Counting ---" }

func (d *Day1) PartOne() string {
	return strconv.Itoa(d.calories[len(d.calories)-1])
}

func (d *Day1) PartTwo() string {
	res := common.Sum(d.calories[len(d.calories)-3:])
	return strconv.Itoa(res)
}
