package day7

import (
	"embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input embed.FS

func Execute() error {
	fmt.Println("*** Day 7 ***")
	bytes, err := input.ReadFile("input.txt")
	if err != nil {
		return err
	}
	lines := strings.Split(string(bytes), "\n")
	var currLine int

	rootDir := directory{
		dirName:          "/",
		directoryContent: []content{},
		parent:           nil,
	}

	currDir := &rootDir
	for currLine < len(lines) {
		line := lines[currLine]
		if line == "" {
			currLine++
			continue
		}
		runes := []rune(line)
		if runes[0] == '$' {
			cmd := strings.Split(string(runes[2:]), " ")
			cmdType := cmd[0]
			if cmdType == "ls" {
				psl := parseLs(currDir, lines[currLine+1:])
				currLine += psl
			}
			if cmdType == "cd" {
				if cmd[1] == "/" {
					currDir = &rootDir
				} else if cmd[1] == ".." {
					par := currDir.parentDir()
					if par != nil {
						currDir = par
					}
				} else {
					ch := cmd[1]
					for _, chc := range currDir.directoryContent {
						switch t := chc.(type) {
						case *directory:
							if chc.name() == ch {
								currDir = t
								break
							}
						default:
							continue
						}
					}
				}
			}
		}
		currLine++
	}
	totalSize := dirSize(rootDir)
	fmt.Printf("Total size: %d\n", totalSize)
	return nil
}

func parseLs(dir *directory, lines []string) int {
	var currId int
	for currId < len(lines) {
		line := lines[currId]
		if line == "" {
			currId++
			continue
		}
		runes := []rune(line)
		if runes[0] == '$' {
			break
		}
		sp := strings.Split(line, " ")
		if sp[0] == "dir" {
			dir.addContent(&directory{
				dirName:          sp[1],
				directoryContent: []content{},
				parent:           dir,
			})
			currId++
			continue
		}
		size, err := strconv.Atoi(sp[0])
		if err != nil {
			panic(err)
		}
		dir.addContent(&file{
			fileName: sp[1],
			fileSize: size,
			parent:   dir,
		})
		currId++
	}
	return currId
}

func dirSize(dir directory) int {
	var totalSize int
	for _, cons := range dir.directoryContent {
		switch t := cons.(type) {
		case *directory:
			if t.size() < 100_000 {
				totalSize += t.size()
			}
			totalSize += dirSize(*t)
		default:
			continue
		}
	}
	return totalSize
}
