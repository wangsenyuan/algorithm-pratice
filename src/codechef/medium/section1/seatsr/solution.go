package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
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
		scanner.Scan()
		x := scanner.Bytes()
		scanner.Scan()
		y := scanner.Bytes()
		z := readNNums(scanner, 3)
		a, b, k := z[0], z[1], z[2]
		fmt.Println(solve(x, y, a, b, k))
		t--
	}
}

func solve(s, t []byte, a, b, k int) int {
	if a == 0 {
		return 0
	}
	m := len(s)
	n := len(t)

	if b == 0 {
		res := abs(m-n) * a
		if res > k {
			return -1
		}
		return res
	}

	if k == 0 {
		eq := reflect.DeepEqual(s, t)
		if eq {
			return 0
		}
		return -1
	}
	matrixPrev := make([]int, n+1)
	for i := 0; i <= n; i++ {
		matrixPrev[i] = i * a
	}

	k1 := k + 2
	matrixCurr := make([]int, n+1)
	for i := 0; i < m; i++ {
		matrixCurr[i] = (i + 1) * a
		if i-k1 > 0 {
			// out of range
			matrixCurr[i-k1] = k1
		}

		for j := max(0, i-k1); j < min(n, i+k1); j++ {
			var cost int
			if s[i] != t[j] {
				cost = b
			}
			matrixCurr[j+1] = min3(matrixCurr[j]+a, matrixPrev[j+1]+a, matrixPrev[j]+cost)
		}

		for j := max(0, i-k1); j < min(n, i+k1); j++ {
			matrixPrev[j] = matrixCurr[j]
		}
		matrixPrev[min(n, i+k1)] = k1
	}

	res := matrixCurr[min(m-1+k1, n)]
	if res > k {
		return -1
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func min3(a, b, c int) int {
	if a <= b && a <= c {
		return a
	}
	if b <= a && b <= c {
		return b
	}
	return c
}
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
