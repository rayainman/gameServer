package game_logic

import (
	"fmt"
	"gameserver/internal/game/game_rules"
	"gameserver/internal/game/model/player"
	"gameserver/internal/game/model/room"
	"sync"
	"time"
)

const turnDuration = 30 * time.Second
const maxTurns = 100

func StartGame(room *room.Room) string {

	room.Game.Start()

	for round := 1; round <= maxTurns; round++ {
		// fmt.Printf("Room %d is in stage %s\n", room.ID, room.Stage)

		// Simulate players making decisions
		waitForPlayerDecisions(room)
		// Perform game updates and synchronization
		// ...
		if err := handleStageResults(room); err != nil {
			return err.Error()
		}

		// Sleep to simulate a game tick
		time.Sleep(1 * time.Second)

		if room.Game.Stage.IsFinish() {
			fmt.Printf("Room %s has finished\n", room.ID)
			// room.DeleteRoom()
			return ""
		}
	}

	return ""
}

func waitForPlayerDecisions(room *room.Room) {
	var wg sync.WaitGroup
	for _, seat := range room.Game.Table.GetActiveSeats() {
		wg.Add(1)
		go makeDecision(seat.Player, room, &wg)
	}
	wg.Wait()
}

func makeDecision(player *player.Player, room *room.Room, wg *sync.WaitGroup) {

	currentStage := room.Game.CurrentStage()
	timeout, ok := game_rules.DecisionTimeouts[currentStage]

	if !ok {
		fmt.Printf("Error: no timeout defined for stage %s\n", currentStage)
		return
	}

	defer wg.Done()
	select {
	case <-time.After(timeout):
	// fmt.Printf("Player %s in Room %d exceeded decision timeout\n", player.Username, room.ID)
	// Handle timeout logic, e.g., default decision or penalization
	case input := <-player.GetChannel():
		// Simulate player decision logic
		fmt.Printf("Player %s in Room %s is making a decision : %s \n", player.Username, room.ID, input)
		// Update player's state based on decision
		// ...
		// switch currentStage {
		// case game_rules.StageOpen:
		// 	// Handle open stage input logic
		// case game_rules.StageDeal:
		// 	// Handle deal stage input logic
		// case game_rules.StageBet:
		// 	// Handle bet stage input logic
		// case game_rules.StageShowdown:
		// 	// Handle showdown stage input logic
		// case game_rules.StageEnd:
		// 	// Handle end stage input logic
		// }
	}
}
func handleStageResults(room *room.Room) error {
	currentStage := room.Game.CurrentStage()

	err := error(nil)
	switch currentStage {
	case game_rules.StageOpen:
		err = handleOpenStage(room)
	case game_rules.StageDeal:
		err = handleDealStage(room)
	case game_rules.StageBet:
		err = handleBetStage(room)
	case game_rules.StageShowdown:
		err = handleShowdownStage(room)
	case game_rules.StageEnd:
		err = handleEndStage(room)
	}

	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		return err
	}

	return nil
}

func handleOpenStage(room *room.Room) error {

	if !room.Game.Stage.IsOpen() {
		return fmt.Errorf("room %s is not in the open stage", room.ID)
	}

	fmt.Printf("Room %s is in the open stage\n", room.ID)

	room.Game.Table.Shuffle()

	room.Game.Stage.Next()
	return nil
}

func handleDealStage(room *room.Room) error {
	if !room.Game.Stage.IsDeal() {
		return fmt.Errorf("room %s is not in the deal stage", room.ID)
	}

	fmt.Printf("Room %s is in the deal stage\n", room.ID)

	room.Game.Stage.Next()

	return nil
}

func handleBetStage(room *room.Room) error {

	if !room.Game.Stage.IsBet() {
		return fmt.Errorf("room %s is not in the bet stage", room.ID)
	}

	fmt.Printf("Room %s is in the bet stage\n", room.ID)

	game_rules.PickWinners(room.Game.Table)
	room.Game.Stage.Next()

	return nil
}

func handleShowdownStage(room *room.Room) error {

	if !room.Game.Stage.IsShowdown() {
		return fmt.Errorf("room %s is not in the showdown stage", room.ID)
	}

	fmt.Printf("Room %s is in the showdown stage\n", room.ID)

	activeSeats := len(room.Game.Table.GetActiveSeats())

	// 印出存活玩家人數
	fmt.Printf("Room %s is in the showdown stage, %d players alive\n", room.ID, activeSeats)
	if activeSeats == 1 {
		room.Game.Stage.Next()
	} else {
		room.Game.Stage.Previous()
	}
	return nil

}

func handleEndStage(room *room.Room) error {

	if !room.Game.Stage.IsEnd() {
		return fmt.Errorf("room %s is not in the end stage", room.ID)
	}

	fmt.Printf("Room %s is in the end stage\n %s is winner \n", room.ID, game_rules.GetWinner(room.Game.Table).Username)

	room.Game.Stage.Next()

	return nil

}
