package game_rules

import "time"

var DecisionTimeouts = map[GameStage]time.Duration{
	StageOpen:     1 * time.Second,
	StageDeal:     0 * time.Second,
	StageBet:      0 * time.Second,
	StageShowdown: 0 * time.Second,
	StageEnd:      1 * time.Second,
}
