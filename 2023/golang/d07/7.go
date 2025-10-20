package d07

import (
	"aoc2023/src/util"
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/charmbracelet/log"
)

/*
	Author's note: I am not proud
*/

const (
	HIGH_CARD       = 0
	ONE_PAIR        = 1
	TWO_PAIR        = 2
	THREE_OF_A_KIND = 3
	FULL_HOUSE      = 4
	FOUR_OF_A_KIND  = 5
	FIVE_OF_A_KIND  = 6
)

type Hand struct {
	Cards    []byte
	Bid      int
	TypeRank int
}

func (h Hand) String() string {
	return fmt.Sprintf("\n\tHand[cards: %s, type: %d, jokers: %d, bid: %d]", string(h.Cards), h.TypeRank, CountJokers(h.Cards), h.Bid)
}

func Main() {
	data, err := util.ReadInput(7)
	if err != nil {
		log.Error("failed to read input")
		panic(err)
	}
	one(data)
	two(data)
}

func one(data []byte) {
	hands := ParseInputOne(data)
	sort.Slice(hands[:], func(i, j int) bool {
		first := hands[i]
		second := hands[j]
		if first.TypeRank != second.TypeRank {
			return first.TypeRank < second.TypeRank
		}
		pointer := 0
		for pointer < len(first.Cards) {
			firstCard := first.Cards[pointer]
			secondCard := second.Cards[pointer]
			if firstCard == secondCard {
				pointer++
				continue
			}
			firstCardRank := GetCardRank(first.Cards[pointer])
			secondCardRank := GetCardRank(second.Cards[pointer])
			return firstCardRank < secondCardRank
		}
		log.Error(fmt.Sprintf("unsortable %s %s", first, second))
		return false
	})

	sum := 0
	for i, hand := range hands {
		sum += (i + 1) * hand.Bid
	}
	log.Info(fmt.Sprintf("part 1 solution: %d", sum))
}

func two(data []byte) {
	hands := ParseInputTwo(data)
	sort.Slice(hands[:], func(i, j int) bool {
		first := hands[i]
		second := hands[j]
		if first.TypeRank != second.TypeRank {
			return first.TypeRank < second.TypeRank
		}
		pointer := 0
		for pointer < len(first.Cards) {
			firstCard := first.Cards[pointer]
			secondCard := second.Cards[pointer]
			if firstCard == secondCard {
				pointer++
				continue
			}
			firstCardRank := GetCardRankTwo(first.Cards[pointer])
			secondCardRank := GetCardRankTwo(second.Cards[pointer])
			return firstCardRank < secondCardRank
		}
		log.Error(fmt.Sprintf("unsortable %s %s", first, second))
		return false
	})

	sum := 0
	for i, hand := range hands {
		sum += (i + 1) * hand.Bid
	}
	log.Info(fmt.Sprintf("part 2 solution: %d", sum))
}

func GetCardRank(card byte) int {
	cardRanks := "123456789TJQKA"
	return strings.Index(cardRanks, string(card))
}

func GetCardRankTwo(card byte) int {
	cardRanks := "J123456789TQKA"
	return strings.Index(cardRanks, string(card))
}

func ParseInputOne(data []byte) []Hand {
	hands := []Hand{}

	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		segments := strings.Split(line, " ")
		hand := strings.TrimSpace(segments[0])
		bid, err := strconv.Atoi(strings.TrimSpace(segments[1]))
		if err != nil {
			log.Error(fmt.Sprintf("failed to parse %s to integer", segments[1]))
			panic(err)
		}
		hands = append(hands, NewHand([]byte(hand), bid))
	}
	return hands
}

func ParseInputTwo(data []byte) []Hand {
	hands := []Hand{}

	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		segments := strings.Split(line, " ")
		hand := strings.TrimSpace(segments[0])
		bid, err := strconv.Atoi(strings.TrimSpace(segments[1]))
		if err != nil {
			log.Error(fmt.Sprintf("failed to parse %s to integer", segments[1]))
			panic(err)
		}
		hands = append(hands, NewHandTwo([]byte(hand), bid))
	}
	return hands
}

