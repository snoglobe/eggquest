package main

type GameMap struct {
	Dungeons []Dungeon
}

func NewGameMap() GameMap {
	//Return a new game map of a single level for now
	l := NewLevel()
	levels := make([]Level, 0)
	levels = append(levels, l)
	d := Dungeon{Name: "default", Levels: levels}
	dungeons := make([]Dungeon, 0)
	dungeons = append(dungeons, d)
	gm := GameMap{Dungeons: dungeons}
	return gm
}
