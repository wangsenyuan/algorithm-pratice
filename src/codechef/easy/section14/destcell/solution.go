package main

import (
	"bufio"
	"fmt"
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

func readInt64(bytes []byte, from int, val *int64) int {
	i := from
	var tmp int64
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + int64(bytes[i]-'0')
		i++
	}
	*val = tmp
	return i
}

func readNInt64Nums(scanner *bufio.Scanner, n int) []int64 {
	res := make([]int64, n)
	x := -1
	scanner.Scan()
	for i := 0; i < n; i++ {
		x = readInt64(scanner.Bytes(), x+1, &res[i])
	}
	return res
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	tc := readNum(scanner)

	for tc > 0 {
		tc--
		n, m := readTwoNums(scanner)

		res := solve(n, m)

		for i := 0; i < len(res); i++ {
			if i < len(res)-1 {
				fmt.Printf("%d ", res[i])
			} else {
				fmt.Printf("%d\n", res[i])
			}
		}
	}
}

func solve(n int, m int) []int {
	K := n * m
	res := make([]int, K)
	res[0] = K
	for k := 2; k <= K; k++ {
		for x := 0; x < K; x += k {
			res[k-1]++
			c, r := x/n, x%n
			y := r*m + c
			if y%k != 0 {
				res[k-1]++
			}
		}

	}

	return res
}

func solve1(n int, m int) []int {
	K := n * m

	res := make([]int, K)

	for i := 0; i < K; i++ {
		res[i] = 1
	}
	//res[0] = K

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			x := i*m + j
			y := j*n + i

			a := 1

			for a*a <= x {
				if x%a == 0 {
					res[a-1]++
					b := x / a
					if a != b && b > 1 {
						res[b-1]++
					}
				}
				a++
			}

			a = 1

			for a*a <= y {
				if y%a == 0 {
					if x%a != 0 {
						res[a-1]++
					}
					b := y / a
					if a != b && b > 1 && x%b != 0 {
						res[b-1]++
					}
				}
				a++
			}
		}
	}

	return res
}
