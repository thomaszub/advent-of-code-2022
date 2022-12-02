package day1

import (
	"embed"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

//go:embed input.txt
var file embed.FS

func Execute() error {
	fmt.Println("*** Day 1 ***")
	bytes, err := file.ReadFile("input.txt")
	if err != nil {
		return err
	}
	content := string(bytes)
	lines := strings.Split(content, "\n")

	var values []int
	var currentValue int

	for _, lineRaw := range lines {
		line := strings.Trim(lineRaw, " \t")
		if line == "" {
			values = append(values, currentValue)
			currentValue = 0
			continue
		}
		value, err := strconv.Atoi(line)
		if err != nil {
			return err
		}
		currentValue += value
	}

	sort.Sort(sort.Reverse(sort.IntSlice(values)))
	maxValue := values[0] + values[1] + values[2]
	fmt.Printf("The maximum values are: %d\n", maxValue)
	return nil
}
