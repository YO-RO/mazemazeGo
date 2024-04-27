package main

type Pos struct {
	X int
	Y int
}

func IsSamePos(p1, p2 Pos) bool {
	return p1.X == p2.X && p1.Y == p2.Y
}

func AddPos(p1, p2 Pos) Pos {
	return Pos{X: p1.X + p2.X, Y: p1.Y + p2.Y}
}
