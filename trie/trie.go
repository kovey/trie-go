package trie

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

type Trie struct {
	Root *Node
}

func NewTrie() *Trie {
	return &Trie{Root: NewNode()}
}

func (t *Trie) Insert(text string) {
	text = strings.ReplaceAll(text, " ", "")
	if text == "" {
		return
	}

	if text == "你好" {
		fmt.Printf("text: %s", text)
	}

	keys := []rune(text)
	node := t.Root
	for _, key := range keys {
		if !node.Exists(key) {
			node.Children[key] = NewNode()
		}

		node = node.Child(key)
	}

	node.End = true
}

func (t *Trie) Has(text string) bool {
	text = strings.ReplaceAll(text, " ", "")
	if text == "" {
		return false
	}

	keys := []rune(text)
	node := t.Root
	for i, key := range keys {
		if !node.Exists(key) {
			if i == 0 {
				continue
			}

			node = t.Root
			if !node.Exists(key) {
				continue
			}
		}

		node = node.Child(key)
		if node.End {
			return true
		}
	}

	return false
}

func (t *Trie) Replace(text string) string {
	if text == "" {
		return text
	}

	keys := []rune(text)
	node := t.Root
	length := len(keys)
	replaces := make(map[int]bool)
	for i := 0; i < length; i++ {
		if !node.Exists(keys[i]) {
			continue
		}

		node = node.Child(keys[i])
		for j := i + 1; j < length; j++ {
			if !node.Exists(keys[j]) {
				i = j - 1
				break
			}

			node = node.Child(keys[j])
			if !node.End {
				continue
			}

			for r := i; r <= j; r++ {
				replaces[r] = true
			}
		}

		node = t.Root
	}

	for i := range replaces {
		c, _ := utf8.DecodeRuneInString("*")
		keys[i] = c
	}

	return string(keys)
}

func (t *Trie) Print() {
	for key := range t.Root.Children {
		fmt.Print("key:", string(key))
	}
}
