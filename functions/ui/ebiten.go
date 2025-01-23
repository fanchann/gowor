package ui

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/fanchann/gowor/functions/tries"
)

const (
	ScreenWidth  = 720
	ScreenHeight = 480

	delay          = 30
	interval       = 3
	suggestionTime = 250 * time.Millisecond
)

type IWordDictionary interface {
	Update() error
	Draw(screen *ebiten.Image)
	Layout(outSideWidth, outsideHeight int) (int, int)

	repeatingKeyPressed(key ebiten.Key) bool
}

type WordDictionary struct {
	Runes       []rune
	Text        string
	LastText    string
	Counter     int
	Suggestions []string
	Meaning     string
	ShowPopup   bool
	LastUpdate  time.Time

	Trie *tries.Trie
}

func NewWordDictionary(textOnScreen string, trie *tries.Trie) IWordDictionary {
	return &WordDictionary{Text: textOnScreen, Trie: trie}
}
