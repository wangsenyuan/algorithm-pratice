package main

import "fmt"

func main() {
	//input := "2-1-1"
	input := "2*3-4*5"
	fmt.Println(diffWaysToCompute(input))
}

func diffWaysToCompute(input string) []int {
	cache := make(map[string][]int)

	var compute func(s string) []int

	compute = func(s string) []int {
		if nums, ok := cache[s]; ok {
			return nums
		}

		var nums []int

		isNum := true
		x := 0
		for i := 0; i < len(s); i++ {
			if s[i] == '+' || s[i] == '-' || s[i] == '*' {
				isNum = false
				left := compute(s[:i])
				right := compute(s[i+1:])
				nums = update(nums, left, right, s[i])
			}
			if isNum {
				x = x*10 + int(s[i]-'0')
			}
		}

		if isNum {
			nums = append(nums, x)
		}
		cache[s] = nums

		return nums
	}

	return compute(input)
}

func update(nums, left, right []int, op byte) []int {
	for _, a := range left {
		for _, b := range right {
			switch op {
			case '+':
				nums = append(nums, a+b)
			case '-':
				nums = append(nums, a-b)
			case '*':
				nums = append(nums, a*b)
			}
		}
	}

	return nums
}
