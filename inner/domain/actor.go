package domain

import "github.com/tentharmy/rogue/inner/geom"

type Actor struct {
	Position geom.Vec2
}

func (a *Actor) Move(move geom.Vec2) bool {
	a.Position.X += move.X
	a.Position.Y += move.Y

	return true
}
