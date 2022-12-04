package common

import "fmt"

type Solution interface {
	// Title returns the title of the solution
	Title() string

	// Init initializes the solution
	Init(input string)

	// PartOne is the solution to the first part of the puzzle
	PartOne() string

	// PartTwo is the solution to the second part of the puzzle
	PartTwo() string
}

func Print(s Solution) {
	fmt.Println(s.Title())
	fmt.Println("Part One:", s.PartOne())
	fmt.Println("Part Two:", s.PartTwo())
	fmt.Println("--------------------")
	fmt.Println("")
}
