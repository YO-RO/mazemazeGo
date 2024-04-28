package main

import (
	"errors"
	"fmt"
)

type PlayerMoveDirection int

const (
	Up PlayerMoveDirection = iota
	Down
	Left
	Right
)

type MazeMazeGo struct {
	Maze                  MazeDesign
	PlayerPos             Pos
	IsShowingCorrectRoute bool
}

func NewMazeMazeGo(maze MazeDesign) MazeMazeGo {
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
	maxX, maxY := len(m.Maze.Design[0])-1, len(m.Maze.Design)-1
	if newPlayerPos.X < minX || newPlayerPos.X > maxX ||
		newPlayerPos.Y < minY || newPlayerPos.Y > maxY {
		return true
	}

	// 壁と衝突していないか確認する
	return m.Maze.Design[newPlayerPos.Y][newPlayerPos.X] == 1
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

func (m *MazeMazeGo) ToggleToShowCorrectRoute() {
	m.IsShowingCorrectRoute = !m.IsShowingCorrectRoute
}

func (m MazeMazeGo) String() string {
	var playAreaCells [10][10]string
	// maze design
	for y, rows := range m.Maze.Design {
		for x, cell := range rows {
			switch cell {
			case 0:
				playAreaCells[y][x] = "  "
			case 1:
				playAreaCells[y][x] = "##"
			default:
				playAreaCells[y][x] = "??"
			}
		}
	}
	// player
	boldFront := func(s string) string {
		return fmt.Sprintf("\033[1m%s\033[m", s)
	}
	playAreaCells[m.PlayerPos.Y][m.PlayerPos.X] = boldFront("OO")
	// maze correct route
	if m.IsShowingCorrectRoute {
		for y, rows := range m.Maze.CorrectRoute {
			for x, cell := range rows {
				switch cell {
				case 0: // do nothing
				case 1:
					frontStr := playAreaCells[y][x]
					redBack := func(s string) string {
						return fmt.Sprintf("\033[41m%s\033[m", s)
					}
					blackFront := func(s string) string {
						return fmt.Sprintf("\033[30m%s\033[m", s)
					}
					playAreaCells[y][x] = redBack(blackFront(frontStr))
				default:
					playAreaCells[y][x] = "??"
				}
			}
		}
	}

	var playAreaStr string
	for _, rows := range playAreaCells {
		for _, cell := range rows {
			playAreaStr += cell
		}
		playAreaStr += "\n"
	}
	return playAreaStr
}
