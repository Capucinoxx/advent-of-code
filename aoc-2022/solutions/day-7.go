package solutions

import (
	"os"
	"strconv"
	"strings"

	"github.com/Capucinoxx/advent-of-code/aoc-2022/common"
)

type Day7 struct {
	sizes map[string]int
}

func (d *Day7) Init(input string) {
	file, _ := os.Open(input)
	defer file.Close()

	path := make([]string, 0)
	d.sizes = make(map[string]int, 0)

	common.ReadLines(file, func(line string) {
		if strings.HasPrefix(line, "$ cd ") {
			dir := line[5:]
			switch dir {
			case "/":
				path = []string{}
			case "..":
				path = path[:len(path)-1]
			default:
				path = append(path, dir)
			}

			return
		}

		els := strings.Split(line, " ")
		if len(els) != 2 {
			return
		}

		size, err := strconv.Atoi(els[0])
		if err != nil {
			return
		}

		for i := 0; i < len(path)+1; i++ {
			pwd := "/" + strings.Join(path[:i], "/")
			if _, ok := d.sizes[pwd]; !ok {
				d.sizes[pwd] = 0
			}
			d.sizes[pwd] += size
		}
	})
}

func (d *Day7) Title() string { return "--- Day 7: No Space Left On Device ---" }

func (d *Day7) PartOne() string {
	sum := 0

	for _, size := range d.sizes {
		if size <= 100_000 {
			sum += size
		}
	}

	return strconv.Itoa(sum)
}

func (d *Day7) PartTwo() string {
	missingSpace := d.sizes["/"] - 40_000_000
	min := int(^uint(0) >> 1)

	for _, size := range d.sizes {
		if size < min && size > missingSpace {
			min = size
		}
	}

	return strconv.Itoa(min)
}
