package creature

import (
	"github.com/SpacedMonkeyTCT/hexagon/internal/hexagon"
	"github.com/SpacedMonkeyTCT/hexagon/internal/hexmap"
	"github.com/faiface/pixel/imdraw"
	"golang.org/x/image/colornames"
)

// Creature moves around the hexmap
type Creature struct {
	hex     hexagon.Hexagon
	hm      hexmap.HexMap
	x       int
	y       int
	waitFor int
}

const (
	delay = 50
)

// New creates a creature that lives on a HexMap
func New(hm hexmap.HexMap) *Creature {
	return &Creature{
		hex:     hexagon.New(hm.Size / 2),
		hm:      hm,
		x:       0,
		y:       0,
		waitFor: delay,
	}
}

func (c *Creature) Update() {
	if c.waitFor > 0 {
		c.waitFor--
		return
	}
	c.waitFor = delay
	if c.x == 0 && c.y < c.hm.MapHeight-1 {
		c.y++
	} else if c.y == c.hm.MapHeight-1 && c.x < c.hm.MapWidth-1 {
		c.x++
	} else if c.x == c.hm.MapWidth-1 && c.y > 0 {
		c.y--
	} else if c.y == 0 && c.x > 0 {
		c.x--
	}
}

// DrawTo draws the creature to an IMDraw
func (c Creature) DrawTo(imd *imdraw.IMDraw) {
	imd.Color = colornames.Violet
	sx, sy := c.hm.ToScreen(c.x, c.y)
	c.hex.DrawTo(imd, sx, sy)
}
