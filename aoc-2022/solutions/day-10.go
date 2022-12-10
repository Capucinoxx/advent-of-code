package solutions

import (
	"os"
	"strconv"
	"strings"

	"github.com/Capucinoxx/advent-of-code/aoc-2022/common"
)

type Day10 struct {
	regVal []int
}

func (d *Day10) Init(input string) {
	file, _ := os.Open(input)
	defer file.Close()

	d.regVal = make([]int, 1, 221)
	d.regVal[0] = 1

	common.ReadLines(file, func(line string) {
		els := strings.Split(line, " ")
		instruction := els[0]

		d.regVal = append(d.regVal, d.regVal[len(d.regVal)-1])

		if instruction == "addx" {
			x, _ := strconv.Atoi(els[1])
			d.regVal = append(d.regVal, d.regVal[len(d.regVal)-1]+x)
		}
	})
}

func (d *Day10) Title() string { return "--- Day 10: Cathode-Ray Tube ---" }

func (d *Day10) PartOne() string {
	signalStrength := 0

	for _, cycle := range []int{20, 60, 100, 140, 180, 220} {
		signalStrength += d.regVal[cycle-1] * cycle
	}

	return strconv.Itoa(signalStrength)
}

func (d *Day10) PartTwo() string {
	wide := 40
	high := 6

	lines := make([]string, high)

	for i := 0; i < high; i++ {
		for j := 0; j < wide; j++ {
			x := d.regVal[i*wide+j]
			if x-1 <= j && j <= x+1 {
				lines[i] += "#"
			} else {
				lines[i] += "."
			}
		}
	}

	return strings.Join(lines, "\n")
}
