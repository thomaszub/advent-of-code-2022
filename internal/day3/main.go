package day3

import (
	"embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var file embed.FS

func Execute() error {
	fmt.Println("*** Day 3 ***")
	bytes, err := file.ReadFile("input.txt")
	if err != nil {
		return err
	}
	content := string(bytes)
	lines := strings.Split(content, "\n")

	rucksacks, err := makeRucksacks(lines)
	if err != nil {
		return err
	}

	sum := 0
	for i := 0; i <= len(rucksacks)-3; i += 3 {
		group := rucksacks[i : i+3]
		sameP := findSame(group[0], group[1])
		same := findSame(sameP, group[2])
		if len(same) < 1 {
			return fmt.Errorf("%v does not contain exactly one equal character: %v", group, same)
		}
		prio := runeToPrio(same[0])
		sum += prio
	}

	fmt.Printf("Sum of priorities: %d\n", sum)
	return nil
}

func makeRucksacks(lines []string) ([][]rune, error) {
	var groups [][]rune
	for _, line := range lines {
		if line == "" {
			continue
		}
		groups = append(groups, []rune(line))
	}
	if len(groups)%3 != 0 {
		return nil, fmt.Errorf("%v does not contain a multiple of three", groups)
	}
	return groups, nil
}

func findSame(first, second []int32) []int32 {
	var same []int32
	for _, f := range first {
		for _, s := range second {
			if f == s {
				same = append(same, s)
			}
		}
	}
	return same
}

func runeToPrio(r rune) int {
	prio := r - 96
	if prio <= 0 {
		prio += 58
	}
	return int(prio)
}
