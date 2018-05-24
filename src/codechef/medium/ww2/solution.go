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

func readTwoNum(scanner *bufio.Scanner) (a int, b int) {
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
		n, m := readTwoNum(scanner)
		fmt.Println(solve(n, m))
		t--
	}
}

const MOD = 1000000007

func solve(n, m int) int64 {
	if m == 1 {
		if n == 1 {
			return 1
		}
		return 0
	}

	even := newMatrix(m)
	odd := newMatrix(m)

	for i := 0; i < m; i++ {
		odd[i][i] = 1
		if i > 0 {
			even[i][i-1] = 1
			odd[i][i-1] = 1
		}
		if i < m-1 {
			even[i][i+1] = 1
			odd[i][i+1] = 1
		}
	}
	t := multiply(even, odd)

	f := fastPow(t, (n-1)/2)
	if n%2 == 0 {
		f = multiply(f, even)
	}
	var ans int64
	for i := 0; i < m; i++ {
		for j := 0; j < m; j++ {
			ans += f[i][j]
			ans %= MOD
		}
	}
	return ans
}

type matrix [][]int64

func newMatrix(k int) matrix {
	c := make(matrix, k)
	for i := 0; i < k; i++ {
		c[i] = make([]int64, k)
	}
	return c
}

func multiply(a, b matrix) matrix {
	k := len(a)
	c := newMatrix(k)

	for i := 0; i < k; i++ {
		for j := 0; j < k; j++ {
			for x := 0; x < k; x++ {
				c[i][j] += (a[i][x] * b[x][j]) % MOD
				c[i][j] %= MOD
			}
		}
	}
	return c
}

func identity(k int) matrix {
	c := newMatrix(k)
	for i := 0; i < k; i++ {
		c[i][i] = 1
	}
	return c
}

func fastPow(t matrix, n int) matrix {
	if n == 0 {
		return identity(len(t))
	}

	if n == 1 {
		return t
	}

	ret := fastPow(t, n/2)
	ret = multiply(ret, ret)
	if n%2 == 0 {
		return ret
	}
	return multiply(ret, t)
}
