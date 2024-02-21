package filter

import (
	"os"
	"strings"

	"github.com/kovey/trie-go/trie"
)

var keywords = trie.NewTrie()

func LoadByJson(path string) error {
	content := &Content{}
	if err := content.Load(path); err != nil {
		return err
	}

	Batch(content.Get())
	return nil
}

func Load(path, split string) error {
	content, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	Batch(strings.Split(string(content), split))

	return nil
}

func Batch(words []string) {
	for _, word := range words {
		keywords.Insert(word)
	}
}

func Has(text string) bool {
	return keywords.Has(text)
}

func Replace(text string) string {
	return keywords.Replace(text)
}
