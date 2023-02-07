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
		n, k := readTwoNums(reader)
		A := readNNums(reader, n)
		res := solve(A, k)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}

	fmt.Print(buf.String())
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

func solve(A []int, k int) int64 {
	// 考虑A全部相等的情况，那么从0开始往后一直获取
	// 考虑经过时间i后的收获， 1 + 2 + 3 + 4 + + i - 1
	// 如果到了i处往回走，肯定是不明智的。
	n := len(A)
	var res int64
	if k >= n {
		for _, num := range A {
			res += int64(num)
		}
		// 一共是n个点，如果不取的话，最后可以获取到 n * k个，
		// 假设让它们生长，在最后的(k - n)秒开始收获
		res += int64(n) * int64(k-1)
		res -= int64(n-1) * int64(n) / 2
		return res
	}
	// k < n
	// 找一段k个和最大的部分
	var sum int64
	for i := 0; i < k-1; i++ {
		sum += int64(A[i])
	}

	for i := k - 1; i < n; i++ {
		sum += int64(A[i])
		res = max(res, sum)
		sum -= int64(A[i-(k-1)])
	}

	res += int64(k) * int64(k-1) / 2

	return res
}

func max(a, b int64) int64 {
	if a >= b {
		return a
	}
	return b
}
