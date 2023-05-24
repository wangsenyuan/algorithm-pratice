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
		s, _ := reader.ReadBytes('\n')
		var n int
		pos := readInt(s, 0, &n)
		var k uint64
		readUint64(s, pos+1, &k)

		res := solve(n, int64(k))

		if len(res) == 0 {
			buf.WriteString("-1\n")
		} else {
			for i := 0; i < n; i++ {
				buf.WriteString(fmt.Sprintf("%d ", res[i]))
			}
			buf.WriteByte('\n')
		}
	}
	fmt.Print(buf.String())
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

const N = 100010

var F [N]int64

func init() {
	F[0] = 1
	F[1] = 1
	var sum int64 = 2

	for i := 2; i < N; i++ {
		F[i] = sum
		sum += F[i]
	}
}

func solve(n int, k int64) []int {
	if n < 62 && k > 1<<uint64(n-1) {
		return nil
	}
	k--
	P := make([]int, n)

	for b := 0; b <= min(60, n-1); b++ {
		if (k>>b)&1 == 1 {
			P[n-2-b] = 1
		}
	}
	ans := make([]int, 0, n)
	cur, sz := 1, 1
	for i := 0; i < n; i++ {
		if P[i] == 0 {
			for j := cur + sz - 1; j >= cur; j-- {
				ans = append(ans, j)
			}
			cur += sz
			sz = 1
		} else {
			sz++
		}
	}

	return ans
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}
