package player

import (
	"gameserver/internal/game/model/card"

	pb "gameserver/proto/pb/room_manager"
)

type stream pb.RoomManager_EnterRoomServer

type Player struct {
	*Audience
	Hand    []card.Card
	Status  Status
	Score   int64
	isHuman bool
}

func NewPlayer(audience *Audience) *Player {

	return &Player{
		Audience: audience,
		isHuman:  true,
	}
}

func (p *Player) SetHand(hand []card.Card) {
	p.Hand = hand
}

func (p *Player) GetHand() []card.Card {
	return p.Hand
}

func (p *Player) GetConnection() stream {
	return p.Connection
}

func (p *Player) SetConnection(connection stream) {
	p.Connection = connection
}

func GetActivePlayers(player []*Player) []*Player {
	var activePlayers []*Player
	for _, p := range player {
		if !p.IsOut() {
			activePlayers = append(activePlayers, p)
		}
	}
	return activePlayers
}

func (p *Player) IsOut() bool {
	return p.Status.IsOut()
}

// 出局
func (p *Player) Out() {
	p.Status = StatusLost
}

func (p *Player) Win() {
	p.Status = StatusWon
}

func (p *Player) SetChannel(channel chan string) {
	p.InputChannel = channel
}

func (p *Player) GetChannel() <-chan string {
	return p.InputChannel
}

// // 隨機產生分數 100-10000
// func (p *Player) GetPlayerFakeScore() int64 {

// 	// 隨機延遲 0.001 - 0.005 秒
// 	time.Sleep(time.Duration(rand.Intn(5)+1) * time.Millisecond)

// 	seed := time.Now().UnixNano()
// 	rng := rand.New(rand.NewSource(seed))
// 	p.Score = rng.Int63n(10000-100) + 100

// 	return p.Score

// }

func (p *Player) InitScore() {
	p.Score = 0
}
