package main

import (
	"strconv"
	"fmt"
)

func main() {
	fmt.Println(complexNumberMultiply("1+1i", "1+1i"))
	fmt.Println(complexNumberMultiply("1+-1i", "1+-1i"))
}

func complexNumberMultiply(a string, b string) string {

	parse := func(str string) (int, int) {
		var x, y int
		i := 0
		for i < len(str) && str[i] != '+' {
			i++
		}
		if i == len(str) {
			x, _ = strconv.Atoi(str)
			return x, 0
		}

		x, _ = strconv.Atoi(str[:i])
		y, _ = strconv.Atoi(str[i+1: len(str)-1])
		return x, y
	}

	x, y := parse(a)
	u, v := parse(b)

	m := x*u - y*v
	n := x*v + u*y
	return fmt.Sprintf("%d+%di", m, n)
}
