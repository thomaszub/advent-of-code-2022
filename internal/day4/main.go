package day4

import (
	"embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var file embed.FS

type section struct {
	lower  int
	higher int
}

func (s section) contains(other section) bool {
	return s.lower <= other.lower && s.higher >= other.higher
}

type pair struct {
	first  section
	second section
}

func Execute() error {
	fmt.Println("*** Day 4 ***")
	bytes, err := file.ReadFile("input.txt")
	if err != nil {
		return err
	}
	content := string(bytes)
	lines := strings.Split(content, "\n")
	pairs, err := getPairs(lines)
	if err != nil {
		return err
	}

	var sumContains int
	for _, pair := range pairs {
		first := pair.first
		second := pair.second
		if first.contains(second) || second.contains(first) {
			sumContains += 1
		}
	}
	fmt.Printf("containing pairs: %d\n", sumContains)

	return nil
}

func getPairs(lines []string) ([]pair, error) {
	var pairs []pair
	for _, line := range lines {
		if line == "" {
			continue
		}
		p, err := getPair(line)
		if err != nil {
			return nil, err
		}
		pairs = append(pairs, p)
	}
	return pairs, nil
}

func getPair(txt string) (pair, error) {
	p := strings.Split(txt, ",")
	if len(p) != 2 {
		return pair{}, fmt.Errorf("%s does not represent a section pair", txt)
	}
	f, err := getSection(p[0])
	if err != nil {
		return pair{}, err
	}
	s, err := getSection(p[1])
	if err != nil {
		return pair{}, err
	}
	return pair{
		first:  f,
		second: s,
	}, err
}

func getSection(txt string) (section, error) {
	s := strings.Split(txt, "-")
	if len(s) != 2 {
		return section{}, fmt.Errorf("%s does not represent a section", txt)
	}
	first, err := strconv.Atoi(s[0])
	if err != nil {
		return section{}, err
	}
	sec, err := strconv.Atoi(s[1])
	if err != nil {
		return section{}, err
	}
	return section{
		lower:  first,
		higher: sec,
	}, err
}
