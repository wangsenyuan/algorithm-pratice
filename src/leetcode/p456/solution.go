package main

import "fmt"

func main() {
	fmt.Println(find132pattern([]int{1, 2, 3, 4}))         //false
	fmt.Println(find132pattern([]int{3, 1, 4, 2}))         //true
	fmt.Println(find132pattern([]int{-1, -1, 0, 1, 2, 0})) //true
	fmt.Println(find132pattern([]int{1, 0, 1, -4, -3}))    //false
	fmt.Println(find132pattern([]int{-2, 1, 2, -2, 1, 2})) //true
	fmt.Println(find132pattern([]int{2, 4, 1, 3}))         //true
	fmt.Println(find132pattern([]int{-2, 1, 1}))           //false
	fmt.Println(find132pattern([]int{4, 1, 3, 2}))         //true

}

func find132pattern(nums []int) bool {
	n := len(nums)
	xs := make([]int, n)
	stack := make([]int, n)
	p := 0
	for i, num := range nums {
		xs[i] = num
		if i > 0 && xs[i] > xs[i-1] {
			xs[i] = xs[i-1]
		}
		for p > 0 && nums[stack[p-1]] <= num {
			p--
		}
		if p > 0 {
			j := stack[p-1]
			if xs[j] < num {
				return true
			}
		}

		stack[p] = i
		p++
	}

	return false
}
