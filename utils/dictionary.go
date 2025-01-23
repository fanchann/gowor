package utils

import (
	"strings"

	"github.com/fanchann/gowor/functions/tries"
)

func GetSuggestions(trie *tries.Trie, input string) []string {
	// Dapatkan saran dari trie
	suggestions := trie.SearchByWord(input)

	// Batasi jumlah saran yang dikembalikan (maksimal 6)
	maxSuggestions := 6
	if len(suggestions) > maxSuggestions {
		suggestions = suggestions[:maxSuggestions]
	}

	return suggestions
}
func GetMeaning(trie *tries.Trie, word string) string {
	// Dapatkan arti kata dari trie
	meaning, found := trie.Search(word)
	if found {
		return meaning
	}
	return ""
}

// Fungsi untuk memecah teks menjadi beberapa baris berdasarkan jumlah maksimum karakter per baris
func WrapText(text string, maxCharsPerLine int) []string {
	var lines []string
	var currentLine string

	words := strings.Fields(text)
	for _, word := range words {
		// Cek apakah penambahan kata melebihi jumlah maksimum karakter per baris
		if len(currentLine)+len(word)+1 > maxCharsPerLine {
			// Jika melebihi, tambahkan baris saat ini ke lines dan mulai baris baru
			lines = append(lines, currentLine)
			currentLine = word
		} else {
			// Jika tidak, tambahkan kata ke baris saat ini
			if len(currentLine) > 0 {
				currentLine += " "
			}
			currentLine += word
		}
	}

	// Tambahkan baris terakhir
	if currentLine != "" {
		lines = append(lines, currentLine)
	}

	return lines
}
