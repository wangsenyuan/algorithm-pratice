package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var buf bytes.Buffer

	tc := readNum(reader)

	for tc > 0 {
		tc--
		n := readNum(reader)
		A := readNNums(reader, n)
		res := solve(A)
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
		if s[i] == '\n' || s[i] == '\r' {
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

func solve(P []int) int64 {
	// all distinct
	// a, b, c
	// a < b < c already werid
	// a > b > c, werid too, after all negative
	// a > b < c and a > c  3, 1, 2 werid
	//            or a < c   2, 1, 3 werid
	// a < b > c     a > c,  2, 3, 1 not werid
	//               a < c,  1, 3, 2 not werid
	// 如果给定a, c 如果 a < b, 且存在一个c b > c, 则 a....c 是 good (not werid)
	// a b c 且 a < b > c, 则包含它们的所有subarray 都是 good
	n := len(P)
	N := int64(n)
	res := N * (N + 1) / 2
	if n <= 2 {
		return res
	}

	last := -1

	for i := 1; i+1 < n; i++ {
		if P[i-1] < P[i] && P[i] > P[i+1] {
			tmp := int64(i-1-last) * int64(n-(i+1))
			res -= tmp
			// 只要不包括i-1, 就能保证不重复
			last = i - 1
		}
	}

	return res
}
