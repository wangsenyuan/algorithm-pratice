package main

import (
	"strconv"
	"fmt"
)

func main() {
	fmt.Println(convertTo7(100))
	fmt.Println(convertTo7(-7))

}

func convertTo7(num int) string {
	if num == 0 {
		return "0"
	}

	sign := num >= 0

	if num < 0 {
		num = -num
	}

	var res string

	for num > 0 {
		x := num % 7
		res = strconv.Itoa(x) + res
		num = num / 7
	}

	if !sign {
		return "-" + res
	}
	return res
}
