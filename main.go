package main

import (
	"embed"
	"log"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/fanchann/gowor/functions/tries"
	"github.com/fanchann/gowor/functions/ui"
	"github.com/fanchann/gowor/utils"
)

//go:embed dictionary/en_dictionary.txt
var EnDictionaryEmbedded embed.FS

func main() {
	// embedded dictionary
	dataEnDictionary, errReadEmbed := EnDictionaryEmbedded.ReadFile("dictionary/en_dictionary.txt")
	if errReadEmbed != nil {
		log.Fatal("Failed to read embedded dictionary:", errReadEmbed)
	}
	trie := tries.NewTrie()

	err := utils.LoadDictionaryFromEmbed(dataEnDictionary, &trie)
	if err != nil {
		log.Fatalf("Failed to load dictionary: %v", err)
	}

	g := ui.NewWordDictionary("Press Space to Show Meaning\nType On Keyboard\n", &trie)

	ebiten.SetWindowSize(ui.ScreenWidth, ui.ScreenHeight)
	ebiten.SetWindowTitle("OxFord Dictionary By Fanchann")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
