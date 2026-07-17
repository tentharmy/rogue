package domain

import "github.com/tentharmy/rogue/inner/geom"

type Game struct {
	Player *Actor
}

func (g *Game) Advance(playerMove geom.Vec2) {
	g.Player.Move(playerMove)
}
