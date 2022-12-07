package main

import (
	"fmt"
	"strconv"

	"github.com/Capucinoxx/advent-of-code/aoc-2022/common"
	"github.com/Capucinoxx/advent-of-code/aoc-2022/solutions"
)

var (
	sols = []common.Solution{
		&solutions.Day1{},
		&solutions.Day2{},
		&solutions.Day3{},
		&solutions.Day4{},
		&solutions.Day5{},
		&solutions.Day6{},
		&solutions.Day7{},
	}
)

func main() {
	fmt.Println("Solution for AOC 2022\n")
	for i, s := range sols {
		s.Init("inputs/day-" + strconv.Itoa(i+1) + ".txt")
		common.Print(s)
	}
}
