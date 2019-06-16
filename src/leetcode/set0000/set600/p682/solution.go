package main

import (
	"strconv"
	"fmt"
)

func calPoints(ops []string) int {
	nums := make([]int, len(ops))
	p := 0

	res := 0
	for i := 0; i < len(ops); i++ {
		op := ops[i]
		if op == "C" {
			x := nums[p-1]
			res -= x
			p -= 1
		} else if op == "D" {
			x := 2 * nums[p-1]
			res += x
			nums[p] = x
			p++
		} else if op == "+" {
			a, b := nums[p-1], nums[p-2]
			x := a + b
			res += x
			nums[p] = x
			p++
		} else {
			x, _ := strconv.Atoi(op)
			res += x
			nums[p] = x
			p++
		}
	}

	return res
}

func main() {
	//fmt.Println(calPoints([]string{"5", "2", "C", "D", "+"}))
	fmt.Println(calPoints([]string{"5", "-2", "4", "C", "D", "9", "+", "+"}))
}
