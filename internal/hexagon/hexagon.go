// Draws a regular hexagon, used by hexmap
package hexagon

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
)

// Hexagon draws a regular hexagon of a given size to an IMDraw
type Hexagon struct {
	r       float64
	dx      float64
	dy      float64
	outline float64
	angle   float64
	origin  pixel.Vec
}

const (
	cos30 = 0.86602540378
	sin30 = 0.5
)

// New creates a hexagon with radius, r, centered at O
// in this orientation:
//        *
//    *       *
//        O
//    *       *
//        *
func New(r int) *Hexagon {
	rf := float64(r)
	return &Hexagon{
		r:       rf,
		dx:      rf * cos30,
		dy:      rf * sin30,
		outline: 0,
		angle:   0,
		origin:  pixel.V(0, 0),
	}
}

// DrawTo draws the hexagon to an IMDraw
func (h Hexagon) DrawTo(imd *imdraw.IMDraw) {
	imd.EndShape = imdraw.RoundEndShape
	imd.Push(
		pixel.V(0, h.r).Rotated(h.angle).Add(h.origin),
		pixel.V(h.dx, h.dy).Rotated(h.angle).Add(h.origin),
		pixel.V(h.dx, -h.dy).Rotated(h.angle).Add(h.origin),
		pixel.V(0, -h.r).Rotated(h.angle).Add(h.origin),
		pixel.V(-h.dx, -h.dy).Rotated(h.angle).Add(h.origin),
		pixel.V(-h.dx, h.dy).Rotated(h.angle).Add(h.origin))
	imd.Polygon(h.outline)
}

// Outline moves the origin of the hexagon to coords (x, y)
func (h *Hexagon) Outline(o float64) {
	h.outline = o
}

// MoveTo moves the origin of the hexagon to coords (x, y)
func (h *Hexagon) MoveTo(p pixel.Vec) {
	h.origin = p
}

// Rotated sets the rotation to angle, a, in radians
func (h *Hexagon) Rotated(a float64) {
	h.angle = a
}
