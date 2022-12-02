package day2

import (
	"embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var file embed.FS

type pair struct {
	first  string
	second string
}

type strategy map[string]int

const (
	rock    = 1
	paper   = 2
	scissor = 3
)

const (
	lost = 0
	draw = 3
	won  = 6
)

var opponentStrategy = strategy{"A": rock, "B": paper, "C": scissor}

// first opponent
var winMatrix = map[int]map[int]int{
	rock: {
		lost: scissor,
		draw: rock,
		won:  paper,
	},
	paper: {
		lost: rock,
		draw: paper,
		won:  scissor,
	},
	scissor: {
		lost: paper,
		draw: scissor,
		won:  rock,
	},
}

func Execute() error {
	fmt.Println("*** Day 2 ***")
	bytes, err := file.ReadFile("input.txt")
	if err != nil {
		return err
	}
	content := string(bytes)
	lines := strings.Split(content, "\n")

	var pairs []pair
	for _, line := range lines {
		if line == "" {
			continue
		}
		splitted := strings.Split(line, " ")
		if len(splitted) != 2 {
			return fmt.Errorf("%s does not contain two strings", line)
		}
		pairs = append(pairs, pair{first: splitted[0], second: splitted[1]})
	}

	strat := strategy{"X": lost, "Y": draw, "Z": won}

	points, err := play(pairs, strat)
	if err != nil {
		return err
	}
	fmt.Printf("Strategy: %v, points: %d\n", strat, points)

	return nil
}

func play(pairs []pair, strat strategy) (int, error) {
	var points int
	for _, pair := range pairs {
		oppC, ok := opponentStrategy[pair.first]
		if !ok {
			return 0, fmt.Errorf("opponent: %s is not a known character", pair.first)
		}
		ownResult, ok := strat[pair.second]
		if !ok {
			return 0, fmt.Errorf("own: %s is not a known character", pair.second)
		}

		result, ok := winMatrix[oppC][ownResult]
		if !ok {
			return 0, fmt.Errorf("opponent %d or/and own %d are not valid choices", oppC, ownResult)
		}
		points += ownResult + result
	}
	return points, nil
}
