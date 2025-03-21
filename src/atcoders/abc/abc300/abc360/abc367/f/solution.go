package main

import (
	"bufio"
	"bytes"
	"os"

	"math/rand"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	ans := process(reader)
	var buf bytes.Buffer
	for _, x := range ans {
		if x {
			buf.WriteString("Yes\n")
		} else {
			buf.WriteString("No\n")
		}
	}

	buf.WriteTo(os.Stdout)
}

func readInt(bytes []byte, from int, val *int) int {
	i := from
	sign := 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	tmp := 0
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
}

func readNum(reader *bufio.Reader) (a int) {
	bs, _ := reader.ReadBytes('\n')
	readInt(bs, 0, &a)
	return
}

func readTwoNums(reader *bufio.Reader) (a int, b int) {
	res := readNNums(reader, 2)
	a, b = res[0], res[1]
	return
}

func readThreeNums(reader *bufio.Reader) (a int, b int, c int) {
	res := readNNums(reader, 3)
	a, b, c = res[0], res[1], res[2]
	return
}

func readNNums(reader *bufio.Reader, n int) []int {
	res := make([]int, n)
	x := 0
	bs, _ := reader.ReadBytes('\n')
	for i := 0; i < n; i++ {
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') && bs[x] != '-' {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
	return res
}

func process(reader *bufio.Reader) []bool {
	n, m := readTwoNums(reader)
	A := readNNums(reader, n)
	B := readNNums(reader, n)
	Q := make([][]int, m)
	for i := 0; i < m; i++ {
		Q[i] = readNNums(reader, 4)
	}
	return solve(A, B, Q)
}

const mod = 1 << 61

func solve(A []int, B []int, Q [][]int) []bool {
	n := len(A)
	as := make([]int, n+1)
	bs := make([]int, n+1)

	r := rand.New(rand.NewSource(99))

	hash := make([]int, n)
	for i := 0; i < n; i++ {
		hash[i] = r.Intn(mod)
	}

	for i := 0; i < n; i++ {
		as[i+1] = as[i] + hash[A[i]-1]
		bs[i+1] = bs[i] + hash[B[i]-1]
	}

	ans := make([]bool, len(Q))

	calc := func(num int) int {
		if num < 0 {
			return num + mod
		}
		return num
	}

	for i, cur := range Q {
		l1, r1, l2, r2 := cur[0], cur[1], cur[2], cur[3]
		if r1-l1 == r2-l2 {
			s1 := as[r1] - as[l1-1]
			s2 := bs[r2] - bs[l2-1]
			ans[i] = calc(s1) == calc(s2)
		}
	}

	return ans
}
