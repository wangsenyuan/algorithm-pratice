package main

import "fmt"

type trie struct {
	children []*trie
	word     bool
}
type MagicDictionary struct {
	root *trie
}

/** Initialize your data structure here. */
func Constructor() MagicDictionary {
	return MagicDictionary{root: &trie{children: make([]*trie, 26)}}
}

func putIntoTrie(root *trie, word string) {
	if len(word) == 0 {
		root.word = true
		return
	}

	head := word[0]
	if root.children[head-'a'] == nil {
		root.children[head-'a'] = &trie{children: make([]*trie, 26)}
	}

	putIntoTrie(root.children[head-'a'], word[1:])
}

/** Build a dictionary through a list of words */
func (this *MagicDictionary) BuildDict(dict []string) {
	for _, word := range dict {
		putIntoTrie(this.root, word)
	}
}

func findInTrie(root *trie, word string, cnt int) bool {
	if len(word) == 0 || cnt > 1 {
		return root.word && cnt == 1
	}

	if root.children == nil {
		return false
	}

	for i := 0; i < 26; i++ {
		if root.children[i] != nil {
			var tmp bool
			if i == int(word[0]-'a') {
				tmp = findInTrie(root.children[i], word[1:], cnt)
			} else {
				tmp = findInTrie(root.children[i], word[1:], cnt+1)
			}
			if tmp {
				return true
			}
		}
	}
	return false
}

/** Returns if there is any word in the trie that equals to the given word after modifying exactly one character */
func (this *MagicDictionary) Search(word string) bool {
	if len(word) == 0 {
		return false
	}

	return findInTrie(this.root, word, 0)
}

/**
 * Your MagicDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.BuildDict(dict);
 * param_2 := obj.Search(word);
 */

func main() {
	dict := Constructor()
	dict.BuildDict([]string{"hello", "hallo", "leetcode"})
	fmt.Println(dict.Search("hello"))
	fmt.Println(dict.Search("hhllo"))
	fmt.Println(dict.Search("hell"))
	fmt.Println(dict.Search("leetcoded"))
	fmt.Println(dict.Search("hella"))
	fmt.Println(dict.Search("helloa"))

}
