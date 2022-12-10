package solutions

import (
	"os"
	"strconv"
	"strings"

	"github.com/Capucinoxx/advent-of-code/aoc-2022/common"
)

type Day04 struct {
	pairsList [][2][2]int
}

func (d *Day04) Init(input string) {
	file, _ := os.Open(input)
	defer file.Close()

	common.ReadLines(file, func(line string) {
		pairsStr := strings.FieldsFunc(line, func(r rune) bool {
			return strings.ContainsRune(",-", r)
		})

		d.pairsList = append(d.pairsList, [2][2]int{})
		for i, v := range pairsStr {

			d.pairsList[len(d.pairsList)-1][i/2][i%2], _ = strconv.Atoi(v)
		}
	})
}

func (d *Day04) Title() string { return "--- Day 4: Camp Cleanup ---" }

func (d *Day04) PartOne() string {
	sum := 0

	contains := func(a [2]int, b [2]int) bool {
		return (a[0] <= b[0] && a[1] >= b[1]) || (b[0] <= a[0] && b[1] >= a[1])
	}

	for _, pairs := range d.pairsList {
		if contains(pairs[0], pairs[1]) {
			sum++
		}
	}

	return strconv.Itoa(sum)
}

func (d *Day04) PartTwo() string {
	sum := 0

	overlap := func(a [2]int, b [2]int) bool {
		return a[1] >= b[0] && a[1] <= b[1] || a[1] >= b[0] && a[0] <= b[1]
	}

	for _, pairs := range d.pairsList {
		if overlap(pairs[0], pairs[1]) {
			sum++
		}
	}

	return strconv.Itoa(sum)
}
