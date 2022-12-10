package common

import (
	"fmt"
	"strings"
)

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

// Print prints the solution
func Print(s Solution) {
	var sb strings.Builder

	// title
	sb.WriteString(s.Title())
	sb.WriteString("\n")

	// part one
	AppendPart(&sb, "Part One: ", s.PartOne())

	// part two
	AppendPart(&sb, "Part Two: ", s.PartTwo())

	// footer
	sb.WriteString("--------------------\n\n")

	fmt.Print(sb.String())
}

// AppendPart appends the part to the string builder. If the solution has
// multiple lines, it will be indented by the length of the part string.
func AppendPart(sb *strings.Builder, part string, solution string) {
	lines := strings.Split(solution, "\n")
	if len(lines) == 1 {
		sb.WriteString(part)
		sb.WriteString(solution)
		sb.WriteString("\n")
		return
	}

	sb.WriteString(part)
	sb.WriteString(lines[0])
	sb.WriteString("\n")

	space := strings.Repeat(" ", len(part))
	for _, line := range lines[1:] {
		sb.WriteString(space)
		sb.WriteString(line)
		sb.WriteString("\n")
	}
}
