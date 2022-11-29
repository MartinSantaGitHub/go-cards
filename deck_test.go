package main

import (
	"os"
	"testing"
)

func getDeck() deck {
	return newDeck()
}

func TestNewDeckCorrectOrder(t *testing.T) {
	d := getDeck()

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card Ace of Spades, but got %s", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected last card of Four of Clubs, but got %s", d[len(d)-1])
	}
}

func TestNewDeckLength(t *testing.T) {
	d := getDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %d", len(d))
	}
}

func TestSaveToFile(t *testing.T) {
	d := getDeck()

	d.saveToFile("_decktesting.txt")
}

func TestNewDeckFromFile(t *testing.T) {
	d := newDeckFromFile("_decktesting.txt")

	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %d", len(d))
	}

	os.Remove("_decktesting.txt")
}
