package main

import (
	"bufio"
	"os"
	"fmt"
)

var factors [][]int

const N = 1000000

var primes []int

func init() {
	factors = make([][]int, N+1)

	for i := 0; i <= N; i++ {
		factors[i] = make([]int, 0, 10)
	}

	sieve := make([]bool, N+1)

	primes = make([]int, 78498)
	var k int
	for i := 2; i < N; i++ {
		if !sieve[i] {
			primes[k] = i
			factors[i] = append(factors[i], i)
			k++
			for j := i + i; j <= N; j += i {
				sieve[j] = true
				factors[j] = append(factors[j], i)
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

func readNum(scanner *bufio.Scanner) (a int) {
	scanner.Scan()
	readInt(scanner.Bytes(), 0, &a)
	return
}

func readTwoNums(scanner *bufio.Scanner) (a int, b int) {
	scanner.Scan()
	x := readInt(scanner.Bytes(), 0, &a)
	readInt(scanner.Bytes(), x+1, &b)
	return
}

func readNNums(scanner *bufio.Scanner, n int) []int {
	res := make([]int, n)
	x := -1
	scanner.Scan()
	for i := 0; i < n; i++ {
		x = readInt(scanner.Bytes(), x+1, &res[i])
	}
	return res
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	t := readNum(scanner)

	for t > 0 {
		n := readNum(scanner)
		A := readNNums(scanner, n)
		fmt.Println(solve(n, A))
		t--
	}
}

func solve(n int, A []int) int {
	p := make([]int, N+1)
	for i := 0; i <= N; i++ {
		p[i] = -1
	}
	cal := func(i int) int {
		j := -1

		for _, factor := range factors[A[i]] {
			if p[factor] > j {
				j = p[factor]
			}
			p[factor] = i
		}

		return i - j
	}
	cal(0)

	ans, prev := 1, 1
	for i := 1; i < n; i++ {
		prev = min(prev+1, cal(i))
		if prev > ans {
			ans = prev
		}
	}
	if ans > 1 {
		return ans
	}
	return -1
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
