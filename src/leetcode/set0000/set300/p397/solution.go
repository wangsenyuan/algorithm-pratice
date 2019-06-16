package main

import "fmt"

func main() {
	//fmt.Println(integerReplacement(8))
	//fmt.Println(integerReplacement(7))
	fmt.Println(integerReplacement(1024))
	fmt.Println(integerReplacement(1234))
	fmt.Println(integerReplacement(65535))
}

func integerReplacement(n int) int {
	res := 0
	for n != 1 {
		if n&1 == 0 {
			n >>= 1
		} else if n == 3 || (n>>1)&1 == 0 {
			n--
		} else {
			n++
		}
		res++
	}

	return res
}
