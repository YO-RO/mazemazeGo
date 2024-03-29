package main

import "fmt"

var mazeMap = [10][10]int{
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
}

var entrancePos, exitPos = Position{1, 0}, Position{6, 9}

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

func makeStageStr(mazeMap [10][10]int, player Position) string {
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
	player := entrancePos
	stageStr := makeStageStr(mazeMap, player)
	fmt.Print(stageStr)
}
