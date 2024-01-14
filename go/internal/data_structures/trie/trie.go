package trie

type Trie struct {
	IsEnd bool
	Next  map[rune]*Trie
}

/** Initialize your data structure here. */
func NewTrie() Trie {
	return Trie{
		IsEnd: false,
		Next:  make(map[rune]*Trie),
	}
}

/** Inserts a word into the trie. */
func (t *Trie) Insert(word string) {
	for _, c := range word {
		if _, ok := t.Next[c]; !ok {
			t.Next[c] = &Trie{
				IsEnd: false,
				Next:  make(map[rune]*Trie),
			}
		}
		t = t.Next[c]
	}
	t.IsEnd = true
}

/** Returns if the word is in the trie. */
func (t *Trie) Search(word string) bool {
	for _, c := range word {
		if _, ok := t.Next[c]; !ok {
			return false
		}
		t = t.Next[c]
	}
	return t.IsEnd
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (t *Trie) StartsWith(prefix string) bool {
	for _, c := range prefix {
		if _, ok := t.Next[c]; !ok {
			return false
		}
		t = t.Next[c]
	}
	return true
}