func NewHand(cards []byte, bid int) Hand {
	return Hand{
		Cards:    cards,
		TypeRank: DetermineCardRank(cards),
		Bid:      bid,
	}
}

func NewHandTwo(cards []byte, bid int) Hand {
	return Hand{
		Cards:    cards,
		TypeRank: DetermineCardRankTwo(cards),
		Bid:      bid,
	}
}

func DetermineCardRank(cards []byte) int {
	cardMapping := map[byte]int{}
	for _, card := range cards {
		cardMapping[card]++
	}
	switch len(cardMapping) {
	case 0:
		return FIVE_OF_A_KIND
	case 1:
		return FIVE_OF_A_KIND
	case 2:
		// FOUR OF A KIND or FULL HOUSE
		for _, val := range cardMapping {
			if val == 4 {
				return FOUR_OF_A_KIND
			}
		}
		return FULL_HOUSE
	case 3:
		// TWO PAIR or THREE OF A KIND
		for _, val := range cardMapping {
			if val == 3 {
				return THREE_OF_A_KIND
			}
		}
		return TWO_PAIR
	case 4:
		return ONE_PAIR
	case 5:
		return HIGH_CARD
	}
	log.Error("impossible situation, hand of 5 cards has more than 5 different cards", "hand", cards)
	return -1
}

func CountJokers(cards []byte) int {
	count := 0
	for _, card := range cards {
		if card == 'J' {
			count++
		}
	}
	return count
}

var TypeRemapping map[int]map[int]int = map[int]map[int]int{
	HIGH_CARD: {
		1: ONE_PAIR,
		2: THREE_OF_A_KIND,
		3: FOUR_OF_A_KIND,
		4: FIVE_OF_A_KIND,
	},
	ONE_PAIR: {
		1: THREE_OF_A_KIND,
		2: FOUR_OF_A_KIND,
		3: FIVE_OF_A_KIND,
	},
	TWO_PAIR: {
		1: FULL_HOUSE,
		2: FOUR_OF_A_KIND,
	},
	THREE_OF_A_KIND: {
		1: FOUR_OF_A_KIND,
		2: FIVE_OF_A_KIND,
	},
	FOUR_OF_A_KIND: {
		1: FIVE_OF_A_KIND,
	},
}

func HandRankWithJoker(hand Hand) int {
	jokers := CountJokers(hand.Cards)
	if jokers == 0 || hand.TypeRank == FULL_HOUSE {
		return hand.TypeRank
	}
	if jokers == 5 {
		return FIVE_OF_A_KIND
	}

	return TypeRemapping[hand.TypeRank][jokers]
}

func FilterOutJokers(cards []byte) []byte {
	filtered := []byte{}
	for _, card := range cards {
		if card != 'J' {
			filtered = append(filtered, card)
		}
	}
	return filtered
}

func DetermineCardRankTwo(cards []byte) int {
	cards = FilterOutJokers(cards)
	jokers := 5 - len(cards)
	if jokers == 0 {
		return DetermineCardRank(cards)
	}
	if jokers == 5 {
		return FIVE_OF_A_KIND
	}

	cardMapping := map[byte]int{}
	for _, card := range cards {
		cardMapping[card]++
	}

	counts := []int{}
	for _, val := range cardMapping {
		counts = append(counts, val)
	}

	sort.Slice(counts[:], func(i, j int) bool {
		return counts[i] > counts[j]
	})

	if counts[0]+jokers == 5 {
		return FIVE_OF_A_KIND
	}

	switch counts[0] {
	case 3:
		return FOUR_OF_A_KIND
	case 2:
		switch jokers {
		case 1:
			if len(counts) > 1 && counts[1] == 2 {
				return FULL_HOUSE
			} else {
				return THREE_OF_A_KIND
			}
		case 2:
			return FOUR_OF_A_KIND
		}
	case 1:
		switch jokers {
		case 1:
			return ONE_PAIR
		case 2:
			return THREE_OF_A_KIND
		case 3:
			return FOUR_OF_A_KIND
		}
	}
	log.Error("uncaught scenario", "jokers", jokers, "hand", string(cards), "counts", counts)
	return 0
}
