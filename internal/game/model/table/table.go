package table

import (
	"gameserver/internal/game/model/card"
	"gameserver/internal/game/model/player"
	"math/rand"
	"time"
)

type Table struct {
	ID    int
	Cards []card.Card
	rng   *rand.Rand
	Seats []*Seat
}

type Seat struct {
	Player *player.Player
}

var maxSeat = 100

// NewTable 新建桌子
func NewTable() *Table {

	table := &Table{
		ID:    1,
		Cards: []card.Card{},
	}

	source := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(source)

	table.Cards = table.OriginCards()
	table.rng = rng
	table.initSeats()

	return table
}

// 原始牌组
func (d *Table) OriginCards() []card.Card {
	var cards []card.Card
	for _, suit := range card.Suits {
		for _, rank := range card.Ranks {
			cards = append(cards, card.NewCard(suit, rank))
		}
	}
	return cards
}

func (d *Table) Shuffle() {
	d.rng.Shuffle(len(d.Cards), func(i, j int) {
		d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i]
	})
}

// 发牌
func (d *Table) Deal(num int) []card.Card {
	cards := d.Cards[:num]
	d.Cards = d.Cards[num:]
	return cards
}

func (t *Table) initSeats() {
	// 初始化座位
	for i := 0; i < maxSeat; i++ {
		t.Seats = append(t.Seats, &Seat{})
	}
}

// 隨機安排座位
func (t *Table) ArrangeSeats(audiences map[string]*player.Audience) {

	for _, audience := range audiences {
		t.JoinSeat(audience)
	}
}

// JoinSeat 加入座位
func (t *Table) JoinSeat(audience *player.Audience) {

	player := player.NewPlayer(audience)

	for _, seat := range t.Seats {
		if seat.Player == nil {
			seat.Player = player
			break
		}
	}
}

// LeaveSeat 离开座位
func (t *Table) LeaveSeat(p *player.Player) {
	for _, seat := range t.Seats {
		if seat.Player == p {
			seat.Player = nil
			break
		}
	}
}

// GetActiveSeats 获取活跃座位
func (t *Table) GetActiveSeats() []*Seat {
	var activeSeats []*Seat
	for _, seat := range t.Seats {
		if seat.Player == nil {
			continue
		}
		if !seat.Player.Status.IsOut() {
			activeSeats = append(activeSeats, seat)
		}
	}
	return activeSeats
}

// 取得所有存活玩家
func (t *Table) GetActivePlayers() []*player.Player {
	var activePlayers []*player.Player
	for _, seat := range t.Seats {
		if seat.Player == nil {
			continue
		}
		if !seat.Player.Status.IsOut() {
			activePlayers = append(activePlayers, seat.Player)
		}
	}
	return activePlayers
}

// 所有存活玩家發一張牌
// func (t *Table) Deal() {
// 	for _, seat := range t.GetActiveSeats() {
// 		// seat.Player.Deal()
// 		fmt.Print(seat.Player.Username)
// 	}
// }
