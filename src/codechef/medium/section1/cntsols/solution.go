package main

import (
	"bufio"
	"fmt"
	"os"
)

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
		line := readNNums(scanner, 4)
		fmt.Println(solve(line[0], line[1], line[2], line[3]))
		t--
	}
}

const MOD = 1000000007

func solve(U, D, M, N int) int64 {

	NN := int64(N)
	MM := int64(M)

	fn := func(x int, d int) int64 {
		y := int64(x)
		var res int64 = 1
		for d > 0 {
			if d&1 == 1 {
				res = (res * y) % NN
			}
			y = (y * y) % NN
			d >>= 1
		}
		return res
	}

	A := make([]int64, N)

	for i := 0; i < N; i++ {
		A[i] = fn(i, D)
	}

	// (a^^d)% N = (a % N)^^d

	count := func(x int) int {
		cnt := U / N
		if U%N > 0 && x <= U%N {
			return cnt + 1
		}
		return cnt
	}
	var ans int64

	for x := 0; x < N; x++ {
		for y := 0; y < N; y++ {
			for z := 0; z < N; z++ {
				sum := A[x] + A[y] + A[z]
				if sum%NN == MM {
					cx := int64(count(x)) % MOD
					cy := int64(count(y)) % MOD
					cz := int64(count(z)) % MOD
					tmp := (((cx * cy) % MOD) * cz) % MOD
					ans = (ans + tmp) % MOD
				}
			}
		}
	}

	return ans
}
