package solutions

import (
	"os"
	"strconv"
	"strings"

	"github.com/Capucinoxx/advent-of-code/aoc-2022/common"
)

type Day11 struct {
	monkeysOrder []string

	monkeysPartOne map[string]*Day11Monkey
	monkeysPartTwo map[string]*Day11Monkey
}

type Day11Monkey struct {
	items     []uint64
	operation [2]string
	test      uint64
	true      uint64
	false     uint64
	count     uint64
}

func (d *Day11) Init(input string) {
	file, _ := os.Open(input)
	defer file.Close()

	d.monkeysPartOne = make(map[string]*Day11Monkey)
	d.monkeysPartTwo = make(map[string]*Day11Monkey)

	var lines []string

	common.ReadLines(file, func(line string) {
		if line == "" {
			return
		}
		if len(lines) != 5 {
			lines = append(lines, line)
			return
		}
		lines = append(lines, line)

		m := Day11Monkey{}
		m.items = common.ToUInt64s(common.TrimSplit(lines[1], "  Starting items: ", ", ")...)
		m.test = common.ToUInt64s(strings.TrimPrefix(lines[3], "  Test: divisible by "))[0]
		m.true = common.ToUInt64s(strings.TrimPrefix(lines[4], "    If true: throw to monkey "))[0]
		m.false = common.ToUInt64s(strings.TrimPrefix(lines[5], "    If false: throw to monkey "))[0]

		copy(m.operation[:], common.TrimSplit(lines[2], "  Operation: new = old ", " "))
		d.monkeysOrder = append(d.monkeysOrder, lines[0])

		monkeyPartOne := m
		d.monkeysPartOne[lines[0]] = &monkeyPartOne

		monkeyPartTwo := m
		d.monkeysPartTwo[lines[0]] = &monkeyPartTwo

		lines = nil
	})
}

func (d *Day11) Title() string { return "--- Day 11: Monkey in the Middle ---" }

func (d *Day11) apply(operation [2]string, item uint64) uint64 {
	arg := item
	if operation[1] != "old" {
		v, _ := strconv.Atoi(operation[1])
		arg = uint64(v)
	}

	switch operation[0] {
	case "*":
		return item * arg
	case "+":
		return item + arg
	default:
		return 1
	}
}

func (d *Day11) PartOne() string {
	appendItems := func(monkey string, item uint64) {
		d.monkeysPartOne[monkey].items = append(d.monkeysPartOne[monkey].items, item)
	}

	for i := 0; i < 20; i++ {
		for _, monkeyID := range d.monkeysOrder {
			monkey := d.monkeysPartOne[monkeyID]

			for _, item := range monkey.items {
				worry := d.apply(monkey.operation, item) / 3

				target := monkey.false
				if worry%monkey.test == 0 {
					target = monkey.true
				}

				appendItems(d.monkeysOrder[target], worry)
				monkey.count++
			}
			monkey.items = nil
		}
	}

	max := make([]uint64, 2)
	for _, monkey := range d.monkeysPartOne {
		if monkey.count > max[0] {
			max[1] = max[0]
			max[0] = monkey.count
		} else if monkey.count > max[1] {
			max[1] = monkey.count
		}
	}

	return strconv.Itoa(int(max[0] * max[1]))
}

func (d *Day11) PartTwo() string {
	appendItems := func(monkey string, item uint64) {
		d.monkeysPartTwo[monkey].items = append(d.monkeysPartTwo[monkey].items, item)
	}

	mod := uint64(1)
	for _, monkey := range d.monkeysPartTwo {
		mod *= monkey.test
	}

	for i := 0; i < 10_000; i++ {
		for _, monkeyID := range d.monkeysOrder {
			monkey := d.monkeysPartTwo[monkeyID]

			for _, item := range monkey.items {
				worry := d.apply(monkey.operation, item) % mod

				target := monkey.false
				if worry%monkey.test == 0 {
					target = monkey.true
				}

				appendItems(d.monkeysOrder[target], worry)
				monkey.count++
			}
			monkey.items = nil
		}
	}

	max := make([]uint64, 2)

	for _, monkey := range d.monkeysPartTwo {
		if monkey.count > max[0] {
			max[1] = max[0]
			max[0] = monkey.count
		} else if monkey.count > max[1] {
			max[1] = monkey.count
		}
	}

	return strconv.Itoa(int(max[0] * max[1]))
}
