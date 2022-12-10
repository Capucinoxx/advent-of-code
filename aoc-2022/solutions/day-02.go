package solutions

import (
	"os"
	"strconv"
	"strings"

	"github.com/Capucinoxx/advent-of-code/aoc-2022/common"
)

type Day02 struct {
	sumPartOne int
	sumPartTwo int
}

func (d *Day02) Init(input string) {
	file, _ := os.Open(input)
	defer file.Close()

	pointsPartA := map[string]map[string]int{
		"A": {"X": 4, "Y": 8, "Z": 3},
		"B": {"X": 1, "Y": 5, "Z": 9},
		"C": {"X": 7, "Y": 2, "Z": 6},
	}

	pointsPartB := map[string]map[string]int{
		"A": {"X": 3, "Y": 4, "Z": 8},
		"B": {"X": 1, "Y": 5, "Z": 9},
		"C": {"X": 2, "Y": 6, "Z": 7},
	}

	common.ReadLines(file, func(line string) {
		el := strings.Split(line, " ")

		d.sumPartOne += pointsPartA[el[0]][el[1]]
		d.sumPartTwo += pointsPartB[el[0]][el[1]]
	})
}

func (d *Day02) Title() string { return "--- Day 2: Rock Paper Scissors ---" }

func (d *Day02) PartOne() string { return strconv.Itoa(d.sumPartOne) }

func (d *Day02) PartTwo() string { return strconv.Itoa(d.sumPartTwo) }
