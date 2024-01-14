package template

import (
	"github.com/stretchr/testify/assert"
	"github.com/wildegor/daily_tasks/go/internal/data_structures/trie"
	"testing"
)

func Test_trie_1(t *testing.T) {
	trie := trie.NewTrie()
	trie.Insert("apple")
	assert.True(t, trie.Search("apple"))
	assert.False(t, trie.Search("app"))
	assert.True(t, trie.StartsWith("app"))
	trie.Insert("app")
	assert.True(t, trie.Search("app"))
}
