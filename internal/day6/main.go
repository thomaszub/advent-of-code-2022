package day6

import (
	"embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var file embed.FS

const markerLength = 4

func Execute() error {
	fmt.Println("*** Day 6 ***")
	bytes, err := file.ReadFile("input.txt")
	if err != nil {
		return err
	}
	content := strings.Trim(string(bytes), " \t\n")
	fmt.Println(content)
	runes := []rune(content)
	if len(runes) < markerLength {
		return fmt.Errorf("content %q is too short", content)
	}

	var numChars int
	for pos := markerLength; pos <= len(runes); pos++ {
		marker := runes[pos-markerLength : pos]
		if checkMarker(marker) {
			numChars = pos
			break
		}
	}
	fmt.Printf("The last marker position is: %d\n", numChars)
	return nil
}

func checkMarker(marker []rune) bool {
	mp := map[rune]bool{}
	for _, m := range marker {
		mp[m] = true
	}
	return len(mp) == markerLength
}
