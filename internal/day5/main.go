package day5

import (
	"embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var file embed.FS

type move struct {
	count int
	from  int
	to    int
}

type stack struct {
	num        int
	containers []string
}

func Execute() error {
	fmt.Println("*** Day 5 ***")
	bytes, err := file.ReadFile("input.txt")
	if err != nil {
		return err
	}
	content := string(bytes)
	lines := strings.Split(content, "\n")
	var secondPart int
	for _, line := range lines {
		if line == "" {
			break
		}
		secondPart += 1
	}

	stacks := makeStacks(lines[:secondPart])
	moves, err := makeMoves(lines[secondPart:])
	if err != nil {
		return err
	}

	for _, move := range moves {
		stackFrom := &stacks[move.from-1]
		stackTo := &stacks[move.to-1]
		cons := stackFrom.containers
		for id := 1; id <= move.count; id++ {
			stackTo.containers = append(stackTo.containers, cons[len(cons)-id])
		}
		stackFrom.containers = stackFrom.containers[:len(stackFrom.containers)-move.count]
	}

	var top []string
	for _, stack := range stacks {
		cons := stack.containers
		top = append(top, cons[len(cons)-1])
	}
	fmt.Println(strings.Join(top, ""))
	return nil
}

func makeStacks(lines []string) []stack {
	var contRow [][]string
	var contNums []int
	for _, line := range lines {
		if line == "" {
			continue
		}
		var cons []string
		runes := []rune(line)
		var id int
		var isNums bool
		for id <= len(runes)-3 {
			part := strings.Trim(line[id:id+3], " ")
			id += 3
			if id < len(runes) {
				id += 1
			}
			if part == "" {
				cons = append(cons, "")
				continue
			}
			num, err := strconv.Atoi(part)
			if err != nil {
				cons = append(cons, string(part[1]))
			} else {
				isNums = true
				contNums = append(contNums, num)
			}
		}
		if !isNums {
			contRow = append(contRow, cons)
		}
	}

	var stacks []stack
	for id, contNum := range contNums {
		var containers []string
		for _, conts := range contRow {
			cont := conts[id]
			if cont == "" {
				continue
			}
			containers = append([]string{cont}, containers...)
		}
		stacks = append(stacks, stack{
			num:        contNum,
			containers: containers,
		})
	}
	return stacks
}

func makeMoves(lines []string) ([]move, error) {
	var moves []move
	for _, line := range lines {
		if line == "" {
			continue
		}
		splitted := strings.Split(line, " ")
		if len(splitted) != 6 {
			return nil, fmt.Errorf("movement line %s is not correctly formatted", line)
		}
		move, err := makeMove([]string{splitted[1], splitted[3], splitted[5]})
		if err != nil {
			return nil, err
		}
		moves = append(moves, move)
	}
	return moves, nil
}

func makeMove(vals []string) (move, error) {
	if len(vals) != 3 {
		return move{}, fmt.Errorf("%v has thr wrong number of values", vals)
	}
	count, err := strconv.Atoi(vals[0])
	if err != nil {
		return move{}, fmt.Errorf("%s is not a number", vals[0])
	}

	from, err := strconv.Atoi(vals[1])
	if err != nil {
		return move{}, fmt.Errorf("%s is not a number", vals[1])
	}

	to, err := strconv.Atoi(vals[2])
	if err != nil {
		return move{}, fmt.Errorf("%s is not a number", vals[2])
	}

	return move{
		count: count,
		from:  from,
		to:    to,
	}, err
}
