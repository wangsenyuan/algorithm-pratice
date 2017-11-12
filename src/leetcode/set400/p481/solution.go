package main

import "fmt"

func magicalString(n int) int {
	if n == 0 {
		return 0
	}
	if n <= 3 {
		return 1
	}

	str := make([]int, n)
	str[0] = 1
	str[1] = 2
	str[2] = 2

	cnt := 1
	p := 2
	i := 3
	for i < n {
		x := str[p]
		y := 3 - str[i - 1]
		str[i] = y

		if y == 1 {
			cnt++
		}

		i++

		if x == 2 && i < n {
			str[i] = y
			i++

			if y == 1 {
				cnt++
			}
		}
		p++
	}

	return cnt
}

func main() {
	//fmt.Println(magicalString(6))
	fmt.Println(magicalString(19))
	fmt.Println(magicalString(4))

}
