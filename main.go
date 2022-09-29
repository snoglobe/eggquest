package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	_ "image/png"
	"log"
)

type Game struct {
	Map GameMap
}

func NewGame() *Game {
	g := &Game{}
	g.Map = NewGameMap()
	return g
}

// Update is called each tic.
func (g *Game) Update() error {
	return nil
}

// Draw is called each draw cycle and is where we will blit.
func (g *Game) Draw(screen *ebiten.Image) {
	//Draw the Map
	gd := NewGameData()
	level := g.Map.Dungeons[0].Levels[0]
	for x := 0; x < gd.ScreenWidth; x++ {
		for y := 0; y < gd.ScreenHeight; y++ {
			tile := level.Tiles[level.GetIndexFromXY(x, y)]
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(tile.PixelX), float64(tile.PixelY))
			screen.DrawImage(tile.Image, op)
		}
	}
}

// Layout will return the screen dimensions.
func (g *Game) Layout(w, h int) (int, int) {
	gd := NewGameData()
	return gd.TileWidth * gd.ScreenWidth, gd.TileHeight * gd.ScreenHeight
}

func main() {
	g := NewGame()
	ebiten.SetWindowResizable(true)
	ebiten.SetWindowTitle("eggquest")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
