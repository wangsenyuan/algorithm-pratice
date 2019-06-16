package main

import (
	"fmt"
	"sort"
)

type PhoneDirectory struct {
	used     map[int]bool
	released []int
	index    int
	maxIndex int
}

/** Initialize your data structure here
  @param maxNumbers - The maximum numbers that can be stored in the phone directory. */
func Constructor(maxNumbers int) PhoneDirectory {
	return PhoneDirectory{make(map[int]bool), make([]int, 0), 0, maxNumbers}
}

/** Provide a number which is not assigned to anyone.
  @return - Return an available number. Return -1 if none is available. */
func (this *PhoneDirectory) Get() int {
	if len(this.used) == this.maxIndex {
		return -1
	}
	num := -1
	if len(this.released) == 0 {
		num = this.index
		this.index++
	} else {
		num = this.released[0]
		this.released = this.released[1:]
	}
	this.used[num] = true
	return num
}

/** Check if a number is available or not. */
func (this *PhoneDirectory) Check(number int) bool {
	return !this.used[number]
}

/** Recycle or release a number. */
func (this *PhoneDirectory) Release(number int) {
	if !this.used[number] {
		return
	}

	i := sort.Search(len(this.released), func(i int) bool {
		return this.released[i] > number
	})

	if i == len(this.released) {
		this.released = append(this.released, number)
	} else if i == 0 {
		this.released = append([]int{number}, this.released...)
	} else {
		a := this.released[0:i]
		c := make([]int, i, i+1)
		copy(c, a)
		c = append(c, number)
		this.released = append(c, this.released[i:]...)
	}

	delete(this.used, number)
}

func main() {
	// Init a phone directory containing a total of 3 numbers: 0, 1, and 2.
	directory := Constructor(3)

	// It can return any available phone number. Here we assume it returns 0.
	fmt.Println(directory.Get())

	// Assume it returns 1.
	fmt.Println(directory.Get())

	// The number 2 is available, so return true.
	fmt.Println(directory.Check(2))

	// It returns 2, the only number that is left.
	fmt.Println(directory.Get())

	// The number 2 is no longer available, so return false.
	fmt.Println(directory.Check(2))

	// Release number 2 back to the pool.
	directory.Release(2)

	// Number 2 is available again, return true.
	fmt.Println(directory.Check(2))
}
