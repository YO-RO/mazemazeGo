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

func makeStageStr(design MapDesignType, player position.Pos) string {
	var mazeMapStr string
	for y, rows := range design {
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

func movePlayer(mazeMap MapDesignType, player position.Pos, directionPos position.Pos) (position.Pos, bool) {
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

func escaped(player, exitPos position.Pos) bool {
	return position.IsSame(player, exitPos)
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
	player := maze.EntrancePos

GAME_LOOP:
	for {
		stageStr := makeStageStr(maze.Design, player)
		fmt.Println(stageStr)

		if escaped(player, maze.ExitPos) {
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
			if p, ok := movePlayer(maze.Design, player, up); ok {
				player = p
			}
		case 's':
			down := position.Pos{X: 0, Y: 1}
			if p, ok := movePlayer(maze.Design, player, down); ok {
				player = p
			}
		case 'a':
			left := position.Pos{X: -1, Y: 0}
			if p, ok := movePlayer(maze.Design, player, left); ok {
				player = p
			}
		case 'd':
			right := position.Pos{X: 1, Y: 0}
			if p, ok := movePlayer(maze.Design, player, right); ok {
				player = p
			}
		}
	}
}
