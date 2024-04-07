package utils

import (
	"fmt"
	"strings"
)

var (
	NO_MISMATCH = -1
)

type Node struct {
	Value  interface{}
	Edges  map[byte]*Edge
	IsLeaf bool
}

func NewNode(isLeaf bool) *Node {
	return &Node{
		Edges:  make(map[byte]*Edge),
		IsLeaf: isLeaf,
	}
}

type Edge struct {
	Label string
	Next  *Node
}

func NewEdge(label string) *Edge {
	return &Edge{
		Label: label,
		Next:  NewNode(true),
	}
}

func (n *Node) TotalEdges() int {
	return len(n.Edges)
}

func (n *Node) AddEdge(label string, next Node) {
	n.Edges[label[0]] = NewEdge(label)
}

func (n *Node) GetTransition(transitionCharacter byte) *Edge {
	return n.Edges[transitionCharacter]
}

type RadixTree struct {
	Root *Node
}

func NewRadixTree() RadixTree {
	return RadixTree{
		Root: NewNode(false),
	}
}

func (t *RadixTree) Insert(word string) {
	current := t.Root
	currIdx := 0
	for {
		if currIdx >= len(word) {
			break
		}

		transitionChar := word[currIdx]
		currentEdge := current.GetTransition(transitionChar)
		currStr := word[currIdx:]

		if currentEdge == nil {
			current.Edges[transitionChar] = NewEdge(currStr)
			break
		}

		splitIndex := GetFirstMismatchLetter(currStr, currentEdge.Label)
		if splitIndex == NO_MISMATCH {
			//The edge and leftover string are the same length
			//so finish and update the next node as a word node
			if len(currStr) == len(currentEdge.Label) {
				currentEdge.Next.IsLeaf = true
				break
			} else if len(currStr) < len(currentEdge.Label) {
				//The leftover word is a prefix to the edge string, so split
				suffix := currentEdge.Label[len(currStr):]
				currentEdge.Label = currStr
				newNext := NewNode(true)
				afterNewNext := currentEdge.Next
				currentEdge.Next = newNext
				newNext.AddEdge(suffix, *afterNewNext)
				break
			} else {
				splitIndex = len(currentEdge.Label)
			}
		} else {
			suffix := currentEdge.Label[splitIndex:]
			currentEdge.Label = currentEdge.Label[:splitIndex]
			prevNext := currentEdge.Next
			currentEdge.Next = NewNode(false)
			currentEdge.Next.AddEdge(suffix, *prevNext)
		}
		current = currentEdge.Next
		currIdx += splitIndex
	}
}

func (t *RadixTree) PrintAllWords() {
	printAllWords(t.Root, "")
}

func printAllWords(current *Node, result string) {
	if current.IsLeaf {
		fmt.Println(result)
	}

	for _, edge := range current.Edges {
		printAllWords(edge.Next, result+edge.Label)
	}
}

func (t *RadixTree) Search(word string) (*Node, bool) {
	current := t.Root
	currIndex := 0

	for {
		if currIndex >= len(word) {
			break
		}
		transitionChar := word[currIndex]
		edge := current.GetTransition(transitionChar)
		if edge == nil {
			return nil, false
		}
		currSubstring := word[currIndex:]
		if !strings.HasPrefix(currSubstring, edge.Label) {
			return nil, false
		}
		currIndex += len(edge.Label)
		current = edge.Next
	}
	if current.IsLeaf {
		return current, current.IsLeaf
	}
	return nil, current.IsLeaf
}

func GetFirstMismatchLetter(s1, s2 string) int {
	n1 := len(s1)
	n2 := len(s2)
	var minLen int

	if n1 > n2 {
		minLen = n2
	} else {
		minLen = n1
	}

	for i := 0; i < minLen; i++ {
		if s1[i] != s2[i] {
			return i
		}
	}
	return NO_MISMATCH
}
