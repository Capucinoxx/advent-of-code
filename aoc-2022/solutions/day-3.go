package solutions

import (
	"os"
	"strconv"

	"github.com/Capucinoxx/advent-of-code/aoc-2022/common"
)

type Day3 struct {
	items  []string
	groups [][3]string
}

func (d *Day3) Init(input string) {
	file, _ := os.Open(input)
	defer file.Close()

	d.groups = make([][3]string, 1)

	idx := 0
	common.ReadLines(file, func(line string) {
		d.items = append(d.items, line[0:len(line)/2], line[len(line)/2:])

		d.groups[len(d.groups)-1][idx] = line
		idx++
		if idx == 3 {
			idx = 0
			d.groups = append(d.groups, [3]string{})
		}
	})
}

func (d *Day3) Title() string { return "--- Day 3: Rucksack Reorganization ---" }

func (d *Day3) PartOne() string {
	sum := 0

	for i := 0; i < len(d.items); i += 2 {
		sharedItems := itemsIntersection([]byte(d.items[i]), []byte(d.items[i+1]))
		for _, ch := range sharedItems {
			sum += d.toScore(ch)
		}
	}

	return strconv.Itoa(sum)
}

func (d *Day3) PartTwo() string {
	sum := 0

	for _, group := range d.groups {
		sharedItems := itemsIntersection([]byte(group[0]), []byte(group[1]), []byte(group[2]))
		for _, ch := range sharedItems {
			sum += d.toScore(ch)
		}
	}

	return strconv.Itoa(sum)
}

func itemsIntersection[T comparable](s ...[]T) (inter []T) {
	length := len(s)

	hash := make(map[T]int)
	for _, ss := range s {
		uniq := make(map[T]struct{})
		for _, ch := range ss {
			uniq[ch] = struct{}{}
		}
		for ch := range uniq {
			hash[ch]++
		}
	}

	for ch, count := range hash {
		if count == length {
			inter = append(inter, ch)
		}
	}

	return
}

func (d *Day3) toScore(ch byte) int {
	if ch <= 'Z' {
		return int(ch - 'A' + 27)
	}
	return int(ch - 'a' + 1)
}
