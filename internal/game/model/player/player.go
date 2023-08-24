package player

import (
	"gameserver/internal/game/model/card"

	pb "gameserver/proto/pb/room_manager"
)

type stream pb.RoomManager_EnterRoomServer

type Player struct {
	*Audience
	Hand   []card.Card
	Status Status
}

func NewPlayer(audience *Audience) *Player {
	return &Player{
		Audience: audience,
		Status:   0,
	}
}

func (p *Player) SetHand(hand []card.Card) {
	p.Hand = hand
}

func (p *Player) GetHand() []card.Card {
	return p.Hand
}

func (p *Player) GetConnection() stream {
	return p.connection
}

func (p *Player) SetConnection(connection stream) {
	p.connection = connection
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
