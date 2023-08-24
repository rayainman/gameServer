package game_rules

import (
	"gameserver/internal/game/model/player"
	"gameserver/internal/game/model/table"
	"math/rand"
	"time"
)

// 隨機挑選兩位未出局的玩家
// 比較牌型
// 假設玩家a贏了
// 設定玩家a狀態為StatusWon
// 設定玩家b狀態為StatusLost
func PickWinners(table *table.Table) {
	// 選出兩位存活玩家
	seed := time.Now().UnixNano()
	rng := rand.New(rand.NewSource(seed))

	activePlayers := table.GetActivePlayers()

	rng.Shuffle(len(activePlayers), func(i, j int) {
		activePlayers[i], activePlayers[j] = activePlayers[j], activePlayers[i]
	})
	for i := 0; i < len(activePlayers)-1; i = i + 2 {
		p1 := activePlayers[i]
		p2 := activePlayers[i+1]

		p1.Status = player.StatusWon
		p2.Status = player.StatusLost

	}

	// return winners
}

func PickWinner(table *table.Table) *player.Player {

	activityPlayers := table.GetActiveSeats()

	if len(activityPlayers) == 1 {
		return activityPlayers[0].Player
	}

	if len(activityPlayers) != 0 {
		return nil
	}

	return nil
}
