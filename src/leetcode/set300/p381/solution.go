package main

import (
	"fmt"
	"math/rand"
)

func main() {
	set := Constructor()
	fmt.Println(set.Insert(-1))
	fmt.Println(set.Remove(-2))
	fmt.Println(set.Insert(-2))
	fmt.Println(set.GetRandom())
	fmt.Println(set.Remove(-1))
	fmt.Println(set.Insert(-2))
	fmt.Println(set.GetRandom())
}

type RandomizedCollection struct {
	vals map[int][]int
	idxs []int
}

/** Initialize your data structure here. */
func Constructor() RandomizedCollection {
	rand.Seed(41)
	return RandomizedCollection{make(map[int][]int), make([]int, 0)}
}

/** Inserts a value to the collection. Returns true if the collection did not already contain the specified element. */
func (this *RandomizedCollection) Insert(val int) bool {
	idx, found := this.vals[val]

	if idx == nil {
		idx = make([]int, 0, 1)
	}

	x := len(this.idxs)
	idx = append(idx, x)
	this.vals[val] = idx
	this.idxs = append(this.idxs, val)

	if found {
		return false
	}

	return true
}

/** Removes a value from the collection. Returns true if the collection contained the specified element. */
func (this *RandomizedCollection) Remove(val int) bool {
	idx, found := this.vals[val]
	if !found {
		return false
	}

	x := idx[0]
	y := idx[1:]
	z := len(this.idxs) - 1
	if x < z {
		a := this.idxs[z]
		b := this.vals[a]
		b = b[:len(b)-1]
		b = append(b, x)
		this.vals[a] = b
		this.idxs[x] = a
	}
	this.idxs = this.idxs[0:z]

	if len(y) == 0 {
		delete(this.vals, val)
	} else {
		this.vals[val] = y
	}

	return true
}

/** Get a random element from the collection. */
func (this *RandomizedCollection) GetRandom() int {
	selected := rand.Intn(len(this.idxs))
	return this.idxs[selected]
}
