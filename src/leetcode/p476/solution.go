package main

import "fmt"

func findComplement(num int) int {
	if num == 0 {
		return 1
	}

	ans := 0

	for i := 0; num > 0; i++ {
		if num&1 == 0 {
			ans |= 1 << uint(i)
		}
		num = num >> 1
	}

	return ans
}

func main() {
	fmt.Println(findComplement(5))
	fmt.Println(findComplement(1))
	fmt.Println(findComplement(0))
	fmt.Println(findComplement(8))

}
