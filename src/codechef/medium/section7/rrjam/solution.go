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

func readThreeNums(scanner *bufio.Scanner) (a int, b int, c int) {
	res := readNNums(scanner, 3)
	a, b, c = res[0], res[1], res[2]
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
	scanner := bufio.NewScanner(os.Stdin)
	n := readNum(scanner)
	A := readNNums(scanner, n)
	B := readNNums(scanner, n)
	m := readNum(scanner)
	L, R, X, Y := make([]int, m), make([]int, m), make([]int, m), make([]int, m)
	for i := 0; i < m; i++ {
		line := readNNums(scanner, 4)
		L[i], R[i], X[i], Y[i] = line[0], line[1], line[2], line[3]
	}
	res := solve(n, A, B, m, L, R, X, Y)

	for i := 0; i < n; i++ {
		if i < n-1 {
			fmt.Printf("%d ", res[i])
		} else {
			fmt.Printf("%d\n", res[i])
		}
	}
}

func solve(n int, A []int, B []int, m int, L []int, R []int, X []int, Y []int) []int {
	begin := make([][]int, n+1)
	end := make([][]int, n+1)

	for i := 1; i <= n; i++ {
		begin[i] = make([]int, 0, 10)
	}

	for i := 0; i < m; i++ {
		begin[L[i]] = append(begin[L[i]], i+1)
		end[R[i]] = append(end[R[i]], i+1)
	}

	sumx := make([]int64, 4*m)
	sumy := make([]int64, 4*m)

	var update func(i int, left, right int, u int, x, y int64)

	update = func(i int, left, right int, u int, x, y int64) {
		if u < left || right < u {
			return
		}

		if left == right {
			sumx[i] = x
			sumy[i] = y
			return
		}
		mid := (left + right) / 2
		update(2*i, left, mid, u, x, y)
		update(2*i+1, mid+1, right, u, x, y)

		sumx[i] = sumx[2*i] + sumx[2*i+1]
		sumy[i] = sumy[2*i] + sumy[2*i+1]
	}

	var get func(i int, left int, right int, u int, v int64) int
	get = func(i int, left int, right int, u int, v int64) int {
		if left == right {
			return left
		}
		mid := (left + right) / 2
		cur := sumx[2*i] + sumy[2*i]*int64(u)
		if cur >= v {
			return get(2*i, left, mid, u, v)
		}
		return get(2*i+1, mid+1, right, u, v-cur)
	}

	ans := make([]int, n)

	// V[i] (finaly value of V) after operation j is X[j] + (j - L[j]) * Y[j] = X[j] - L[j] * Y[j] + Y[j] * j
	// here X[j] - L[j] * Y[j] is determined, and Y[j] changes along with j
	for pos := 1; pos <= n; pos++ {
		for _, idx := range begin[pos] {
			x, y := X[idx-1], Y[idx-1]
			// l == pos
			update(1, 1, m, idx, int64(x)-int64(pos)*int64(y), int64(y))
		}

		if B[pos-1] <= A[pos-1] {
			ans[pos-1] = 0
		} else if sumx[1]+sumy[1]*int64(pos) >= int64(B[pos-1]-A[pos-1]) {
			ans[pos-1] = get(1, 1, m, pos, int64(B[pos-1]-A[pos-1]))
		} else {
			ans[pos-1] = -1
		}

		for _, idx := range end[pos] {
			update(1, 1, m, idx, 0, 0)
		}
	}

	return ans
}
