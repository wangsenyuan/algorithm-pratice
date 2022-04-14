package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)
	var buf bytes.Buffer
	for tc > 0 {
		tc--
		n := readNum(reader)
		A := readNNums(reader, n)
		B := readNNums(reader, n)
		res := solve(A, B)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}
	fmt.Print(buf.String())
}

func readUint64(bytes []byte, from int, val *uint64) int {
	i := from

	var tmp uint64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + uint64(bytes[i]-'0')
		i++
	}
	*val = tmp

	return i
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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' {
			return s[:i]
		}
	}
	return s
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

const H = 20

func solve(A []int, B []int) int64 {
	// a ^ b, a & b, 如果a和b的最高位一致， a ^ b < a & b, else a ^ b > a & b
	var res int64
	n := len(A)

	sort.Ints(A)
	sort.Ints(B)

	reverse(A)
	reverse(B)

	dp := count(A)
	fp := count(B)

	get := func(xp [][]int, a, b int, d int) int {
		if b < a {
			return 0
		}
		res := xp[b][d]
		if a > 0 {
			res -= xp[a-1][d]
		}
		return res
	}

	var p, q int
	for d := H - 1; d >= 0; d-- {
		i, j := p, q
		for p < n && (A[p]>>uint(d))&1 == 1 {
			p++
		}
		for q < n && (B[q]>>uint(d))&1 == 1 {
			q++
		}
		// A[i...p] B[j...q] all have highest digit d set
		// it contributes A[i..p] & B[j..q] to result and A[i..p] ^ B[q..], A[p..] ^ B[j...q]
		for u := d; u >= 0; u-- {
			x := get(dp, i, p-1, u)
			y := get(fp, j, q-1, u)
			res += int64(x) * int64(y) * (1 << uint64(u))
			a := get(dp, p, n-1, u)
			b := get(fp, q, n-1, u)
			res += int64(x) * int64(n-q-b) * (1 << uint64(u))
			res += int64(p-i-x) * int64(b) * (1 << uint64(u))
			res += int64(y) * int64(n-p-a) * (1 << uint64(u))
			res += int64(q-j-y) * int64(a) * (1 << uint64(u))
		}
	}

	return res
}

func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func count(arr []int) [][]int {
	n := len(arr)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, H)
		if i > 0 {
			copy(dp[i], dp[i-1])
		}
		for j := 0; j < H; j++ {
			dp[i][j] += (arr[i] >> uint(j)) & 1
		}
	}
	return dp
}
