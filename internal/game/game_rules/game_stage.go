package game_rules

import "fmt"

type GameStage int

const (
	StageOpen     GameStage = iota // 開始
	StageDeal                      // 發牌
	StageBet                       // 下注
	StageShowdown                  // 比牌
	StageEnd                       // 結束
	StageFinish                    // 真的結束
)

var stageNames = []string{
	"Open",
	"Deal",
	"Bet",
	"Showdown",
	"End",
}

func (stage GameStage) String() string {
	if stage >= 0 && int(stage) < len(stageNames) {
		return stageNames[stage]
	}
	return fmt.Sprintf("Unknown Stage (%d)", stage)
}

func NewGameStage() *GameStage {
	stage := StageOpen
	return &stage
}

func (stage *GameStage) CurrentStage() GameStage {
	return *stage
}

func (stage *GameStage) Start() {
	*stage = StageOpen
}

func (stage *GameStage) Next() {
	*stage = *stage + 1
}

func (stage *GameStage) Previous() {
	*stage = *stage - 1
}

func (stage GameStage) IsOpen() bool {
	return stage == StageOpen
}

func (stage GameStage) IsDeal() bool {
	return stage == StageDeal
}

func (stage GameStage) IsBet() bool {
	return stage == StageBet
}

func (stage GameStage) IsShowdown() bool {
	return stage == StageShowdown
}

func (stage GameStage) IsEnd() bool {
	return stage == StageEnd
}

func (stage GameStage) IsFinish() bool {
	return stage == StageFinish
}
