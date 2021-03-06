package rung_test

import (
	"testing"

	"github.com/minhajuddinkhan/pattay"
	"github.com/minhajuddinkhan/rung"
	"github.com/minhajuddinkhan/rung/dataset"
	"github.com/stretchr/testify/assert"
)

func TestGame_HasFourPlayers(t *testing.T) {

	game := rung.NewGame()
	assert.Equal(t, len(game.Players()), 4)
}

func TestGame_EachPlayerHasZeroCardsBeforeDistribution(t *testing.T) {

	game := rung.NewGame()
	players := game.Players()

	for _, player := range players {
		assert.Equal(t, len(player.CardsAtHand()), 0)
	}

}

func TestGame_EachPlayerHasThirteenCardsAfterDistribution(t *testing.T) {
	game := rung.NewGame()
	err := game.DistributeCards()
	assert.Nil(t, err)
	players := game.Players()
	for _, p := range players {
		assert.Equal(t, len(p.CardsAtHand()), 13)
	}

}

func TestGame_NoTwoPlayersHaveSameCard(t *testing.T) {

	game := rung.NewGame()
	err := game.DistributeCards()
	assert.Nil(t, err)
	players := game.Players()

	secondPlayer := players[1]

	cardWithfirstPlayer := players[0].CardsAtHand()[0]
	playerOneHasAceOfSpade := false
	playerTwoHasAceOfSpade := false

	for _, card := range secondPlayer.CardsAtHand() {
		if card.House() == cardWithfirstPlayer.House() && cardWithfirstPlayer.Number() == pattay.Ace {
			playerTwoHasAceOfSpade = true
		}
	}

	assert.NotEqual(t, playerOneHasAceOfSpade, playerTwoHasAceOfSpade)

}

func TestGame_FirstHandMustHaveFourCards(t *testing.T) {
	game := rung.NewGame()
	game.ShuffleDeck(20)
	assert.Nil(t, game.DistributeCards())
	players := game.Players()

	for _, p := range players {
		for i, c := range p.CardsAtHand() {
			if c.House() == pattay.Club {
				p.ThrowCard(i)
				break
			}
		}
	}

	handOutCome, err := game.PlayHand(0, nil, nil)
	assert.Nil(t, err)
	assert.Equal(t, len(handOutCome.Cards()), 4)
}

func TestGame_FirstHandMustHaveTwoOfClubs(t *testing.T) {

	game := rung.NewGame()
	game.ShuffleDeck(1)
	game.DistributeCards()
	p1, i1 := dataset.PlayerWithTwoOfClubs(game)
	others := dataset.PLayersWithoutTwoOfClubs(game)

	assert.True(t, p1.HasHouse(pattay.Club))
	assert.True(t, others[0].HasHouse(pattay.Club))
	assert.True(t, others[1].HasHouse(pattay.Club))
	assert.True(t, others[2].HasHouse(pattay.Club))

	p1.ThrowCard(i1)
	for _, px := range others {
		for j, c := range px.CardsAtHand() {
			if c.House() == pattay.Club {
				px.ThrowCard(j)
			}
		}
	}

	hand, err := game.PlayHand(0, nil, nil)
	assert.Nil(t, err)
	assert.Equal(t, len(hand.Cards()), 4)
	has, _ := hand.HasCard(pattay.NewCard(pattay.Club, pattay.Two))
	assert.True(t, has)
}

func TestGame_ConsecutiveHeadsPlayerShouldWinHandsAtTable(t *testing.T) {
	game := rung.NewGame()
	game.ShuffleDeck(20)
	game.DistributeCards()
	trump := pattay.Spade
	players := game.Players()

	var biggestPlayer pattay.Player

	var spades []pattay.Card
	for _, x := range players {
		spade, at, err := x.AnySpade()
		assert.Nil(t, err)
		spades = append(spades, spade)
		if spade.Number() == pattay.GetBiggestCard(spades, pattay.Spade).Number() {
			biggestPlayer = x
		}
		x.ThrowCard(at)
	}

	hand, err := game.PlayHand(1, &trump, biggestPlayer)
	assert.Nil(t, err)
	player, err := hand.Head()
	assert.Nil(t, err)
	assert.Equal(t, player.Name(), biggestPlayer.Name())
	assert.Equal(t, 0, len(game.HandsOnGround()), "Hands on ground should be zero after player has won hand")
	assert.Equal(t, 1, game.HandsWonBy(biggestPlayer), "Hands won be player should be 1")
}

func TestTwelvthHandShouldNotAddWinToAnyPlayer(t *testing.T) {

	game := rung.NewGame()
	game.ShuffleDeck(20)
	game.DistributeCards()

	trump := pattay.Spade
	biggestPlayerCard, aceAt := dataset.PlayerWithAceOfSpade(game)
	assert.NotEqual(t, -1, aceAt)
	others := dataset.PlayersWithoutAceOfSpade(game)

	biggestPlayerCard.ThrowCard(aceAt)
	for _, px := range others {
		for j, c := range px.CardsAtHand() {
			if c.House() == pattay.Spade {
				px.ThrowCard(j)
			}
		}
	}
	_, err := game.PlayHand(11, &trump, biggestPlayerCard)
	assert.Equal(t, 0, game.HandsWonBy(biggestPlayerCard))

	assert.Nil(t, err)

}

//TEST Todo: throw a card from player that it doens't have. expect error
//TEST todo: create invalid player in the game and call next
