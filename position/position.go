package position

type Pos struct {
	X int
	Y int
}

func IsSame(p1, p2 Pos) bool {
	return p1.X == p2.X && p1.Y == p2.Y
}

func Move(base, ex Pos) Pos {
	return Pos{X: base.X + ex.X, Y: base.Y + ex.Y}
}
