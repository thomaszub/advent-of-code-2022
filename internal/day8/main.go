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
	height  int
	visible bool
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

	//outer
	for idRow := 0; idRow < rows; idRow++ {
		for idCol := 0; idCol < cols; idCol++ {
			if idCol == 0 || idRow == 0 || idCol == cols-1 || idRow == rows-1 {
				tree := &trees[idRow][idCol]
				tree.visible = true
			}
		}
	}

	//left
	for idRow := 1; idRow < rows-1; idRow++ {
		for idCol := 1; idCol < cols-1; idCol++ {
			tree := &trees[idRow][idCol]
			h := tree.height

			visible := true
			for idR := 0; idR < idRow; idR++ {
				t := &trees[idR][idCol]
				if t.height >= h {
					visible = false
				}
			}
			if visible {
				tree.visible = visible
			}

			visible = true
			for idC := 0; idC < idCol; idC++ {
				t := &trees[idRow][idC]
				if t.height >= h {
					visible = false
				}
			}
			if visible {
				tree.visible = visible
			}

			visible = true
			for idR := rows - 1; idR > idRow; idR-- {
				t := &trees[idR][idCol]
				if t.height >= h {
					visible = false
				}
			}
			if visible {
				tree.visible = visible
			}

			visible = true
			for idC := cols - 1; idC > idCol; idC-- {
				t := &trees[idRow][idC]
				if t.height >= h {
					visible = false
				}
			}
			if visible {
				tree.visible = visible
			}
		}
	}

	visibleCount := visibleTrees(trees)
	fmt.Printf("Visible trees: %d\n", visibleCount)
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
			height:  h,
			visible: false,
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

func visibleTrees(trees [][]tree) int {
	var visibleCount int
	for _, treeRow := range trees {
		for _, tree := range treeRow {
			if tree.visible {
				visibleCount++
			}
		}
	}
	return visibleCount
}
