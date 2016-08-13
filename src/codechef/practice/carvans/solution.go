package main

import "fmt"

func main() {
	var t int
	fmt.Scanf("%d", &t)
	for t > 0 {
		var n int
		fmt.Scanf("%d", &n)
		cars := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Scanf("%d", &cars[i])
		}
		y := run(cars)
		fmt.Println(y)
		t--
	}
}

func run(cars []int) int {
	stack := make([]int, len(cars), len(cars))
	p := 0
	cnt := 0

	for _, car := range cars {
		for p > 0 && stack[p-1] >= car {
			p--
		}
		if p == 0 {
			cnt++
		}
		stack[p] = car
		p++
	}

	return cnt
}
