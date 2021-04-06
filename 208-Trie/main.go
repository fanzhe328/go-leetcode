package main

import (
	"fmt"
)

func main() {
	obj := Constructor()
	obj.Insert("a")

	fmt.Println(obj.Search("a"))
	fmt.Println(obj.StartsWith("a"))
}

type Trie struct {
	IsEnd    bool
	Children map[byte]*Trie
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{
		Children: make(map[byte]*Trie),
	}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	cur := this
	for i := 0; i < len(word); i++ {
		if _, ok := cur.Children[word[i]]; !ok {
			cur.Children[word[i]] = &Trie{
				Children: make(map[byte]*Trie),
			}
		}
		cur = cur.Children[word[i]]
	}
	cur.IsEnd = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	cur := this.find(word)
	if cur == nil {
		return false
	}
	return cur.IsEnd
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	cur := this.find(prefix)
	if cur == nil {
		return false
	}
	if cur.IsEnd || len(cur.Children) != 0 {
		return true
	}
	return false
}

func (this *Trie) find(prefix string) *Trie {
	cur := this
	for i := 0; i < len(prefix); i++ {
		if _, ok := cur.Children[prefix[i]]; !ok {
			return nil
		}
		cur = cur.Children[prefix[i]]
	}
	return cur
}
