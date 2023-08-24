package card

type Card struct {
	Suit string
	Rank string
}

var Suits = []string{"黑桃", "紅心", "方塊", "梅花"}
var Ranks = []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}

// 用撲克牌初始化
func NewCard(suit string, rank string) Card {
	return Card{
		Suit: suit,
		Rank: rank,
	}
}

// 比較兩張牌花色和點數的大小
func CardCompareCards(card1 Card, card2 Card) int {
	// 比较两张牌的大小
	if card1.Suit == card2.Suit {
		return rankCompare(card1.Rank, card2.Rank)
	} else {
		return suitCompare(card1.Suit, card2.Suit)
	}
}

// 比較兩張牌的點數的大小
func rankCompare(rank1 string, rank2 string) int {
	// 比较两张牌的大小
	var rank1Index int
	var rank2Index int
	for index, rank := range Ranks {
		if rank1 == rank {
			rank1Index = index
		}
		if rank2 == rank {
			rank2Index = index
		}
	}
	if rank1Index > rank2Index {
		return 1
	} else if rank1Index < rank2Index {
		return -1
	} else {
		return 0
	}
}

// 比較兩張牌的花色的大小
func suitCompare(suit1 string, suit2 string) int {
	// 比较两张牌的大小

	var suit1Index int
	var suit2Index int
	for index, suit := range Suits {
		if suit1 == suit {
			suit1Index = index
		}
		if suit2 == suit {
			suit2Index = index
		}
	}
	if suit1Index > suit2Index {
		return 1
	} else if suit1Index < suit2Index {
		return -1
	} else {
		return 0
	}
}
