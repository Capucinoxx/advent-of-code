package solutions

import (
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/Capucinoxx/advent-of-code/aoc-2022/common"
)

type Day05Phase string

type Day05 struct {
	phase       Day05Phase
	matrixPart1 [][]uint8
	matrixPart2 [][]uint8
	actions     [][3]int
}

const (
	Day5PhaseCrates    Day05Phase = "crates"
	Day5PhaseColumnsId Day05Phase = "columnsId"
	Day5PhaseActions   Day05Phase = "actions"

	Day5CratesColumns = 9
)

func (d *Day05) Init(input string) {
	file, _ := os.Open(input)
	defer file.Close()

	d.phase = Day5PhaseCrates

	d.matrixPart1 = make([][]uint8, Day5CratesColumns)
	d.matrixPart2 = make([][]uint8, Day5CratesColumns)

	reIsCrate := regexp.MustCompile(".*\\[.*")
	regxActionsGetInts := regexp.MustCompile(`[-]?\d[\d,]*[\.]?[\d{2}]*`)

	common.ReadLines(file, func(line string) {
		switch d.phase {
		case Day5PhaseCrates:
			if !reIsCrate.MatchString(line) {
				d.phase = Day5PhaseColumnsId
				return
			}

			for i := 0; i < len(line)/4+1; i++ {
				crate := line[i*4+1]
				if crate != ' ' {
					d.matrixPart1[i] = append(d.matrixPart1[i], crate)
					d.matrixPart2[i] = append(d.matrixPart2[i], crate)
				}
			}
		case Day5PhaseColumnsId:
			d.phase = Day5PhaseActions
			break
		case Day5PhaseActions:
			ints := regxActionsGetInts.FindAllString(line, -1)
			n, _ := strconv.Atoi(ints[0])
			from, _ := strconv.Atoi(ints[1])
			to, _ := strconv.Atoi(ints[2])

			d.actions = append(d.actions, [3]int{n, from - 1, to - 1})
		}
	})
}

func (d *Day05) Title() string { return "--- Day 5: Supply Stacks ---" }

func (d *Day05) PartOne() string {
	for _, action := range d.actions {
		n, from, to := action[0], action[1], action[2]

		for i := 0; i < n; i++ {
			d.matrixPart1[to] = append([]uint8{d.matrixPart1[from][0]}, d.matrixPart1[to]...)
			d.matrixPart1[from] = d.matrixPart1[from][1:]
		}
	}

	sb := strings.Builder{}
	for _, column := range d.matrixPart1 {
		if len(column) == 0 {
			continue
		}
		sb.WriteByte(column[0])
	}

	return sb.String()
}

func (d *Day05) PartTwo() string {
	for _, action := range d.actions {
		n, from, to := action[0], action[1], action[2]

		movement := make([]uint8, n)
		copy(movement, d.matrixPart2[from][:n])

		d.matrixPart2[to] = append(movement, d.matrixPart2[to]...)
		d.matrixPart2[from] = d.matrixPart2[from][n:]
	}

	sb := strings.Builder{}
	for _, column := range d.matrixPart2 {
		if len(column) == 0 {
			continue
		}
		sb.WriteByte(column[0])
	}

	return sb.String()
}
