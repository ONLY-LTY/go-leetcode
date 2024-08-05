package graph

// Trie 前缀树
type Trie struct {
	isWorld  bool
	children map[rune]*Trie
}

func Constructor() Trie {
	return Trie{isWorld: false, children: make(map[rune]*Trie)}
}

func (t *Trie) Insert(word string) {
	parent := t
	for _, ch := range word {
		if child, ok := parent.children[ch]; ok {
			parent = child
		} else {
			newChild := &Trie{children: make(map[rune]*Trie)}
			parent.children[ch] = newChild
			parent = newChild
		}
	}
	parent.isWorld = true
}

func (t *Trie) Search(word string) bool {
	parent := t
	for _, ch := range word {
		if child, ok := parent.children[ch]; ok {
			parent = child
			continue
		}
		return false
	}
	return parent.isWorld
}

func (t *Trie) StartsWith(prefix string) bool {
	parent := t
	for _, ch := range prefix {
		if child, ok := parent.children[ch]; ok {
			parent = child
			continue
		}
		return false
	}
	return true
}
