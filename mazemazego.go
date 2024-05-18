package main

import (
	"github.com/fatih/color"
)

type MazeMazeGo struct {
	maze                Maze
	player              Player
	displayCorrectRoute bool
}

func NewMazeMazeGo(maze Maze) MazeMazeGo {
	m := MazeMazeGo{
		maze:   maze,
		player: NewPlayer(maze.EntrancePos, NewTile("OO", color.Bold)),
	}
	return m
}

func (m MazeMazeGo) GameIsEnd() bool {
	return m.isEscaped()
}

func (m MazeMazeGo) isEscaped() bool {
	return IsSamePos(m.player.Pos, m.maze.ExitPos)
}

func (m *MazeMazeGo) MovePlayer(moveDir PlayerMoveDirection) bool {
	movedPlayer := m.player.Moved(moveDir)
	if !m.maze.IsRoad(movedPlayer.Pos) {
		return false
	}
	m.player = movedPlayer
	return true
}

func (m *MazeMazeGo) ToggleCorrectRouteDisplay() {
	m.displayCorrectRoute = !m.displayCorrectRoute
}

func (m MazeMazeGo) String() string {
	var tiles [10][10]Tile
	if m.displayCorrectRoute {
		tiles = m.maze.TileLayoutWithCorrectRoute()
	} else {
		tiles = m.maze.TileLayout()
	}
	tiles[m.player.Pos.Y][m.player.Pos.X] =
		tiles[m.player.Pos.Y][m.player.Pos.X].Overwrite(m.player.Tile)

	var str string
	for _, rows := range tiles {
		for _, tile := range rows {
			str += tile.String()
		}
		str += "\n"
	}

	if m.isEscaped() {
		str += "\nSuccessful Escape!\n"
	}
	return str
}
