package solutions

import (
	"os"
	"strconv"

	"github.com/Capucinoxx/advent-of-code/aoc-2022/common"
)

type Day12 struct {
	grid  [][]int
	start [2]int
	end   [2]int
}

func (d *Day12) Init(input string) {
	file, _ := os.Open(input)
	defer file.Close()

	d.grid = [][]int{}

	common.ReadLines(file, func(line string) {
		d.grid = append(d.grid, make([]int, len(line)))
		for x, ch := range line {
			switch ch {
			case 'S':
				d.start = [2]int{x, len(d.grid) - 1}
				d.grid[len(d.grid)-1][x] = int('a')
				break
			case 'E':
				d.end = [2]int{x, len(d.grid) - 1}
				d.grid[len(d.grid)-1][x] = int('z')
				break
			default:
				d.grid[len(d.grid)-1][x] = int(ch)
			}
		}
	})
}

func (d *Day12) Title() string { return "--- Day 12: Hill Climbing Algorithm ---" }

func (d *Day12) PartOne() string {
	return strconv.Itoa(d.bfs(d.start))
}

func (d *Day12) PartTwo() string {
	best := 1_000_000_000

	for i := 0; i < len(d.grid); i++ {
		for j := 0; j < len(d.grid[i]); j++ {
			if d.grid[i][j] == int('a') {
				newBest := d.bfs([2]int{j, i})
				if newBest < best {
					best = newBest
				}
			}
		}
	}

	return strconv.Itoa(best)
}

func (d *Day12) neighbors(x, y int) [][2]int {
	var neighbors [][2]int

	height := d.grid[y][x]

	for _, n := range [][2]int{{x - 1, y}, {x + 1, y}, {x, y - 1}, {x, y + 1}} {
		if !(n[0] >= 0 && n[0] < len(d.grid[0]) && n[1] >= 0 && n[1] < len(d.grid)) {
			continue
		}

		if d.grid[n[1]][n[0]]-height <= 1 {
			neighbors = append(neighbors, n)
		}
	}

	return neighbors
}

func (d *Day12) bfs(start [2]int) int {
	queue := [][2]int{start}
	visited := map[[2]int]int{start: 0}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current == d.end {
			return visited[current]
		}

		for _, n := range d.neighbors(current[0], current[1]) {
			if _, ok := visited[n]; !ok {
				visited[n] = visited[current] + 1
				queue = append(queue, n)
			}
		}
	}
	return 1_000_000_000
}
