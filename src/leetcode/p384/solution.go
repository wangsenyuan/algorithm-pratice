package main

import (
	"fmt"
	"math/rand"
)

type Solution struct {
	nums []int
	last []int
}

func Constructor(nums []int) Solution {
	rand.Seed(41)
	last := make([]int, len(nums))
	copy(last, nums)
	return Solution{nums, last}
}

/** Resets the array to its original configuration and return it. */
func (this *Solution) Reset() []int {
	return this.nums
}

func removeSliceAt(nums []int, j int) []int {
	if j == 0 {
		return nums[1:]
	}

	if j == len(nums) - 1 {
		return nums[:j]
	}

	a := make([]int, j, len(nums) - 1)
	copy(a, nums[:j])
	return append(a, nums[j + 1:]...)
}

func pick(nums []int, result []int, i int) {
	if i == len(result) {
		return
	}
	//fmt.Printf("choose random from %v\n", nums)
	j := rand.Intn(len(nums))
	result[i] = nums[j]
	b := removeSliceAt(nums, j)
	//fmt.Printf("after remove %dth, get %v\n", j, b)
	pick(b, result, i + 1)
}

/** Returns a random shuffling of the array. */
func (this *Solution) Shuffle() []int {
	rs := make([]int, len(this.last))
	pick(this.nums, rs, 0)
	this.last = rs
	return rs
}

func main() {
	nums := []int{1, 2, 3}
	solution := Constructor(nums)
	fmt.Printf("shuffe => %v\n", solution.Shuffle())
	fmt.Printf("reset => %v\n", solution.Reset())
	fmt.Printf("shuffe => %v\n", solution.Shuffle())
	fmt.Printf("shuffe => %v\n", solution.Shuffle())
	fmt.Printf("shuffe => %v\n", solution.Shuffle())
}
