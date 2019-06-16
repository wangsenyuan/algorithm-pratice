package main

import (
	"fmt"
	"math/rand"
)

func main() {
	set := Constructor()
	fmt.Println(set.Insert(1000000000))
	fmt.Println(set.Insert(1000000001))
	fmt.Println(set.Insert(1000000002))

}

type RandomizedSet struct {
	vals map[int]int
	idxs []int
}

/** Initialize your data structure here. */
func Constructor() RandomizedSet {
	rand.Seed(41)
	return RandomizedSet{make(map[int]int), make([]int, 0, 100)}
}

/** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (this *RandomizedSet) Insert(val int) bool {
	if _, has := this.vals[val]; has {
		return false
	}
	idx := len(this.idxs)
	this.vals[val] = idx
	this.idxs = append(this.idxs, val)
	return true
}

/** Removes a value from the set. Returns true if the set contained the specified element. */
func (this *RandomizedSet) Remove(val int) bool {
	if _, has := this.vals[val]; !has {
		return false
	}

	idx := this.vals[val]
	if idx != len(this.idxs)-1 {
		last := this.idxs[len(this.idxs)-1]
		this.idxs[idx] = last
		this.vals[last] = idx
	}

	this.idxs = this.idxs[:len(this.idxs)-1]

	delete(this.vals, val)
	return true
}

/** Get a random element from the set. */
func (this *RandomizedSet) Getrandom() int {
	selected := rand.Intn(len(this.idxs))
	return this.idxs[selected]
}
