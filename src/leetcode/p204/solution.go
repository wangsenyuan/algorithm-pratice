package main

import "fmt"

func main() {
	//fmt.Println(countPrimes(101))
	fmt.Println(countPrimes(0))
	fmt.Println(countPrimes(2))
	fmt.Println(countPrimes(3))
	fmt.Println(countPrimes(10))

}

func countPrimes(n int) int {
	if n < 2 {
		return 0
	}
	ps := make([]bool, n)
	ps[1] = true

	p := 2
	cnt := n - 2
	for p*p < n {
		for i := 2; p*i < n; i++ {
			if !ps[p*i] {
				cnt--
			}
			ps[p*i] = true
		}
		i := p + 1
		for i*i < n {
			if !ps[i] {
				break
			}
			i++
		}
		p = i
	}

	return cnt
}
