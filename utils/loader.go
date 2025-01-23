package utils

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
	"unicode"

	"github.com/fanchann/gowor/functions/tries"
)

func LoadDictionariesIntoTrie(fileName string, trie *tries.Trie) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var currentWord, currentMeaning string

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		if len(line) > 0 && unicode.IsUpper(rune(line[0])) {
			if currentWord != "" {
				trie.Insert(currentWord, currentMeaning)
			}

			parts := strings.SplitN(line, " ", 2)
			if len(parts) < 2 {
				currentWord = parts[0]
				currentMeaning = ""
			} else {
				currentWord = parts[0]
				currentMeaning = strings.TrimSpace(parts[1])
			}
		} else {
			currentMeaning += " " + strings.TrimSpace(line)
		}
	}

	if currentWord != "" {
		trie.Insert(currentWord, currentMeaning)
	}

	return scanner.Err()
}

func LoadDictionaryFromEmbed(data []byte, trie *tries.Trie) error {
	scanner := bufio.NewScanner(bytes.NewReader(data))
	var currentWord, currentMeaning string

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		if len(line) > 0 && unicode.IsUpper(rune(line[0])) {
			if currentWord != "" {
				trie.Insert(currentWord, currentMeaning)
			}

			// Pisahkan kata dan arti
			parts := strings.SplitN(line, " ", 2)
			if len(parts) < 2 {
				currentWord = parts[0]
				currentMeaning = ""
			} else {
				currentWord = parts[0]
				currentMeaning = strings.TrimSpace(parts[1])
			}
		} else {
			currentMeaning += " " + strings.TrimSpace(line)
		}
	}

	if currentWord != "" {
		trie.Insert(currentWord, currentMeaning)
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("gagal membaca data: %w", err)
	}

	return nil
}
