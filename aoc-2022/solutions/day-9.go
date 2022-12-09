package solutions

import (
	"os"
	"strconv"
	"strings"

	"github.com/Capucinoxx/advent-of-code/aoc-2022/common"
)

type Day9 struct {
	ropePartOne *Day9Rope
	ropePartTwo *Day9Rope
}

type Day9Rope struct {
	length  int
	rope    [][2]int
	visited map[[2]int]struct{}
}

func (d *Day9Rope) Move(direction string) {
	switch direction {
	case "R":
		d.rope[0] = [2]int{d.rope[0][0] + 1, d.rope[0][1]}
	case "L":
		d.rope[0] = [2]int{d.rope[0][0] - 1, d.rope[0][1]}
	case "U":
		d.rope[0] = [2]int{d.rope[0][0], d.rope[0][1] - 1}
	case "D":
		d.rope[0] = [2]int{d.rope[0][0], d.rope[0][1] + 1}
	}

	for i := 1; i < d.length; i++ {
		head := d.rope[i-1]
		tail := d.rope[i]

		if common.Abs(head[0]-tail[0]) <= 1 && common.Abs(head[1]-tail[1]) <= 1 {
			continue
		}

		if tail[0] < head[0] {
			d.rope[i][0]++
		}

		if tail[0] > head[0] {
			d.rope[i][0]--
		}

		if tail[1] < head[1] {
			d.rope[i][1]++
		}

		if tail[1] > head[1] {
			d.rope[i][1]--
		}
	}

	d.visited[d.rope[len(d.rope)-1]] = struct{}{}
}

func (d *Day9) Init(input string) {
	file, _ := os.Open(input)
	defer file.Close()

	d.ropePartOne = &Day9Rope{
		length:  2,
		rope:    make([][2]int, 2),
		visited: make(map[[2]int]struct{}),
	}
	d.ropePartTwo = &Day9Rope{
		length:  10,
		rope:    make([][2]int, 10),
		visited: make(map[[2]int]struct{}),
	}

	common.ReadLines(file, func(line string) {
		els := strings.Split(line, " ")

		movement := els[0]
		count, _ := strconv.Atoi(els[1])

		for i := 0; i < count; i++ {
			d.ropePartOne.Move(movement)
			d.ropePartTwo.Move(movement)
		}
	})
}

func (d *Day9) Title() string { return "--- Day 9: Rope Bridge ---" }

func (d *Day9) PartOne() string {
	return strconv.Itoa(len(d.ropePartOne.visited))
}

func (d *Day9) PartTwo() string {
	return strconv.Itoa(len(d.ropePartTwo.visited))
}
