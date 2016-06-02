package main

import (
	"fmt"
	"leetcode/p352/solution"
)

func main() {
	nums := []int{6, 6, 0, 4, 8, 7, 6, 4, 7, 5}
	sr := solution.Constructor()
	for _, num := range nums {
		sr.Addnum(num)
		is := sr.Getintervals()
		fmt.Printf("%v\n", is)
	}
}
