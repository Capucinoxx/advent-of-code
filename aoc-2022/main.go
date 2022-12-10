package main

import (
	"fmt"
	"strconv"

	"github.com/Capucinoxx/advent-of-code/aoc-2022/common"
	"github.com/Capucinoxx/advent-of-code/aoc-2022/solutions"
)

var (
	sols = []common.Solution{
		&solutions.Day01{},
		&solutions.Day02{},
		&solutions.Day03{},
		&solutions.Day04{},
		&solutions.Day05{},
		&solutions.Day06{},
		&solutions.Day07{},
		&solutions.Day08{},
		&solutions.Day09{},
		&solutions.Day10{},
	}
)

func main() {
	fmt.Println("Solution for AOC 2022\n")
	for i, s := range sols {
		s.Init("inputs/day-" + strconv.Itoa(i+1) + ".txt")
		common.Print(s)
	}
}
