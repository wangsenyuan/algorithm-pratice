package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	nums := readFile("../p410/input1.txt")
	fmt.Println(splitArray(nums, 9))
}

func readFile(file string) []int {
	fp, e := filepath.Abs(file)
	dat, e := ioutil.ReadFile(fp)
	if e != nil {
		panic(e)
	}
	content := string(dat)
	strs := strings.Split(content, ",")
	nums := make([]int, len(strs))
	for i := range strs {
		x, _ := strconv.Atoi(strings.Trim(strs[i], " \n"))
		nums[i] = x
	}
	return nums
}

func splitArray(nums []int, m int) int {
	left := 0
	right := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] > left {
			left = nums[i]
		}
		right += nums[i]
	}

	checkAnswer := func(answer int) bool {
		n := 1
		s := nums[0]
		for i := 1; i < len(nums); i++ {
			s += nums[i]
			if s > answer {
				n++
				s = nums[i]
			}
		}
		return n <= m
	}

	for left <= right {
		if left == right {
			return left
		}
		if left+1 == right {
			if checkAnswer(left) {
				return left
			}
			if checkAnswer(right) {
				return right
			}
			return 0
		}
		mid := left + (right-left)/2
		if checkAnswer(mid) {
			right = mid
		} else {
			left = mid
		}
	}

	return 0
}
