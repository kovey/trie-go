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
	length := len(keys)
	for i, key := range keys {
		if !node.Exists(key) {
			continue
		}

		node = node.Child(key)
		for j := i + 1; j < length; j++ {
			if !node.Exists(keys[j]) {
				break
			}
			node = node.Child(keys[j])
			if !node.End {
				continue
			}

			if i <= j {
				return true
			}
			i = j
			node = t.Root
			break
		}

		node = t.Root
	}

	return false
}

func (t *Trie) Replace(text string) string {
	if text == "" {
		return text
	}

	keys := []rune(text)
	node := t.Root
	var replaces []int = make([]int, 0)
	length := len(keys)
	for i, key := range keys {
		if !node.Exists(key) {
			continue
		}

		node = node.Child(key)
		for j := i + 1; j < length; j++ {
			if !node.Exists(keys[j]) {
				break
			}

			node = node.Child(keys[j])
			if !node.End {
				continue
			}

			for r := i; r <= j; r++ {
				replaces = append(replaces, r)
			}

			i = j
			node = t.Root
			break
		}
		node = t.Root
	}

	for _, i := range replaces {
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
