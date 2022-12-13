package solutions

import (
	"os"
	"sort"
	"strconv"

	"github.com/Capucinoxx/advent-of-code/aoc-2022/common"
)

type Day13 struct {
	packets [][]Day13Value
}

type Day13Value struct {
	val   int
	items []Day13Value
}

func (d Day13Value) isInt() bool {
	return d.items == nil
}

func (d *Day13) Init(input string) {
	file, _ := os.Open(input)
	defer file.Close()

	d.packets = append(d.packets, []Day13Value{})

	common.ReadLines(file, func(line string) {
		if line != "" {
			d.packets[len(d.packets)-1] = append(d.packets[len(d.packets)-1], parseValue(line))
		} else {
			d.packets = append(d.packets, []Day13Value{})
		}
	})
}

func parseValue(line string) Day13Value {
	if v, err := strconv.Atoi(line); err == nil {
		return Day13Value{val: v}
	}

	stripped := line[1 : len(line)-1]

	nesting := 0
	var parts []string
	last := 0

	for i := range stripped {
		switch stripped[i] {
		case '[':
			nesting++
		case ']':
			nesting--
		case ',':
			if nesting == 0 {
				parts = append(parts, stripped[last:i])
				last = i + 1
			}
		}
	}
	parts = append(parts, stripped[last:])

	d := &Day13Value{items: make([]Day13Value, 0, len(parts))}
	for _, part := range parts {
		if len(part) > 0 {
			d.items = append(d.items, parseValue(part))
		}
	}

	return *d
}

func comp(v1, v2 Day13Value) int {
	if v1.isInt() && v2.isInt() {
		return v1.val - v2.val
	}

	if !v1.isInt() && !v2.isInt() {
		for i := 0; i < common.Max(len(v1.items), len(v2.items)); i++ {
			if i >= len(v1.items) {
				return -1
			}

			if i >= len(v2.items) {
				return 1
			}

			res := comp(v1.items[i], v2.items[i])
			if res != 0 {
				return res
			}

		}

		return 0
	}

	if v1.isInt() {
		v1 = Day13Value{items: []Day13Value{v1}}
	} else {
		v2 = Day13Value{items: []Day13Value{v2}}
	}

	return comp(v1, v2)
}

func (d *Day13) Title() string { return "--- Day 13: Distress Signal ---" }

func (d *Day13) PartOne() string {
	count := 0

	for i, packet := range d.packets {
		if comp(packet[0], packet[1]) < 0 {
			count += i + 1
		}
	}

	return strconv.Itoa(count)
}

func (d *Day13) PartTwo() string {
	var packets []Day13Value

	for _, p := range d.packets {
		packets = append(packets, p[0], p[1])
	}

	packets = append(packets, Day13Value{items: []Day13Value{{val: 2}}})
	packets = append(packets, Day13Value{items: []Day13Value{{val: 6}}})

	sort.Slice(packets, func(i, j int) bool {
		return comp(packets[i], packets[j]) < 0
	})

	mul := 1
	for i, p := range packets {
		if len(p.items) == 1 && (p.items[0].val == 2 || p.items[0].val == 6) {
			mul *= i + 1
		}
	}

	return strconv.Itoa(mul)
}
