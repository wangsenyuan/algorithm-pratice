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

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	tc := readNum(scanner)

	for tc > 0 {
		tc--
		line := readNNums(scanner, 3)
		N, P, Q := line[0], line[1], line[2]
		A := readNNums(scanner, N)
		B := readNNums(scanner, Q/64+2)
		res := solve(N, P, Q, A, B)
		fmt.Println(res)
	}
}

const MAX_L = 21

func solve(n, p, Q int, A []int, B []int) int {
	tl := make([][]int64, MAX_L)
	tr := make([][]int64, MAX_L)

	for i := 0; i < MAX_L; i++ {
		tl[i] = make([]int64, n)
		tr[i] = make([]int64, n)
	}

	P := int64(p)
	for j := 0; ; j++ {
		sz := 1 << uint(j)
		if sz > n {
			break
		}
		ones := sz - 1

		tmp := int64(1)
		for i := 0; i < n; i++ {
			tmp *= int64(A[i])
			tmp %= P
			tl[j][i] = tmp
			if i&ones == ones {
				tmp = 1
			}
		}
		tmp = 1
		for i := n - 1; i >= 0; i-- {
			tmp *= int64(A[i])
			tmp %= P
			tr[j][i] = tmp
			if i&ones == 0 {
				tmp = 1
			}
		}
	}
	N := int64(n)
	var prev int64
	cnst := 1<<6 - 1
	var l, r int64
	for i := 0; i < Q; i++ {
		if i&cnst == 0 {
			l = int64(B[i>>6]) + prev
			r = int64(B[i>>6+1]) + prev
		} else {
			l += prev
			r += prev
		}
		l %= N
		r %= N
		if l > r {
			l, r = r, l
		}
		var ans int64
		if l == r {
			ans = int64(A[l])
		} else {
			num := 31 - clz(int(l^r))
			ans = tl[num][r] * tr[num][l]
		}
		ans = (ans + 1) % P
		prev = ans
	}
	return int(prev)
}

func clz(num int) int {
	var ans int

	for i := 31; i >= 0; i-- {
		if num&(1<<uint(i)) > 0 {
			break
		}
		ans++
	}

	return ans
}
