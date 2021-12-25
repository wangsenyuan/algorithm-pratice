package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func readInt(bytes []byte, from int, val *int) int {
	i := from
	sign := 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	tmp := 0
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
}

func readNum(scanner *bufio.Scanner) (a int) {
	scanner.Scan()
	readInt(scanner.Bytes(), 0, &a)
	return
}

func readTwoNums(scanner *bufio.Scanner) (a int, b int) {
	res := readNNums(scanner, 2)
	a, b = res[0], res[1]
	return
}

func readNNums(scanner *bufio.Scanner, n int) []int {
	res := make([]int, n)
	x := 0
	scanner.Scan()
	for i := 0; i < n; i++ {
		for x < len(scanner.Bytes()) && scanner.Bytes()[x] == ' ' {
			x++
		}
		x = readInt(scanner.Bytes(), x, &res[i])
	}
	return res
}

func fillNNums(scanner *bufio.Scanner, n int, res []int) {
	x := 0
	scanner.Scan()
	for i := 0; i < n; i++ {
		for x < len(scanner.Bytes()) && scanner.Bytes()[x] == ' ' {
			x++
		}
		x = readInt(scanner.Bytes(), x, &res[i])
	}
}

func readUint64(bytes []byte, from int, val *uint64) int {
	i := from

	var tmp uint64
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + uint64(bytes[i]-'0')
		i++
	}
	*val = tmp

	return i
}

func main() {
	// for i := 2; i < 200; i++ {
	// 	fmt.Printf("prime factor count: %d => %d\n", i, primeFactorsCount[i])
	// }
	scanner := bufio.NewScanner(os.Stdin)

	tc := readNum(scanner)
	for tc > 0 {
		tc--
		n := readNum(scanner)
		A := readNNums(scanner, n)
		fmt.Println(solve(n, A))
	}
}

const MAX_NUM = 100010

var primes []int
var primeFactorsCount []int
var count []int64

func init() {
	set := make([]bool, MAX_NUM)
	smf := make([]int, MAX_NUM)
	var j int
	primes = make([]int, MAX_NUM)
	for x := 2; x < MAX_NUM; x++ {
		if !set[x] {
			primes[j] = x
			smf[x] = x
			j++
			for y := x * x; y < MAX_NUM && y > 0; y += x {
				set[y] = true
				if smf[y] == 0 {
					smf[y] = x
				}
			}
		}
	}
	primes = primes[:j]

	primeFactorsCount = make([]int, MAX_NUM)

	for num := 2; num < MAX_NUM; num++ {
		fs := make(map[int]bool)
		x := num
		for x >= 2 {
			fs[smf[x]] = true
			x /= smf[x]
		}
		primeFactorsCount[num] = len(fs)
	}

	count = make([]int64, MAX_NUM)
}

func solve(n int, A []int) int64 {
	for i := 0; i < MAX_NUM; i++ {
		count[i] = 0
	}
	for i := 0; i < n; i++ {
		num := A[i]
		count[1]++
		if num > 1 {
			count[num]++
		}
		sqr := int(math.Sqrt(float64(num)))
		for x := 2; x <= sqr; x++ {
			if num%x == 0 {
				count[x]++
				if x < sqr {
					count[num/x]++
				}
			}
		}

	}

	var best int64

	for i := 2; i < MAX_NUM; i++ {
		tmp := count[i] * int64(primeFactorsCount[i])
		if tmp > best {
			best = tmp
		}
	}

	return best
}
