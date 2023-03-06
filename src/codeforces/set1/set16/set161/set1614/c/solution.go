package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	//var buf bytes.Buffer
	tc := readNum(reader)

	for tc > 0 {
		tc--
		n, m := readTwoNums(reader)
		segs := make([][]int, m)
		for i := 0; i < m; i++ {
			segs[i] = readNNums(reader, 3)
		}
		res := solve(n, segs)

		fmt.Println(res)
		//buf.WriteString(fmt.Sprintf("%d\n", res))
	}
	//fmt.Print(buf.String())
}

func normalize(s string) string {

	for i := len(s); i > 0; i-- {
		if s[i-1] >= 'a' && s[i-1] <= 'z' {
			return s[:i]
		}
	}
	return ""
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

const MOD = 1e9 + 7

func modMul(a, b int) int {
	return int(int64(a) * int64(b) % MOD)
}

func pow(a, b int) int {
	r := 1
	for b > 0 {
		if b&1 == 1 {
			r = modMul(r, a)
		}
		a = modMul(a, a)
		b >>= 1
	}
	return r
}

func solve(n int, segs [][]int) int {
	// digit by digit
	// each sub-seq of A
	// 考虑包含A[0]的子序列
	// 如果包含0的seg,or[d] = 0, 那么 A[0][d] = 0
	//  否则 A[0][d] = 1
	var x int
	for _, seg := range segs {
		x |= seg[2]
	}

	res := modMul(x, pow(2, n-1))

	return int(res)
}
