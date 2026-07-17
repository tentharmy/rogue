package app

import (
	"github.com/tentharmy/rogue/inner/domain"
	"github.com/tentharmy/rogue/inner/geom"
)

type GameService struct {
	game *domain.Game
}

type GameSnapshot struct {
	PlayerPos geom.Vec2
}

func NewGame() *GameService {
	return &GameService{
		game: &domain.Game{
			Player: &domain.Actor{
				Position: geom.Vec2{
					X: 1,
					Y: 1,
				},
			},
		},
	}
}

func (serv *GameService) MovePlayer(vec geom.Vec2) {
	serv.game.Advance(vec)
}

func (serv *GameService) GetSnapshot() GameSnapshot {
	return GameSnapshot{
		PlayerPos: serv.game.Player.Position,
	}
}
