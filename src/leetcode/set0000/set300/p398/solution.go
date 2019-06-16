package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strconv"
	"strings"
)

type Solution struct {
	nums []int
}

func Constructor(nums []int) Solution {
	rand.Seed(41)
	return Solution{nums}
}

func (this *Solution) Pick(target int) int {
	bound := 1
	res := 0
	for i := 0; i < len(this.nums); i++ {
		if this.nums[i] == target {
			if rand.Intn(bound) == 0 {
				res = i
			}
			bound++
		}
	}
	return res
}

func main() {

	dat, err := ioutil.ReadFile("/Users/wangsenyuan/Go/src/leetcode/p398/sample1.txt")
	if err != nil {
		panic(err)
	}

	str := string(dat)
	strs := strings.Split(str, ",")

	nums := make([]int, len(strs))

	for i, s := range strs {
		num, _ := strconv.Atoi(s)
		nums[i] = num
	}

	solution := Constructor(nums)

	fmt.Println(solution.Pick(-760627172))
}
