package main

import (
	"fmt"
	"strconv"
)

func main() {
	/*for i := 1; i < 100; i++ {
		fmt.Println(findNthDigit(i))
	}*/
	fmt.Println(findNthDigit(2))
}

func findNthDigit(n int) int {
	d := 1
	cnt := 9
	start := 1
	for n > d*cnt {
		n -= d * cnt
		start *= 10
		cnt *= 10
		d++
	}
	start += (n - 1) / d
	s := strconv.Itoa(start)
	return int(s[(n-1)%d] - '0')
}
