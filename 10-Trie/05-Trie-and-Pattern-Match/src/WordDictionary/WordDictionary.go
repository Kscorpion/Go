package main

import "fmt"

/// Leetcode 211. Add and Search Word - Data structure design
/// https://leetcode.com/problems/add-and-search-word-data-structure-design/description/
type Node struct {
	isWord bool
	next   map[string]*Node
}

type WordDictionary struct {
	root *Node
}

func getNode() *Node {
	return &Node{
		next: make(map[string]*Node),
	}
}

/** Initialize your data structure here. */
func Constructor() WordDictionary {
	return WordDictionary{
		root: getNode(),
	}
}

/** Adds a word into the data structure. */
func (this *WordDictionary) AddWord(word string) {
	cur := this.root

	for _, w := range []rune(word) {
		c := string(w)
		if cur.next[c] == nil {
			cur.next[c] = getNode()
		}
		cur = cur.next[c]
	}

	cur.isWord = true
}

/** Returns if the word is in the data structure. A word could contain the dot character '.' to represent any one letter. */
func (this *WordDictionary) Search(word string) bool {
	return this.match(this.root, word, 0)
}

func (this *WordDictionary) match(n *Node, word string, index int) bool {
	if index == len(word) {
		return n.isWord
	}

	c := string([]rune(word)[index])

	if c != "." {
		if n.next[c] == nil {
			return false
		}
		return this.match(n.next[c], word, index+1)
	} else {
		for w := range n.next {
			if this.match(n.next[w], word, index+1) {
				return true
			}
		}
		return false
	}
}

func main() {
	obj := Constructor()
	obj.AddWord("bad")
	obj.AddWord("dad")
	obj.AddWord("mad")

	fmt.Println(obj.Search("pad"))
	fmt.Println(obj.Search("bad"))
	fmt.Println(obj.Search(".ad"))
	fmt.Println(obj.Search("b.."))
}
