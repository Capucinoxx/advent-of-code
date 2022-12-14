package solutions

import (
	"os"
	"strconv"

	"github.com/Capucinoxx/advent-of-code/aoc-2022/common"
)

type Day06 struct {
	postPartOne int
	postPartTwo int
}

type Day06Character struct {
	Count    int
	Position int
}

func (d *Day06) Init(input string) {
	file, _ := os.Open(input)
	defer file.Close()

	hash := make(map[int32]Day06Character)
	common.ReadLines(file, func(line string) {
		count := 0
		for i, ch := range line {
			count++

			if old, ok := hash[ch]; !ok || i-old.Position > count {
				hash[ch] = Day06Character{Count: count, Position: i}
			} else {
				count = i - hash[ch].Position
				hash[ch] = Day06Character{Count: count, Position: i}
			}

			if count == 4 && d.postPartOne == 0 {
				d.postPartOne = i + 1
			}

			if count == 14 && d.postPartTwo == 0 {
				d.postPartTwo = i + 1
			}
		}
	})
}

func (d *Day06) Title() string { return "--- Day 6: Tuning Trouble ---" }

func (d *Day06) PartOne() string {
	return strconv.Itoa(d.postPartOne)
}

func (d *Day06) PartTwo() string {
	return strconv.Itoa(d.postPartTwo)
}
