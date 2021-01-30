// Draws a map of hexagonal tiles
package hexmap

import (
	"math"

	"github.com/SpacedMonkeyTCT/hexagon/internal/hexagon"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"golang.org/x/image/colornames"
)

// HexMap draws map of hexagons to an IMDraw, scaled to fit given dimensions
type HexMap struct {
	hex       *hexagon.Hexagon
	originX   int
	originY   int
	offsetX   int
	offsetY   int
	walls     [][]bool
	Size      int
	MapWidth  int
	MapHeight int
}

const (
	cos30 = 0.86602540378
)

// New creates a hexmap to fit a window with dimensions w, h with a comfortable border
func New(mapW, mapH, scrW, scrH int) *HexMap {
	size := int(math.Min(float64(scrW/(mapW+1)), float64(scrH/mapH))) / 2
	offsetX := size * 2
	offsetY := int(float64(offsetX) * cos30)
	drawW := mapW * offsetX
	drawH := mapH * offsetY

	walls := make([][]bool, mapW)
	for c := 0; c < mapW; c++ {
		walls[c] = make([]bool, mapH)
	}
	hex := hexagon.New(size)
	hex.Outline(4)

	return &HexMap{
		hex:       hex,
		originX:   (scrW - drawW + size) / 2,
		originY:   (scrH - drawH + offsetY) / 2,
		offsetX:   offsetX,
		offsetY:   offsetY,
		walls:     walls,
		Size:      size,
		MapWidth:  mapW,
		MapHeight: mapH,
	}
}

func (hm HexMap) AddWall(c, r int) {
	hm.walls[c][r] = true
}

func (hm HexMap) RemoveWall(c, r int) {
	hm.walls[c][r] = false
}

func (hm HexMap) IsWall(c, r int) bool {
	return hm.walls[c][r]
}

// DrawTo draws the hexmap to an IMDraw with borders between tiles.
// The tiles are arranged like:
//     * * * * *
//    * * * * *
//     * * * * *
//    * * * * *
func (hm HexMap) DrawTo(imd *imdraw.IMDraw) {
	for y := 0; y < hm.MapHeight; y++ {
		for x := 0; x < hm.MapWidth; x++ {
			if hm.IsWall(x, y) {
				imd.Color = colornames.Black
			} else {
				imd.Color = colornames.Limegreen
			}
			hm.hex.MoveTo(hm.ToScreen(x, y))
			hm.hex.DrawTo(imd)
		}
	}
}

// toScreen converts map coords to screen coords
func (hm HexMap) ToScreen(x, y int) pixel.Vec {
	xs := hm.originX + hm.offsetX*x + hm.Size*(y%2)
	ys := hm.originY + hm.offsetY*y
	return pixel.V(float64(xs), float64(ys))
}
