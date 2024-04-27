package main

import (
	"fmt"
	"log"
	"mazemazego/position"

	"github.com/mattn/go-tty"
)

type MapDesignType [10][10]int

type MazeDesign struct {
	Design      MapDesignType
	EntrancePos position.Pos
	ExitPos     position.Pos
}

func makeStageStr(design MapDesignType, playerPos position.Pos) string {
	var mazeMapStr string
	for y, rows := range design {
		for x, mass := range rows {
			switch {
			case playerPos.X == x && playerPos.Y == y:
				mazeMapStr += "OO"
			case mass == 1:
				mazeMapStr += "##"
			case mass == 0:
				mazeMapStr += "  "
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

func movePlayer(playerPos position.Pos, directionPos position.Pos, mazeMap MapDesignType) (position.Pos, bool) {
	if !(isValidDirectionPos(directionPos)) {
		return position.Pos{}, false
	}

	newPlayerPos := position.Add(playerPos, directionPos)

	minX, minY := 0, 0
	maxX, maxY := len(mazeMap[0])-1, len(mazeMap)-1
	if newPlayerPos.X < minX || newPlayerPos.X > maxX ||
		newPlayerPos.Y < minY || newPlayerPos.Y > maxY {
		return position.Pos{}, false
	}

	if mazeMap[newPlayerPos.Y][newPlayerPos.X] == 0 {
		return newPlayerPos, true
	}
	return position.Pos{}, false
}

func escaped(playerPos, exitPos position.Pos) bool {
	return position.IsSame(playerPos, exitPos)
}

func main() {
	tty, err := tty.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer tty.Close()

	maze := MazeDesign{
		Design: MapDesignType{
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
	playerPos := maze.EntrancePos

GAME_LOOP:
	for {
		stageStr := makeStageStr(maze.Design, playerPos)
		fmt.Println(stageStr)

		if escaped(playerPos, maze.ExitPos) {
			fmt.Println("Successful Escape!")
			return
		}

		r, err := tty.ReadRune()
		if err != nil {
			log.Fatal(err)
		}

		// handle input
		switch r {
		case 'x':
			break GAME_LOOP
		case 'w':
			up := position.Pos{X: 0, Y: -1}
			if p, ok := movePlayer(playerPos, up, maze.Design); ok {
				playerPos = p
			}
		case 's':
			down := position.Pos{X: 0, Y: 1}
			if p, ok := movePlayer(playerPos, down, maze.Design); ok {
				playerPos = p
			}
		case 'a':
			left := position.Pos{X: -1, Y: 0}
			if p, ok := movePlayer(playerPos, left, maze.Design); ok {
				playerPos = p
			}
		case 'd':
			right := position.Pos{X: 1, Y: 0}
			if p, ok := movePlayer(playerPos, right, maze.Design); ok {
				playerPos = p
			}
		}
	}
}
