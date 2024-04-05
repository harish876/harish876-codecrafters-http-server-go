package utils

type TrieNode struct {
	Children map[string]*TrieNode
	IsEnd    bool
}

func NewTrieNode() *TrieNode {
	return &TrieNode{
		Children: make(map[string]*TrieNode),
		IsEnd:    false,
	}
}

type TrieMethods interface {
	Insert(value rune)
}

type Trie struct {

}
