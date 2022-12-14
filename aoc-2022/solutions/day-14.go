package solutions

import (
	"os"
	"regexp"
	"strconv"

	"github.com/Capucinoxx/advent-of-code/aoc-2022/common"
)

type Day14 struct {
	gridPartOne map[Day14Coord]int
	gridPartTwo map[Day14Coord]int
	minX        int
	minY        int
	maxY        int
	maxX        int
}

type Day14Coord struct {
	x, y int
}

func (d *Day14) Init(input string) {
	file, _ := os.Open(input)
	defer file.Close()

	d.gridPartOne = make(map[Day14Coord]int)
	d.gridPartTwo = make(map[Day14Coord]int)
	d.minY = 99999999999
	d.minX = 99999999999

	common.ReadLines(file, func(line string) {
		positions := common.ToInts(regexp.MustCompile(" -> |,").Split(line, -1)...)

		curr := Day14Coord{positions[0], positions[1]}

		var next Day14Coord
		for i := 2; i < len(positions); i += 2 {
			d.minX = common.Min(d.minX, curr.x)
			d.minY = common.Min(d.minY, curr.y)
			d.maxY = common.Max(d.maxY, curr.y)
			d.maxX = common.Max(d.maxX, curr.x)

			next = Day14Coord{positions[i], positions[i+1]}

			if curr.x != next.x {
				for j := common.Min(curr.x, next.x); j <= common.Max(curr.x, next.x); j++ {
					d.gridPartOne[Day14Coord{j, curr.y}] = -1
					d.gridPartTwo[Day14Coord{j, curr.y}] = -1
				}
			} else {
				for j := common.Min(curr.y, next.y); j <= common.Max(curr.y, next.y); j++ {
					d.gridPartOne[Day14Coord{curr.x, j}] = -1
					d.gridPartTwo[Day14Coord{curr.x, j}] = -1
				}
			}

			curr = next
		}

		d.minX = common.Min(d.minX, curr.x)
		d.minY = common.Min(d.minY, curr.y)
		d.maxY = common.Max(d.maxY, curr.y)
		d.maxX = common.Max(d.maxX, curr.x)
	})
}

func (d *Day14) Title() string { return "--- Day 14: Regolith Reservoir ---" }

func (d *Day14) PartOne() string {
	count := 0

	var produceSand func(coord Day14Coord) bool

	produceSand = func(coord Day14Coord) bool {
		if coord.x < d.minX || coord.x > d.maxX || coord.y > d.maxY {
			return false
		}

		for _, move := range [][2]int{{0, 1}, {-1, 1}, {1, 1}} {
			bellow, ok := d.gridPartOne[Day14Coord{coord.x + move[0], coord.y + move[1]}]
			if !ok || bellow == 0 {
				return produceSand(Day14Coord{coord.x + move[0], coord.y + move[1]})
			}
		}

		d.gridPartOne[coord] = 1
		return true
	}

	for produceSand(Day14Coord{500, 0}) {
		count++
	}

	return strconv.Itoa(count)
}

func (d *Day14) PartTwo() string {
	var produceSand func(coord Day14Coord) bool

	produceSand = func(coord Day14Coord) bool {
		if coord.y == d.maxY+1 {
			d.gridPartTwo[coord] = 1
			return true
		}

		for _, move := range [][2]int{{0, 1}, {-1, 1}, {1, 1}} {
			bellow, ok := d.gridPartTwo[Day14Coord{coord.x + move[0], coord.y + move[1]}]
			if !ok || bellow == 0 {
				return produceSand(Day14Coord{coord.x + move[0], coord.y + move[1]})
			}
		}

		d.gridPartTwo[coord] = 1
		return coord != Day14Coord{500, 0}
	}

	count := 1

	for produceSand(Day14Coord{500, 0}) {
		count++
	}

	return strconv.Itoa(count)
}
