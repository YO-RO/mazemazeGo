package main

import (
	"fmt"
	"log"
	"mazemazego/position"

	"github.com/mattn/go-tty"
)

type MapDetailType [10][10]int

type Maze struct {
	MapDetail   MapDetailType
	EntrancePos position.Pos
	ExitPos     position.Pos
}

func makeStageStr(mazeMap MapDetailType, player position.Pos) string {
	var mazeMapStr string
	for y, rows := range mazeMap {
		for x, mass := range rows {
			switch {
			case player.X == x && player.Y == y:
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

func isValidDirectionPos(directionPos position.Pos) bool {
	return position.IsSame(directionPos, position.Pos{X: 1, Y: 0}) ||
		position.IsSame(directionPos, position.Pos{X: -1, Y: 0}) ||
		position.IsSame(directionPos, position.Pos{X: 0, Y: 1}) ||
		position.IsSame(directionPos, position.Pos{X: 0, Y: -1})
}

func movePlayer(mazeMap MapDetailType, player position.Pos, directionPos position.Pos) (position.Pos, bool) {
	if !(isValidDirectionPos(directionPos)) {
		return position.Pos{}, false
	}

	newPlayer := position.Move(player, directionPos)

	minX, minY := 0, 0
	maxX, maxY := len(mazeMap[0])-1, len(mazeMap)-1
	if newPlayer.X < minX || newPlayer.X > maxX ||
		newPlayer.Y < minY || newPlayer.Y > maxY {
		return position.Pos{}, false
	}

	if mazeMap[newPlayer.Y][newPlayer.X] == 0 {
		return newPlayer, true
	}
	return position.Pos{}, false
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
		EntrancePos: position.Pos{X: 1, Y: 0},
		ExitPos:     position.Pos{X: 6, Y: 9},
	}
	player := maze.EntrancePos

	for {
		stageStr := makeStageStr(maze.MapDetail, player)
		fmt.Println(stageStr)

		if position.IsSame(player, maze.ExitPos) {
			fmt.Println("Successful Escape!")
			return
		}

		r, err := tty.ReadRune()
		if err != nil {
			log.Fatal(err)
		}

		if r == 'x' {
			break
		}

		var dirPos position.Pos
		switch r {
		case 'w':
			dirPos = position.Pos{X: 0, Y: -1}
		case 's':
			dirPos = position.Pos{X: 0, Y: 1}
		case 'a':
			dirPos = position.Pos{X: -1, Y: 0}
		case 'd':
			dirPos = position.Pos{X: 1, Y: 0}
		}
		if p, ok := movePlayer(maze.MapDetail, player, dirPos); ok {
			player = p
		}
	}
}
