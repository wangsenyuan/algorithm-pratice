package main

import "fmt"

func main() {
	fmt.Println(circularArrayLoop([]int{2, -1, 1, 2, 2}))
	fmt.Println(circularArrayLoop([]int{-1, 2}))
	fmt.Println(circularArrayLoop([]int{1, 2, 0, 3, 4}))

}

func circularArrayLoop(nums []int) bool {
	if len(nums) < 2 {
		return false
	}

	for _, num := range nums {
		if num == 0 {
			return false
		}
	}

	n := len(nums)
	m := 2 * len(nums)
	detectLoop := func(forward bool) bool {
		for i, num := range nums {
			if num == 0 {
				continue
			}

			if forward && num < 0 || !forward && num > 0 {
				continue
			}

			one := i
			two := i
			for m > 0 {
				jumpOne := (one + nums[one] + n) % n

				if nums[one] == 0 || one == jumpOne || forward && nums[jumpOne] < 0 || !forward && nums[jumpOne] > 0 {
					break
				}
				one = jumpOne
				jumpTwo := (two + nums[two] + n) % n

				if nums[two] == 0 || two == jumpTwo || forward && nums[jumpTwo] < 0 || !forward && nums[jumpTwo] > 0 {
					break
				}

				two = jumpTwo
				jumpTwo = (two + nums[two] + n) % n
				if nums[two] == 0 || two == jumpTwo || forward && nums[jumpTwo] < 0 || !forward && nums[jumpTwo] > 0 {
					break
				}

				two = jumpTwo
				if one == two {
					return true
				}
				m--
			}

			j := i
			for nums[j] != 0 {
				tmp := nums[j]
				nums[j] = 0
				j = (j + tmp + n) % n
			}
		}
		return false
	}

	if detectLoop(true) || detectLoop(false) {
		return true
	}
	return false
}
