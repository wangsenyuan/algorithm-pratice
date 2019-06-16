package main

import (
	"fmt"
	"math"
)

func smallestFactorization(a int) int {
	if a == 1 {
		return 1
	}
	factors := make([]int, 32)

	var i = 0
	for x := 9; x > 1; x-- {
		for a%x == 0 && i < 32 {
			factors[i] = x
			i++
			a /= x
		}
	}

	if a > 1 {
		return 0
	}

	ans := int64(0)
	for j := i - 1; j >= 0; j-- {
		ans = ans*10 + int64(factors[j])
	}

	if ans > math.MaxInt32 {
		return 0
	}

	return int(ans)
}

func main() {
	//fmt.Println(smallestFactorization(48))
	//fmt.Println(smallestFactorization(15))
	fmt.Println(smallestFactorization(18000000))

}
