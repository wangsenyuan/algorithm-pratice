package main

import "fmt"

func judgeCircle(moves string) bool {
	if len(moves) == 0 {
		return false
	}

	l, r, d, u := 0, 0, 0, 0

	for i := 0; i < len(moves); i++ {
		if moves[i] == 'U' {
			u++
		} else if moves[i] == 'D' {
			d++
		} else if moves[i] == 'L' {
			l++
		} else {
			r++
		}
	}

	return l == r && u == d
}

func main() {

	fmt.Println(judgeCircle("UD"))
	fmt.Println(judgeCircle("LL"))

	fmt.Println(judgeCircle("LDRRLRUULR"))
}
