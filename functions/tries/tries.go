package tries

import "sync"

type ITrie interface {
	Insert(word string)
	SearchByWord(word string) []string
	collectWord(node *TrieNode, prefix string, result *[]string)
}

type TrieNode struct {
	children map[rune]*TrieNode
	isEnd    bool
	meaning string
}

type Trie struct {
	root *TrieNode
	mu   sync.RWMutex
}

func NewTrie() Trie {
	return Trie{root: &TrieNode{children: make(map[rune]*TrieNode)}}
}
