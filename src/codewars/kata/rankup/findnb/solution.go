package main

import "fmt"

var cubs = make(map[int]int)

func init() {
	n := 1
	m := 0
	for {
		m = n*n*n + m
		if m < 0 {
			break
		}
		cubs[m] = n
		n += 1
	}
}

func FindNb(m int) int {
	// your code
	if ans, found := cubs[m]; found {
		return ans
	} else {
		return -1
	}
}

func main() {
	fmt.Println(FindNb(4183059834009))
}
