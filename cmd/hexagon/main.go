// Currently draws a 5 x 4 map of hexagonal tiles using faiface/pixel
package main

import (
	"github.com/SpacedMonkeyTCT/hexagon/internal/creature"
	"github.com/SpacedMonkeyTCT/hexagon/internal/hexmap"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

func main() {
	pixelgl.Run(run)
}

const (
	width  = 1024
	height = 768
)

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Hexagon!",
		Bounds: pixel.R(0, 0, width, height),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	imd := imdraw.New(nil)
	hm := hexmap.New(width, height)
	c := creature.New(hm)
	//nm := navmesh.NewHexMap(5, 4)
	//as := navmesh.NewAstar(nm.GetStart(), nm.GetEnd())

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
