// Currently draws a 5 x 4 map of hexagonal tiles using faiface/pixel
package main

import (
	"math/rand"

	"github.com/SpacedMonkeyTCT/hexagon/internal/creature"
	"github.com/SpacedMonkeyTCT/hexagon/internal/hexmap"
	"github.com/SpacedMonkeyTCT/hexagon/internal/navigation"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

func main() {
	pixelgl.Run(run)
}

const (
	winW = 1024
	winH = 768
	mapW = 7
	mapH = 6
)

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Hexagon!",
		Bounds: pixel.R(0, 0, winW, winH),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	imd := imdraw.New(nil)
	hm := hexmap.New(mapW, mapH, winW, winH)
	n := navigation.NewNavigation(mapW, mapH)
	setWalls(hm, n)
	c := creature.New(hm, n)

	for !win.Closed() {
		win.Clear(colornames.Aliceblue)
		c.Update()
		imd.Clear()
		hm.DrawTo(imd)
		c.DrawTo(imd)
		imd.Draw(win)
		win.Update()
	}
}

func setWalls(hm *hexmap.HexMap, n navigation.Navigation) {
	wallMax := (mapW * mapH) / 4
	for i := 0; i < wallMax; i++ {
		x := rand.Intn(mapW)
		y := rand.Intn(mapH)
		hm.SetWall(x, y)
		n.SetWall(x, y)
	}
}
