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

const INF = 1 << 30

func find132pattern(nums []int) bool {
	n := len(nums)
	// min value before n
	mn := make([]int, n)
	mn[0] = INF
	stack := make([]int, n)
	var p int
	var pivot = nums[0]
	for i := 1; i < n; i++ {
		cur := nums[i]
		for p > 0 && nums[stack[p-1]] <= cur {
			p--
		}
		// p == 0 || stack[p-1] > cur
		if p > 0 && mn[stack[p-1]] < cur {
			return true
		}
		if pivot < cur {
			stack[p] = i
			p++
			mn[i] = pivot
		} else {
			mn[i] = INF
			pivot = cur
		}
	}
	return false
}

func find132pattern1(nums []int) bool {
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
