package main

type Dungeon struct {
	Name   string
	Levels []Level
}

type Level struct {
	Tiles []MapTile
}

func NewLevel() Level {
	l := Level{}
	tiles := l.CreateTiles()
	l.Tiles = tiles
	return l
}
