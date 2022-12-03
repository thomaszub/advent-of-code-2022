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

	sum := 0

	for _, line := range lines {
		if line == "" {
			continue
		}
		runes := []rune(line)
		if len(runes)%2 != 0 {
			return fmt.Errorf("the line %s contains an uneven count of characters", line)
		}
		splitLen := len(runes) / 2
		same, err := findSame(runes[:splitLen], runes[splitLen:])
		if err != nil {
			return err
		}
		sum += runeToPrio(same)
	}

	fmt.Printf("Sum of priorities: %d\n", sum)
	return nil
}

func findSame(first, second []int32) (int32, error) {
	for _, f := range first {
		for _, s := range second {
			if f == s {
				return s, nil
			}
		}
	}

	return 0, fmt.Errorf("%v and %v do not contain a same number", first, second)
}

func runeToPrio(r rune) int {
	prio := r - 96
	if prio <= 0 {
		prio += 58
	}
	return int(prio)
}
