package day8

import (
	"embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var file embed.FS

type tree struct {
	height int
	score  int
}

func Execute() error {
	fmt.Println("*** Day 8 ***")
	bytes, err := file.ReadFile("input.txt")
	if err != nil {
		return err
	}
	lines := strings.Split(string(bytes), "\n")
	trees, err := linesToTrees(lines)
	if err != nil {
		return err
	}

	rows := len(trees)
	cols, err := colLen(trees)
	if err != nil {
		return err
	}

	//left
	for idRow := 1; idRow < rows-1; idRow++ {
		for idCol := 1; idCol < cols-1; idCol++ {
			tree := &trees[idRow][idCol]
			h := tree.height
			currScore := 1

			score := 0
			for idR := idRow + 1; idR < rows; idR++ {
				t := &trees[idR][idCol]
				score++
				if t.height >= h {
					break
				}
			}
			currScore *= score

			score = 0
			for idC := idCol + 1; idC < cols; idC++ {
				t := &trees[idRow][idC]
				score++
				if t.height >= h {
					break
				}
			}
			currScore *= score

			score = 0
			for idR := idRow - 1; idR >= 0; idR-- {
				t := &trees[idR][idCol]
				score++
				if t.height >= h {
					break
				}
			}
			currScore *= score

			score = 0
			for idC := idCol - 1; idC >= 0; idC-- {
				t := &trees[idRow][idC]
				score++
				if t.height >= h {
					break
				}
			}
			currScore *= score

			tree.score = currScore
		}
	}

	score := highestScore(trees)
	fmt.Printf("Highest score: %d\n", score)
	return nil
}

func linesToTrees(lines []string) ([][]tree, error) {
	var trees [][]tree
	for _, line := range lines {
		if line == "" {
			continue
		}
		treeRow, err := lineToTreeRow(line)
		if err != nil {
			return [][]tree{}, err
		}
		trees = append(trees, treeRow)
	}
	return trees, nil
}

func lineToTreeRow(line string) ([]tree, error) {
	var treeRow []tree
	for _, ch := range line {
		h, err := strconv.Atoi(string(ch))
		if err != nil {
			return []tree{}, err
		}
		treeRow = append(treeRow, tree{
			height: h,
			score:  0,
		})
	}
	return treeRow, nil
}

func colLen(trees [][]tree) (int, error) {
	rowLens := map[int]bool{}
	for idRow := 0; idRow < len(trees); idRow++ {
		rowLen := len(trees[idRow])
		rowLens[rowLen] = true
	}
	if len(rowLens) != 1 {
		return 0, fmt.Errorf("trees has not equal length rows")
	}
	for k := range rowLens {
		return k, nil
	}
	return 0, nil
}

func highestScore(trees [][]tree) int {
	var score int
	for _, treeRow := range trees {
		for _, tree := range treeRow {
			if tree.score > score {
				score = tree.score
			}
		}
	}
	return score
}
