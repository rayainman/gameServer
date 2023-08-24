package room

import (
	"gameserver/internal/game/model/game"
	"gameserver/internal/game/model/player"
	"math/rand"
	"sync"
)

type Room struct {
	ID        string
	Audiences map[string]*player.Audience // 玩家列表
	Game      *game.Game
	mu        sync.Mutex
}

var (
	rooms     = make(map[string]*Room)
	roomMutex sync.Mutex
)

func randomID() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, 10)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

func CreateRoom() (string, bool) {
	roomMutex.Lock()
	defer roomMutex.Unlock()

	// 隨機產生一個id
	id := randomID()

	if _, ok := rooms[id]; ok {
		return "", false
	}

	room := &Room{
		ID:        id,
		Audiences: make(map[string]*player.Audience),
		Game:      game.NewGame(),
	}

	rooms[id] = room

	return id, true

}

func GetRoom(id string) *Room {
	roomMutex.Lock()
	defer roomMutex.Unlock()

	return rooms[id]
}

func (r *Room) DeleteRoom() {
	roomMutex.Lock()
	defer roomMutex.Unlock()

	delete(rooms, r.ID)
}

func (r *Room) AddAudience(p *player.Audience) bool {
	r.mu.Lock()
	defer r.mu.Unlock()

	// r.Players = append(r.Players, p...)
	// check if player already exists
	if _, ok := r.Audiences[p.ID]; ok {
		return false
	}

	r.Audiences[p.ID] = p

	return true
}

func (r *Room) RemoveAudience(p *player.Audience) {
	r.mu.Lock()
	defer r.mu.Unlock()

	delete(r.Audiences, p.ID)
}

func (r *Room) GetAudiences() map[string]*player.Audience {
	r.mu.Lock()
	defer r.mu.Unlock()

	return r.Audiences
}

func (r *Room) GetAudiencesCount() int {
	r.mu.Lock()
	defer r.mu.Unlock()

	return len(r.Audiences)
}

// audiences join seat
func (r *Room) JoinSeat(audiences map[string]*player.Audience) {
	r.mu.Lock()
	defer r.mu.Unlock()

	// 確定在房間裡的audience 才能被安排座位
	for _, audience := range audiences {
		if _, ok := r.Audiences[audience.ID]; ok {
			r.Game.Table.JoinSeat(audience)
		}
	}
}

// func (r *Room) StartGame() {
// 	r.mu.Lock()
// 	defer r.mu.Unlock()

// 	r.Game = game.NewGame(r.Players)
// }
