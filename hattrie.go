package safebrowsing

import "github.com/apokalyptik/quicktrie"

type HatTrie struct {
	trie *trie.Trie
}

func NewTrie() *HatTrie {
	out := &HatTrie{
		trie: trie.NewTrie(),
	}
	return out
}

func (h *HatTrie) Delete(key string) {
	h.trie.Del(key)
}

func (h *HatTrie) Set(key string) {
	h.trie.Add(key)
}

func (h *HatTrie) Get(key string) bool {
	return h.trie.Get(key)
}
