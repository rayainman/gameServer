package game_rules

import (
	"fmt"
	"gameserver/internal/game/model/card"
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

	table.Deal(3)

	for i := 0; i < len(activePlayers)-1; i = i + 2 {
		p1 := activePlayers[i]
		p2 := activePlayers[i+1]

		p1.InitScore()
		p2.InitScore()

		fmt.Println("Player1: ", p1.Audience.ID, " Player2: ", p2.Audience.ID)
		// 玩三局
		for j := 0; j <= 2; j++ {
			card1 := p1.GetHand()[j]
			card2 := p2.GetHand()[j]

			var winner, loser *player.Player

			result := card.CardCompareCards(card1, card2)

			switch result {
			case 1:
				winner = p1
				loser = p2
			case 2:
				winner = p2
				loser = p1
			default:
				fmt.Println("Round ", j+1, ": ", "Tie", " winner's card: ", card1, " loser's card: ", card2)
				continue
			}

			winner.Score = winner.Score + 1

			fmt.Println("Round ", j+1, ": ", winner.Audience.ID, " win ", loser.Audience.ID, " lose", " winner's card: ", card1, " loser's card: ", card2)
		}

		// var winner, loser *player.Player
		// if p1.Score > p2.Score {
		// 	winner = p1
		// 	loser = p2
		// } else {
		// 	winner = p2
		// 	loser = p1
		// }

		// winner.Win()
		// loser.Out()

		var winner, loser *player.Player
		switch {
		case p1.Score > p2.Score:
			winner = p1
			loser = p2
		case p1.Score < p2.Score:
			winner = p2
			loser = p1
		default:
			fmt.Println("Tie")
			continue
		}

		winner.Win()
		loser.Out()

		fmt.Println("Winner: ", winner.Audience.ID, " Loser: ", loser.Audience.ID)
	}
}

func GetWinner(table *table.Table) *player.Player {

	activityPlayers := table.GetActiveSeats()

	if len(activityPlayers) == 1 {
		return activityPlayers[0].Player
	}

	if len(activityPlayers) != 0 {
		return nil
	}

	return nil
}
