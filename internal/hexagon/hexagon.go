// Draws a regular hexagon, used by hexmap
package hexagon

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
)

// Hexagon draws a regular hexagon of a given size to an IMDraw
type Hexagon struct {
	r  float64
	dx float64
	dy float64
}

const (
	cos30 = 0.86602540378
	sin30 = 0.5
)

// New creates a hexagon with radius, r
func New(r int) Hexagon {
	return Hexagon{
		r:  float64(r),
		dx: float64(r) * cos30,
		dy: float64(r) * sin30,
	}
}

// DrawTo draws the hexagon to an IMDraw at x, y, centered at O
// in this orientation:
//      *
//    *   *
//      O
//    *   *
//      *
func (h Hexagon) DrawTo(imd *imdraw.IMDraw, x, y int) {
	xf := float64(x)
	yf := float64(y)
	imd.Push(pixel.V(xf, yf+h.r), pixel.V(xf+h.dx, yf+h.dy))
	imd.Push(pixel.V(xf+h.dx, yf-h.dy), pixel.V(xf, yf-h.r))
	imd.Push(pixel.V(xf-h.dx, yf-h.dy), pixel.V(xf-h.dx, yf+h.dy))
	imd.Polygon(4)
}
