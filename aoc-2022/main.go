package main

import (
	"fmt"
	"strconv"

	"github.com/Capucinoxx/advent-of-code/aoc-2022/common"
)

var (
	solutions = []common.Solution{}
)

func main() {
	fmt.Println("Solution for AOC 2022")
	for i, s := range solutions {
		s.Init("inputs/day-" + strconv.Itoa(i+1) + ".txt")
		common.Print(s)
	}
}
