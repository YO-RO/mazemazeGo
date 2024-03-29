package main

import "fmt"

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

func main() {
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
	stageStr := makeStageStr(maze.MapDetail, player)
	fmt.Print(stageStr)
}
