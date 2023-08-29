package card

type Card struct {
	Suit string
	Rank string
}

var Suits = []string{"黑桃", "紅心", "方塊", "梅花"}
var Ranks = []string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K", "A"}

var SuitMap = map[string]int{
	"黑桃": 4,
	"紅心": 3,
	"方塊": 2,
	"梅花": 1,
}

// rank map
var RankMap = map[string]int{
	"2":  2,
	"3":  3,
	"4":  4,
	"5":  5,
	"6":  6,
	"7":  7,
	"8":  8,
	"9":  9,
	"10": 10,
	"J":  11,
	"Q":  12,
	"K":  13,
	"A":  14,
}

// 用撲克牌初始化
func NewCard(suit string, rank string) Card {
	return Card{
		Suit: suit,
		Rank: rank,
	}
}

// 比較兩張牌花色和點數的大小 1: card1 > card2 2: card1 < card2 0: card1 = card2
func CardCompareCards(card1 Card, card2 Card) int {
	// 先比較點數
	rankResult := rankCompare(card1.Rank, card2.Rank)
	if rankResult != 0 {
		return rankResult
	}

	// 點數相同再比較花色
	suitResult := suitCompare(card1.Suit, card2.Suit)
	if suitResult != 0 {
		return suitResult
	}

	return 0
}

// 比較兩張牌的點數的大小
func rankCompare(rank1 string, rank2 string) int {
	score1 := RankMap[rank1]
	score2 := RankMap[rank2]

	if score1 > score2 {
		return 1
	}
	if score1 < score2 {
		return 2
	}
	return 0

}

// 比較兩張牌的花色的大小
func suitCompare(suit1 string, suit2 string) int {
	score1 := SuitMap[suit1]
	score2 := SuitMap[suit2]

	if score1 > score2 {
		return 1
	}
	if score1 < score2 {
		return 2
	}
	return 0
}
