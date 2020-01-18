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

	line := readNNums(scanner, 3)
	n, k, s := line[0], line[1], line[2]
	A := readNNums(scanner, n)
	res := solve(n, k, s, A)
	fmt.Println(res)
}

func solve(N, k, s int, A []int) int64 {
	maxNum := A[0]
	for i := 1; i < N; i++ {
		if A[i] > maxNum {
			maxNum = A[i]
		}
	}

	minPrimeFactor := make([]int, maxNum+1)
	notPrime := make([]bool, maxNum+1)

	for x := 2; x <= maxNum; x++ {
		if notPrime[x] {
			continue
		}
		minPrimeFactor[x] = x
		for y := 2 * x; y <= maxNum; y += x {
			if minPrimeFactor[y] == 0 {
				minPrimeFactor[y] = x
			}
			notPrime[y] = true
		}
	}

	factors := make([]map[int]int, maxNum+1)

	for x := 1; x <= maxNum; x++ {
		factors[x] = make(map[int]int)

		y := x

		for y > 1 {
			factors[x][minPrimeFactor[y]]++
			y /= minPrimeFactor[y]
		}

	}
	K := int64(k)
	S := int64(s)
	var res int64

	for i := 0; i < N; i++ {
		fs := factors[A[i]]
		p := int64(len(fs))

		num := int64(A[i]) * (K - p*S)

		if num > res {
			res = num
		}
	}

	m := (k + s - 1) / s

	for p := 1; p < m; p++ {
		cnt := make(map[int]int)

		var sum int64
		var j int

		for i := 0; i < N; i++ {
			sum += int64(A[i])

			fs := factors[A[i]]
			for f, c := range fs {
				cnt[f] += c
			}
			if len(cnt) < p {
				continue
			}
			for len(cnt) > p {
				fs = factors[A[j]]
				for f, c := range fs {
					cnt[f] -= c
					if cnt[f] == 0 {
						delete(cnt, f)
					}
				}
				sum -= int64(A[j])
				j++
			}
			if len(cnt) == p {
				tmp := sum * (K - int64(p)*S)
				if tmp > res {
					res = tmp
				}
			}
		}
	}

	return res
}
