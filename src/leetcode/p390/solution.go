package main

import "fmt"

func main() {
	fmt.Println(lastRemaining(1))
	fmt.Println(lastRemaining(9))
	fmt.Println(lastRemaining(10))

}

func lastRemaining(n int) int {

	var process func(start, by, cnt int) int

	process = func(start, by, cnt int) int {
		if cnt == 1 {
			return start
		}

		end := start + by*(cnt-1)
		if cnt%2 == 1 {
			end -= by
		}
		return process(end, -2*by, cnt/2)
	}

	return process(1, 1, n)

}
