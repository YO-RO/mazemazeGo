package main

import (
	"fmt"
	"log"

	"github.com/mattn/go-tty"
)

type MapDetailType [10][10]int

type Maze struct {
	MapDetail   MapDetailType
	EntrancePos Position
	ExitPos     Position
}

type Position struct {
	X int
	Y int
}

func (p Position) isSame(other Position) bool {
	return p.X == other.X && p.Y == other.Y
}

func (p Position) isAt(x, y int) bool {
	return p.X == x && p.Y == y
}

func (p *Position) move(pos Position) {
	p.X += pos.X
	p.Y += pos.Y
}

func makeStageStr(mazeMap MapDetailType, player Position) string {
	var mazeMapStr string
	for y, rows := range mazeMap {
		for x, mass := range rows {
			switch {
			case player.isAt(x, y):
				mazeMapStr += "*"
			case mass == 1:
				mazeMapStr += "O"
			case mass == 0:
				mazeMapStr += " "
			}
		}
		mazeMapStr += "\n"
	}
	return mazeMapStr
}

func isValidDirectionPos(directionPos Position) bool {
	return directionPos.isSame(Position{1, 0}) ||
		directionPos.isSame(Position{-1, 0}) ||
		directionPos.isSame(Position{0, 1}) ||
		directionPos.isSame(Position{0, -1})
}

func movePlayer(mazeMap MapDetailType, player Position, directionPos Position) (Position, bool) {
	if !(isValidDirectionPos(directionPos)) {
		return Position{}, false
	}

	player.move(directionPos)
	x, y := player.X, player.Y

	minX, minY := 0, 0
	maxX, maxY := len(mazeMap[0])-1, len(mazeMap)-1
	if player.X < minX || player.X > maxX ||
		player.Y < minY || player.Y > maxY {
		return Position{}, false
	}

	if mazeMap[y][x] == 0 {
		return player, true
	}
	return Position{}, false
}

func main() {
	tty, err := tty.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer tty.Close()

	maze := Maze{
		MapDetail: MapDetailType{
			{1, 0, 1, 1, 1, 1, 1, 1, 1, 1},
			{1, 0, 1, 0, 1, 0, 0, 0, 1, 1},
			{1, 0, 1, 0, 1, 0, 1, 0, 1, 1},
			{1, 0, 0, 0, 0, 0, 1, 0, 0, 1},
			{1, 0, 0, 1, 1, 1, 1, 0, 1, 1},
			{1, 0, 1, 1, 0, 0, 0, 0, 0, 1},
			{1, 0, 1, 0, 0, 1, 1, 1, 0, 1},
			{1, 0, 1, 0, 0, 0, 0, 1, 0, 1},
			{1, 0, 0, 1, 1, 1, 0, 1, 0, 1},
			{1, 1, 1, 1, 1, 1, 0, 1, 1, 1},
		},
		EntrancePos: Position{1, 0},
		ExitPos:     Position{6, 9},
	}
	player := maze.EntrancePos

	for {
		stageStr := makeStageStr(maze.MapDetail, player)
		fmt.Println(stageStr)

		r, err := tty.ReadRune()
		if err != nil {
			log.Fatal(err)
		}

		if r == 'x' {
			break
		}

		var dirPos Position
		switch r {
		case 'w':
			dirPos = Position{0, -1}
		case 's':
			dirPos = Position{0, 1}
		case 'a':
			dirPos = Position{-1, 0}
		case 'd':
			dirPos = Position{1, 0}
		}
		if p, ok := movePlayer(maze.MapDetail, player, dirPos); ok {
			player = p
		}
	}
}
