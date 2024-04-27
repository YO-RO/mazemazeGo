package main

type MapDesignType [10][10]int

type MazeDesign struct {
	Design      MapDesignType
	EntrancePos Pos
	ExitPos     Pos
}

func GetMazeDesign() MazeDesign {
	maze := MazeDesign{
		Design: MapDesignType{
			{1, 0, 1, 1, 1, 1, 1, 1, 1, 1},
			{1, 0, 1, 0, 1, 0, 0, 0, 1, 1},
			{1, 0, 1, 0, 1, 0, 1, 0, 1, 1},
			{1, 0, 0, 0, 0, 0, 1, 0, 0, 1},
			{1, 0, 0, 1, 1, 1, 1, 0, 1, 1},
			{1, 0, 1, 1, 0, 0, 0, 0, 0, 1},
			{1, 0, 1, 0, 0, 1, 1, 1, 0, 1},
			{1, 0, 1, 0, 0, 0, 0, 1, 0, 1},
			{1, 0, 0, 1, 1, 1, 0, 1, 0, 1},
			{1, 1, 1, 1, 1, 1, 0, 1, 1, 1},
		},
		EntrancePos: Pos{X: 1, Y: 0},
		ExitPos:     Pos{X: 6, Y: 9},
	}
	return maze
}
