package cards

type actionLoop bool

const REMOVE_LOOP actionLoop = true
const MOVE_LOOP actionLoop = false

func CardsSliceLoop(cards int) ([]int, int) {
	cardsSlice := addCardsToSliceLoop(cards)
	removedCards, remainingCard := removeRemainingCardsLoop(cardsSlice, []int{}, REMOVE_LOOP)

	return removedCards, remainingCard
}

func addCardsToSliceLoop(cards int) []int {
	cardsSlice := []int{}

	for i := 1; i <= cards; i++ {
		cardsSlice = append(cardsSlice, i)
	}

	return cardsSlice
}

func removeRemainingCardsLoop(cardsSlice []int, removedCards []int, act actionLoop) ([]int, int) {
	remainingCards := []int{}

	for len(cardsSlice) > 1 {
		for _, card := range cardsSlice {

			if act == REMOVE_LOOP {
				removedCards = append(removedCards, card)
			} else {
				remainingCards = append(remainingCards, card)
			}

			act = !act
		}

		cardsSlice = remainingCards
		remainingCards = []int{}
	}

	return removedCards, cardsSlice[0]
}
