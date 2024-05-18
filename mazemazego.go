package main

import (
	"errors"

	"github.com/fatih/color"
)

type PlayerMoveDirection int

const (
	Up PlayerMoveDirection = iota
	Down
	Left
	Right
)

type MazeMazeGo struct {
	Maze                Maze
	PlayerPos           Pos
	DisplayCorrectRoute bool
}

func NewMazeMazeGo(maze Maze) MazeMazeGo {
	m := MazeMazeGo{
		Maze:      maze,
		PlayerPos: maze.EntrancePos,
	}
	return m
}

func (m MazeMazeGo) IsEscaped() bool {
	return IsSamePos(m.PlayerPos, m.Maze.ExitPos)
}

func dirToPos(dir PlayerMoveDirection) (pos Pos, err error) {
	switch dir {
	case Up:
		pos = Pos{X: 0, Y: -1}
	case Down:
		pos = Pos{X: 0, Y: 1}
	case Left:
		pos = Pos{X: -1, Y: 0}
	case Right:
		pos = Pos{X: 1, Y: 0}
	default:
		err = errors.New("invalid dir (PlayerMoveDirection)")
	}
	return
}

func (m MazeMazeGo) detectCollisiton(newPlayerPos Pos) bool {
	// 迷路から飛び出ていないか確認する
	minX, minY := 0, 0
	maxX, maxY := len(m.Maze.wall[0])-1, len(m.Maze.wall)-1
	if newPlayerPos.X < minX || newPlayerPos.X > maxX ||
		newPlayerPos.Y < minY || newPlayerPos.Y > maxY {
		return true
	}

	// 壁と衝突していないか確認する
	return m.Maze.IsWall(newPlayerPos)
}

func (m *MazeMazeGo) MovePlayer(moveDir PlayerMoveDirection) bool {
	dirPos, err := dirToPos(moveDir)
	if err != nil {
		return false
	}
	newPlayerPos := AddPos(m.PlayerPos, dirPos)
	if !m.detectCollisiton(newPlayerPos) {
		m.PlayerPos = newPlayerPos
		return true
	}
	return false
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
	tiles[m.PlayerPos.Y][m.PlayerPos.X] = tiles[m.PlayerPos.Y][m.PlayerPos.X].Overwrite(NewTile("OO", color.Bold))

	var str string
	for _, rows := range tiles {
		for _, tile := range rows {
			str += tile.String()
		}
		str += "\n"
	}
	return str
}
