package tries

import (
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func (t *Trie) Insert(word, meaning string) {
	t.mu.Lock()
	defer t.mu.Unlock()

	current := t.root
	for _, ch := range word {
		if _, exists := current.children[ch]; !exists {
			current.children[ch] = &TrieNode{children: make(map[rune]*TrieNode)}
		}
		current = current.children[ch]
	}
	current.isEnd = true
	current.meaning = meaning
}

func (t *Trie) SearchByWord(word string) []string {
	t.mu.RLock()
	defer t.mu.RUnlock()

	capitalizedWord := cases.Title(language.English).String(word)

	current := t.root

	for _, char := range capitalizedWord {
		if _, exists := current.children[char]; !exists {
			return nil 
		}
		current = current.children[char]
	}

	var results []string
	t.collectWords(current, word, &results)
	return results
}

func (t *Trie) Search(word string) (string, bool) {
	t.mu.RLock()
	defer t.mu.RUnlock()

	capitalizedWord := cases.Title(language.English).String(word)

	current := t.root
	for _, ch := range capitalizedWord {
		if _, exists := current.children[ch]; !exists {
			return "", false 
		}
		current = current.children[ch]
	}
	if current.isEnd {
		return current.meaning, true 
	}
	return "", false 
}

func (t *Trie) collectWords(node *TrieNode, prefix string, result *[]string) {
	if node.isEnd {
		*result = append(*result, prefix) 
	}

	for ch, child := range node.children {
		t.collectWords(child, prefix+string(ch), result)
	}
}
