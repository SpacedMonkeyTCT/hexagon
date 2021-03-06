// Draw a map of hexagonal tiles, some of which are impassable,
// and a hexagonal creature that randomly picks a valid destination and walks there
package main

import (
	"math/rand"
	"time"

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
	winW          = 1024
	winH          = 768
	mapW          = 13
	mapH          = 12
	msPerUpdate   = 10 * time.Millisecond
	creatureCount = 3
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
	setWalls(hm)
	n := navigation.NewNavigation(hm)

	ca := []*creature.Creature{}
	for i := 0; i < creatureCount; i++ {
		ca = append(ca, creature.New(hm, n))
	}

	then := time.Now()
	lag := time.Duration(0)
	for !win.Closed() {
		elapsed := time.Since(then)
		// fmt.Println("Elapsed", elapsed)
		then = time.Now()

		for lag += elapsed; lag >= msPerUpdate; lag -= msPerUpdate {
			for _, c := range ca {
				c.Update()
			}
		}

		imd.Clear()
		hm.DrawTo(imd)
		for _, c := range ca {
			c.DrawTo(imd)
		}
		win.Clear(colornames.Aliceblue)
		imd.Draw(win)
		win.Update()
	}
}

func setWalls(hm *hexmap.HexMap) {
	wallMax := (mapW * mapH) / 4
	for i := 0; i < wallMax; i++ {
		x := rand.Intn(mapW)
		y := rand.Intn(mapH)
		hm.AddWall(x, y)
	}
	hm.RemoveWall(0, 0)
}
