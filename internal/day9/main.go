package day9

import (
	"embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var file embed.FS

type direction int

const (
	UP    = iota
	DOWN  = iota
	LEFT  = iota
	RIGHT = iota
)

type pos struct {
	x int
	y int
}

func (p pos) add(o pos) pos {
	return pos{
		x: p.x + o.x,
		y: p.y + o.y,
	}
}

type move struct {
	dir direction
	num int
}

var (
	upPos    = pos{x: 0, y: 1}
	downPos  = pos{x: 0, y: -1}
	rightPos = pos{x: 1, y: 0}
	leftPos  = pos{x: -1, y: 0}
)

func Execute() error {
	fmt.Println("*** Day 9 ***")
	bytes, err := file.ReadFile("input.txt")
	if err != nil {
		return nil
	}
	lines := strings.Split(string(bytes), "\n")
	moves, err := getMoves(lines)
	if err != nil {
		return err
	}

	num := 2
	var snake []pos
	for id := 0; id < num; id++ {
		snake = append(snake, pos{})
	}

	visited := map[pos]bool{}
	visited[snake[num-1]] = true
	for _, move := range moves {
		for step := 0; step < move.num; step++ {
			var movePos pos
			if move.dir == UP {
				movePos = upPos
			} else if move.dir == DOWN {
				movePos = downPos
			} else if move.dir == LEFT {
				movePos = leftPos
			} else {
				movePos = rightPos
			}

			var newSnake []pos
			newPos := snake[0].add(movePos)
			newSnake = append(newSnake, newPos)
			for id := 1; id < num; id++ {
				tailPos := snake[id]
				if tailMustMove(newSnake[id-1], tailPos) {
					tailPos = snake[id-1]
				}
				newSnake = append(newSnake, tailPos)
			}
			visited[newSnake[num-1]] = true
			snake = newSnake
		}
	}
	fmt.Printf("Visited positions: %d\n", len(visited))
	return nil
}

func getMoves(lines []string) ([]move, error) {
	var moves []move
	for _, line := range lines {
		if line == "" {
			continue
		}
		sp := strings.Split(line, " ")
		if len(sp) != 2 {
			return []move{}, fmt.Errorf("%s is not correctly formatted", line)
		}
		num, err := strconv.Atoi(sp[1])
		if err != nil {
			return []move{}, err
		}
		var dir direction
		switch sp[0] {
		case "U":
			dir = UP
		case "D":
			dir = DOWN
		case "L":
			dir = LEFT
		case "R":
			dir = RIGHT
		default:
			return []move{}, fmt.Errorf("%s is not a valid direction", sp[0])

		}
		moves = append(moves, move{
			dir: dir,
			num: num,
		})
	}
	return moves, nil
}

func tailMustMove(tail, head pos) bool {
	if abs(tail.x-head.x) > 1 || abs(tail.y-head.y) > 1 {
		return true
	}
	return false
}

func abs(val int) int {
	if val > 0 {
		return val
	}
	return -val
}
