// Draws a map of hexagonal tiles
package hexmap

import (
	"math"

	"github.com/SpacedMonkeyTCT/hexagon/internal/hexagon"
	"github.com/faiface/pixel/imdraw"
)

// HexMap draws map of hexagons to an IMDraw, scaled to fit given dimensions
type HexMap struct {
	hex     hexagon.Hexagon
	size    int
	originX int
	originY int
	offsetX int
	offsetY int
}

const (
	mapW  = 5
	mapH  = 4
	cos30 = 0.86602540378
)

// New creates a hexmap to fit a window with dimensions w, h with a comfortable border
func New(w, h int) HexMap {
	size := int(math.Min(float64(w/(mapW+1)), float64(h/mapH))) / 2
	offsetX := size * 2
	offsetY := int(float64(offsetX) * cos30)
	screenW := mapW * offsetX
	screenH := mapH * offsetY

	return HexMap{
		hex:     hexagon.New(size),
		size:    size,
		originX: (w - screenW + size) / 2,
		originY: (h - screenH + offsetY) / 2,
		offsetX: offsetX,
		offsetY: offsetY,
	}
}

// DrawTo draws the hexmap to an IMDraw with borders between tiles.
// The tiles are arranged like:
//     * * * * *
//    * * * * *
//     * * * * *
//    * * * * *
func (hm HexMap) DrawTo(imd *imdraw.IMDraw) {
	for y := 0; y < mapH; y++ {
		for x := 0; x < mapW; x++ {
			xs, ys := hm.toScreen(x, y)
			hm.hex.DrawTo(imd, xs, ys)
		}
	}
}

// toScreen converts map coords to screen coords
func (hm HexMap) toScreen(x, y int) (int, int) {
	xs := hm.originX + hm.offsetX*x + hm.size*(y%2)
	ys := hm.originY + hm.offsetY*y
	return xs, ys
}
