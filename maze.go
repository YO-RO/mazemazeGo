package main

import (
	"github.com/fatih/color"
)

type tileIndex int

type Maze struct {
	wall         [10][10]tileIndex
	correctRoute [10][10]tileIndex
	indexToTile  map[tileIndex]Tile
	EntrancePos  Pos
	ExitPos      Pos
}

func (m Maze) TileLayout() [10][10]Tile {
	var tileLayout [10][10]Tile
	for y, rows := range m.wall {
		for x, idx := range rows {
			// 必ず存在するためカンマok構文は使わない
			tileLayout[y][x] = m.indexToTile[idx]
		}
	}
	return tileLayout
}

func (m Maze) TileLayoutWithCorrectRoute() [10][10]Tile {
	var tileLayout [10][10]Tile
	baseTileLayout := m.TileLayout()
	for y, rows := range m.correctRoute {
		for x, idx := range rows {
			baseTile := baseTileLayout[y][x]
			newTile := baseTile.Overwrite(m.indexToTile[idx])
			tileLayout[y][x] = newTile
		}
	}
	return tileLayout
}

func (m Maze) IsWall(pos Pos) bool {
	return m.wall[pos.Y][pos.X] == 2
}

func GetMaze() Maze {
	ignoreTile := NewTile("")
	wallTile := NewTile("##")
	roadTile := NewTile("  ")
	correctRouteTile := NewTile("  ", color.BgRed, color.FgBlack)
	indexToTile := map[tileIndex]Tile{
		0: ignoreTile,
		1: roadTile,
		2: wallTile,
		3: correctRouteTile,
	}
	maze := Maze{
		// 1: road, 2: wall
		wall: [10][10]tileIndex{
			{2, 1, 2, 2, 2, 2, 2, 2, 2, 2},
			{2, 1, 2, 1, 2, 1, 1, 1, 2, 2},
			{2, 1, 2, 1, 2, 1, 2, 1, 2, 2},
			{2, 1, 1, 1, 1, 1, 2, 1, 1, 2},
			{2, 1, 1, 2, 2, 2, 2, 1, 2, 2},
			{2, 1, 2, 2, 1, 1, 1, 1, 1, 2},
			{2, 1, 2, 1, 1, 2, 2, 2, 1, 2},
			{2, 1, 2, 1, 1, 1, 1, 2, 1, 2},
			{2, 1, 1, 2, 2, 2, 1, 2, 1, 2},
			{2, 2, 2, 2, 2, 2, 1, 2, 2, 2},
		},
		// 0: no tile, 3: correct route
		correctRoute: [10][10]tileIndex{
			{0, 3, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 3, 0, 0, 0, 3, 3, 3, 0, 0},
			{0, 3, 0, 0, 0, 3, 0, 3, 0, 0},
			{0, 3, 3, 3, 3, 3, 0, 3, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 3, 0, 0},
			{0, 0, 0, 0, 3, 3, 3, 3, 0, 0},
			{0, 0, 0, 0, 3, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 3, 3, 3, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 3, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 3, 0, 0, 0},
		},
		indexToTile: indexToTile,
		EntrancePos: Pos{X: 1, Y: 0},
		ExitPos:     Pos{X: 6, Y: 9},
	}
	return maze
}
