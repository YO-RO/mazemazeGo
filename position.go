package main

type Pos struct {
	X int
	Y int
}

func (p Pos) IsSame(other Pos) bool {
	return p.X == other.X && p.Y == other.Y
}

func (p *Pos) Add(ex Pos) {
	p.X += ex.X
	p.Y += ex.Y
}
