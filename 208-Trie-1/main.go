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
	isEnd    bool
	children [26]*Trie
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	for _, v := range word {
		if this.children[v-'a'] == nil {
			this.children[v-'a'] = &Trie{}
		}
		this = this.children[v-'a']
	}
	this.isEnd = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	for _, v := range word {
		if this = this.children[v-'a']; this == nil {
			return false
		}
	}
	return this.isEnd
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	for _, v := range prefix {
		if this = this.children[v-'a']; this == nil {
			return false
		}
	}
	return true
}
