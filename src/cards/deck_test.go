package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length at 16, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got %v", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected last card of Four of Clubs, but got %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")
	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 16 {
		t.Errorf("Expected deck length at 16, but got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}

// ❯ go test
// --- FAIL: TestSaveToDeckAndNewDeckFromFile (0.00s)
//     deck_test.go:32: Expected deck length at 160, but got 16
// FAIL
// exit status 1
// FAIL    _/Users/y0y00rc/dev/mygo/src/cards      0.116s
// ❯ go test
// PASS
// ok      _/Users/y0y00rc/dev/mygo/src/cards      0.104s
