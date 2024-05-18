package main

import (
	"github.com/fatih/color"
)

type MazeMazeGo struct {
	Maze                Maze
	Player              Player
	DisplayCorrectRoute bool
}

func NewMazeMazeGo(maze Maze) MazeMazeGo {
	m := MazeMazeGo{
		Maze:   maze,
		Player: NewPlayer(maze.EntrancePos, NewTile("OO", color.Bold)),
	}
	return m
}

func (m MazeMazeGo) IsEscaped() bool {
	return IsSamePos(m.Player.Pos, m.Maze.ExitPos)
}

func (m *MazeMazeGo) MovePlayer(moveDir PlayerMoveDirection) bool {
	movedPlayer := m.Player.Moved(moveDir)
	if !m.Maze.IsRoad(movedPlayer.Pos) {
		return false
	}
	m.Player = movedPlayer
	return true
}

func (m *MazeMazeGo) ToggleCorrectRouteDisplay() {
	m.DisplayCorrectRoute = !m.DisplayCorrectRoute
}

func (m MazeMazeGo) String() string {
	var tiles [10][10]Tile
	if m.DisplayCorrectRoute {
		tiles = m.Maze.TileLayoutWithCorrectRoute()
	} else {
		tiles = m.Maze.TileLayout()
	}
	tiles[m.Player.Pos.Y][m.Player.Pos.X] =
		tiles[m.Player.Pos.Y][m.Player.Pos.X].Overwrite(m.Player.Tile)

	var str string
	for _, rows := range tiles {
		for _, tile := range rows {
			str += tile.String()
		}
		str += "\n"
	}
	return str
}
