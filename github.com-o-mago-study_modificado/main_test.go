package main

import (
	"main/challenges/cards"
	"testing"
)

var N int = 100000

func BenchmarkCardSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cards.CardsSlice(N)
	}
}

func BenchmarkLinkedList(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cards.CardsLinkedList(N)
	}
}

func BenchmarkContainerList(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cards.CardsContainerList(N)
	}
}

func BenchmarkChannel(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cards.CardsChannel(N)
	}
}

func BenchmarkRecursive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cards.CardsRecursive(N)
	}
}

func BenchmarkSliceLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cards.CardsSliceLoop(N)
	}
}
