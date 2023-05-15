package cards

import (
	"log"
	"main/structures/queue"
)

func Cards(cards int, q queue.Queue[int], name string) ([]int, int) {

	removedCards := addCardsToQueue(cards, q)
	removeRemainingCards(q, &removedCards)

	remainingCard, err := q.Dequeue()
	if err != nil {
		log.Fatal(err)
	}

	return removedCards, remainingCard
}

func addCardsToQueue(cards int, queue queue.Queue[int]) []int {
	removedCards := []int{}

	for i := 1; i <= cards; i += 2 {
		removedCards = append(removedCards, i)
		if i < cards {
			err := queue.Enqueue(i + 1)
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	return removedCards
}

func removeRemainingCards(q queue.Queue[int], removedCards *[]int) {
	for q.Len() > 1 {

		removeCard, err := q.Dequeue()
		if err != nil {
			log.Fatal(err)
		}
		*removedCards = append(*removedCards, removeCard)
		passCard, err := q.Dequeue()
		if err != nil {
			log.Fatal(err)
		}

		err = q.Enqueue(passCard)
		if err != nil {
			log.Fatal(err)
		}
	}
}
