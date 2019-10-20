package rung_test

import (
	"testing"

	"github.com/minhajuddinkhan/rung"
	"github.com/stretchr/testify/assert"
)

func TestDeckShouldHaveFiftyTwoCards(t *testing.T) {

	deck := rung.NewDeck()
	assert.Equal(t, len(deck.CardsInDeck()), 52)
}

func TestNewDeckHasFourOfSpades(t *testing.T) {

	deck := rung.NewDeck()
	found := true
	for _, card := range deck.CardsInDeck() {
		if card.House() == rung.Spade && card.Number() == 4 {
			found = true
		}
	}
	assert.True(t, found)
}

func TestIsCardPresentInDeck(t *testing.T) {

	deck := rung.NewDeck()
	card := rung.NewCard(rung.Spade, rung.Ace)
	assert.True(t, deck.IsCardPresent(card))

}
func TestIsCardNotPresentInDeck(t *testing.T) {

	deck := rung.NewDeck()
	card, err := deck.DrawCard(0)
	assert.Nil(t, err)
	assert.False(t, deck.IsCardPresent(card))
}

func TestAfterDrawingCardFromDeck(t *testing.T) {

	deck := rung.NewDeck()
	card, err := deck.DrawCard(0)
	assert.Nil(t, err)
	assert.False(t, deck.IsCardPresent(card))
	assert.Equal(t, len(deck.CardsInDeck()), 51)

}

func TestDrawCards(t *testing.T) {
	deck := rung.NewDeck()
	cards, err := deck.DrawCards(0, 2)
	assert.Nil(t, err)
	assert.False(t, deck.IsCardPresent(cards[0]))
	assert.False(t, deck.IsCardPresent(cards[1]))
	assert.False(t, deck.IsCardPresent(cards[2]))

}

func TestPutCard(t *testing.T) {

	deck := rung.NewDeck()
	card, err := deck.DrawCard(0)
	assert.Nil(t, err)
	deck.PutCard(card)
	assert.Equal(t, len(deck.CardsInDeck()), 52)
}

func TestPutCards(t *testing.T) {

	deck := rung.NewDeck()
	cards, err := deck.DrawCards(0, 1)
	assert.Nil(t, err)
	assert.Equal(t, len(deck.CardsInDeck()), 52-2)
	err = deck.PutCards(cards)
	assert.Nil(t, err)
	assert.Equal(t, len(deck.CardsInDeck()), 52)
}
func TestAfterShufflingDeck(t *testing.T) {

	deck := rung.NewDeck()
	err := deck.Shuffle(30)
	assert.Nil(t, err)
	assert.Equal(t, len(deck.CardsInDeck()), 52)

}

func TestDrawCardNotPresentInDeck(t *testing.T) {

	deck := rung.NewDeck()
	_, err := deck.DrawCard(52)
	assert.Error(t, err)
}

func TestDrawMoreCardsThenInDeck(t *testing.T) {
	deck := rung.NewDeck()
	_, err := deck.DrawCards(0, 53)
	assert.Error(t, err)
}

//test create deck, draw 50 cards, then shuffle. expect error
func TestDrawCardAndShuffle(t *testing.T) {
	deck := rung.NewDeck()
	deck.DrawCards(0, 51)
	assert.Error(t, deck.Shuffle(5))
}

func TestPutCardAlreadyPresent(t *testing.T) {
	deck := rung.NewDeck()
	err := deck.PutCard(rung.NewCard(rung.Spade, rung.Ace))
	assert.Error(t, err)
}
func TestPutMultipleCardAlreadyPresent(t *testing.T) {
	deck := rung.NewDeck()
	cards := []rung.Card{rung.NewCard(rung.Spade, rung.Ace)}
	err := deck.PutCards(cards)
	assert.Error(t, err)
}
func TestInvalidGetQueryCard(t *testing.T) {
	deck := rung.NewDeck()
	_, err := deck.DrawCards(2, 1)
	assert.Error(t, err)
}
