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
		player: NewPlayer(maze.EntrancePos),
	}
	return m
}

func (m MazeMazeGo) GameIsEnd() bool {
	return m.isEscaped()
}

func (m MazeMazeGo) isEscaped() bool {
	return m.player.Pos.IsSame(m.maze.ExitPos)
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

	if m.isEscaped() {
		message := "Successful Escape!!!"
		attr := []color.Attribute{color.BgCyan, color.FgBlack, color.Bold}
		var messageTiles [10]Tile
		for i := 0; i < 10; i++ {
			// messageTiles[0] -> NewTile("Su", attr...)
			// messageTiles[1] -> NewTile("cc", attr...)
			// ...
			messageTiles[i] = NewTile(message[i*2:i*2+2], attr...)
		}

		displayRowIdx := 4
		for i, t := range tiles[displayRowIdx] {
			tiles[displayRowIdx][i] = t.Overwrite(messageTiles[i])
		}
	}

	var str string
	for _, rows := range tiles {
		for _, tile := range rows {
			str += tile.String()
		}
		str += "\n"
	}
	return str
}
