package cards

type action bool

const REMOVE action = true
const MOVE action = false

func CardsRecursive(cards int) ([]int, int) {
	cardsSlice := addCardsToSlice(cards)
	removedCards, remainingCard := removeRemainingCardsRecursively(cardsSlice, []int{}, REMOVE)

	return removedCards, remainingCard
}

func addCardsToSlice(cards int) []int {
	cardsSlice := []int{}

	for i := 1; i <= cards; i++ {

		cardsSlice = append(cardsSlice, i)
	}

	return cardsSlice
}

func removeRemainingCardsRecursively(cardsSlice []int, removedCards []int, act action) ([]int, int) {
	if len(cardsSlice) == 1 {
		return removedCards, cardsSlice[0]
	}

	remainingCards := []int{}

	for _, card := range cardsSlice {

		if act == REMOVE {
			removedCards = append(removedCards, card)
		} else {
			remainingCards = append(remainingCards, card)
		}

		act = !act
	}

	return removeRemainingCardsRecursively(remainingCards, removedCards, act)
}
