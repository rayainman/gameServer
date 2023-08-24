package game

import (
	"gameserver/internal/game/game_rules"
	"gameserver/internal/game/model/table"
)

type Game struct {
	// Players []*player.Player // 實際參與遊戲的玩家
	Table *table.Table
	Stage *game_rules.GameStage
}

func NewGame() *Game {
	return &Game{
		Stage: game_rules.NewGameStage(),
		Table: table.NewTable(),
	}
}

func (g *Game) Start() {
	// g.Table.ArrangeSeats(players)
}

func (g *Game) CurrentStage() game_rules.GameStage {
	return g.Stage.CurrentStage()
}

// func (g *Game) DealCards() {
// 	for _, player := range g.Players {
// 		playerCard := g.Deck.Deal()
// 		// 分发牌给玩家...
// 	}
// }
