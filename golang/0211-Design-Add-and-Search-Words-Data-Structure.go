package main

import "fmt"

func main() {
	wd := Constructor()
	wd.AddWord("bad")
	wd.AddWord("dad")
	wd.AddWord("mad")
	fmt.Println(wd.Search("pad"))
	fmt.Println(wd.Search("bad"))
	fmt.Println(wd.Search(".ad"))
	fmt.Println(wd.Search("b.."))
}

type WordDictionary struct {
	children [26]*WordDictionary
	isEnd    bool
}

func Constructor() WordDictionary {
	return WordDictionary{}
}

func (this *WordDictionary) AddWord(word string) {
	for i := 0; i < len(word); i++ {
		index := word[i] - 'a'
		//fmt.Println(index)
		if this.children[index] == nil {
			this.children[index] = &WordDictionary{}
		}
		this = this.children[index]
		//fmt.Println(this)
		//fmt.Println("---")
	}
	this.isEnd = true
	//fmt.Println(this)
	//fmt.Println("===============")
}

func (this *WordDictionary) Search(word string) bool {
	if len(word) == 0 {
		return this.isEnd
	}
	fmt.Println(this)
	if word[0] == '.' {
		// try all possible nodes
		for j := 0; j < 26; j++ {
			if this.children[j] != nil && this.children[j].Search(word[1:]) {
				return true
			}
		}
	} else {
		index := word[0] - 'a'
		// if it has child, keep searching
		if child := this.children[index]; child != nil {
			return child.Search(word[1:])
		}
	}
	return false
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
