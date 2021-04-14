package p208

type Trie struct {
	children [26]*Trie
	leaf     bool
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	if len(word) == 0 {
		this.leaf = true
		return
	}
	x := int(word[0] - 'a')
	if this.children[x] == nil {
		this.children[x] = new(Trie)
	}
	this.children[x].Insert(word[1:])
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	if this == nil {
		return false
	}
	if len(word) == 0 {
		return this.leaf
	}
	x := int(word[0] - 'a')
	return this.children[x].Search(word[1:])
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	if this == nil {
		return false
	}
	if len(prefix) == 0 {
		return true
	}
	x := int(prefix[0] - 'a')
	return this.children[x].StartsWith(prefix[1:])
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
