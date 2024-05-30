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

type Player struct {
	Pos  Pos
	Tile Tile
}

func NewPlayer(initPos Pos) Player {
	return Player{
		Pos:  initPos,
		Tile: NewTile("OO", color.Bold),
	}
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

func (p Player) Moved(moveDir PlayerMoveDirection) Player {
	dirPos, err := dirToPos(moveDir)
	if err != nil {
		return p
	}
	p.Pos.Add(dirPos)
	return p
}
