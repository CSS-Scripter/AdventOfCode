package d07

import (
	"strings"
	"testing"
)

type DetermineCardRankTableTest struct {
	Hand     []byte
	Expected int
}

func TestDetermineCardRank(t *testing.T) {
	tableTests := []DetermineCardRankTableTest{
		{[]byte("TTTTT"), FIVE_OF_A_KIND},
		{[]byte("TTTAT"), FOUR_OF_A_KIND},
		{[]byte("A88AA"), FULL_HOUSE},
		{[]byte("KK992"), TWO_PAIR},
		{[]byte("9A78A"), ONE_PAIR},
		{[]byte("T1234"), HIGH_CARD},
	}

	for _, tt := range tableTests {
		actual := DetermineCardRank(tt.Hand)
		if actual != tt.Expected {
			t.Errorf("expected %d, but got %d", tt.Expected, actual)
		}
	}
}

func TestParseInput(t *testing.T) {
	exampleData := []byte(strings.Join([]string{
		"32T3K 765",
		"T55J5 684",
		"KK677 28",
		"KTJJT 220",
		"QQQJA 483",
	}, "\n"))

	expectedHands := []Hand{
		{[]byte("32T3K"), 765, ONE_PAIR},
		{[]byte("T55J5"), 684, THREE_OF_A_KIND},
		{[]byte("KK677"), 28, TWO_PAIR},
		{[]byte("KTJJT"), 220, TWO_PAIR},
		{[]byte("QQQJA"), 483, THREE_OF_A_KIND},
	}

	games := ParseInputOne(exampleData)

	for i, expected := range expectedHands {
		actual := games[i]
		if string(actual.Cards) != string(expected.Cards) {
			t.Errorf("expected %s, got %s", string(expected.Cards), string(actual.Cards))
		}
		if actual.Bid != expected.Bid {
			t.Errorf("expected %d, got %d", expected.Bid, actual.Bid)
		}
		if actual.TypeRank != expected.TypeRank {
			t.Errorf("expected %d, got %d", expected.TypeRank, actual.TypeRank)
		}
	}

}
