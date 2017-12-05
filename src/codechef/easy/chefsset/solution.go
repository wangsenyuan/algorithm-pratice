package main

import (
	"bufio"
	"fmt"
	"os"
)

const N = 1000001

var primeFactors [][]int

func init() {
	primeFactors = make([][]int, N)

	primes := make([]int, N)

	for i := 2; i < N; i++ {
		if primes[i] == 0 {
			for j := i + i; j < N; j += i {
				primes[j] = 1
			}
		}
	}

	// fmt.Println(primes[:20])

	for i := 2; i < N; i++ {
		if primes[i] == 0 {

			for j := i; j < N; j += i {
				val := j
				var cnt int
				for val%i == 0 {
					val /= i
					cnt++
				}
				if cnt&1 == 1 {
					if primeFactors[j] == nil {
						primeFactors[j] = make([]int, 0, 5)
					}
					primeFactors[j] = append(primeFactors[j], i)
				}
			}
		}
	}
}

func readInt(bytes []byte, from int, val *int) int {
	i := from
	tmp := 0
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp
	return i
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var t int
	scanner.Scan()
	readInt(scanner.Bytes(), 0, &t)

	for i := 0; i < t; i++ {
		var n int
		scanner.Scan()
		readInt(scanner.Bytes(), 0, &n)
		scanner.Scan()
		bs := scanner.Bytes()
		x := -1
		ps := make(map[int]int)
		for j := 0; j < n; j++ {
			var num int
			x = readInt(bs, x+1, &num)
			factors := primeFactors[num]
			if len(factors) == 0 {
				continue
			}
			for _, d := range factors {
				ps[d]++
			}
		}
		// fmt.Println(ps)
		var ans int
		for _, c := range ps {
			ans += min(c, n-c)
		}
		fmt.Println(ans)
	}
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
