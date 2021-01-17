package creature

import (
	"math/rand"

	"github.com/SpacedMonkeyTCT/hexagon/internal/hexagon"
	"github.com/SpacedMonkeyTCT/hexagon/internal/hexmap"
	"github.com/SpacedMonkeyTCT/hexagon/internal/navigation"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"golang.org/x/image/colornames"
)

// Creature moves around the hexmap
type Creature struct {
	hex     hexagon.Hexagon
	hm      *hexmap.HexMap
	n       navigation.Navigation
	x       int
	y       int
	waitFor int
	endX    int
	endY    int
	path    []pixel.Vec
}

const (
	delay = 50
)

// New creates a creature that lives on a HexMap
func New(hm *hexmap.HexMap, n navigation.Navigation) *Creature {
	return &Creature{
		hex:     hexagon.New(hm.Size / 2),
		hm:      hm,
		n:       n,
		x:       0,
		y:       0,
		waitFor: delay,
		path:    nil,
	}
}

func (c *Creature) Update() {
	if c.waitFor > 0 {
		c.waitFor--
		return
	}
	c.waitFor = delay

	for len(c.path) == 0 {
		c.endX = rand.Intn(c.hm.MapWidth)
		c.endY = rand.Intn(c.hm.MapHeight)
		if c.n.IsWall(c.endX, c.endY) {
			continue
		}

		c.path = c.n.Find(c.x, c.y, c.endX, c.endY)
		if len(c.path) > 0 {
			return
		}
	}
	nextPos := c.path[len(c.path)-1]
	c.path = c.path[:len(c.path)-1]
	c.x = int(nextPos.X)
	c.y = int(nextPos.Y)
}

// DrawTo draws the creature to an IMDraw
func (c Creature) DrawTo(imd *imdraw.IMDraw) {
	// draw path
	imd.Color = colornames.Violet
	for _, step := range c.path {
		x, y := c.hm.ToScreen(int(step.X), int(step.Y))
		c.hex.DrawTo(imd, x, y)
	}

	// draw destination
	imd.Color = colornames.Red
	sx, sy := c.hm.ToScreen(c.endX, c.endY)
	c.hex.DrawTo(imd, sx, sy)

	// draw creature
	imd.Color = colornames.Cornflowerblue
	sx, sy = c.hm.ToScreen(c.x, c.y)
	c.hex.DrawTo(imd, sx, sy)
}
