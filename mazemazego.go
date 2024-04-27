package main

import "errors"

type PlayerMoveDirection int

const (
	Up PlayerMoveDirection = iota
	Down
	Left
	Right
)

type MazeMazeGo struct {
	Maze      MazeDesign
	PlayerPos Pos
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

func (m MazeMazeGo) String() string {
	var mazeMapStr string
	for y, rows := range m.Maze.Design {
		for x, mass := range rows {
			switch {
			case m.PlayerPos.X == x && m.PlayerPos.Y == y:
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
