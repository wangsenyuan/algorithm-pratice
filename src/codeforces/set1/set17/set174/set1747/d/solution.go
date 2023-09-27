package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, m := readTwoNums(reader)
	v := readNNums(reader, n)
	q := make([][]int, m)
	for i := 0; i < m; i++ {
		q[i] = readNNums(reader, 2)
	}
	res := solve(v, q)

	var buf bytes.Buffer

	for _, x := range res {
		buf.WriteString(fmt.Sprintf("%d\n", x))
	}

	fmt.Print(buf.String())
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

func solve(v []int, queries [][]int) []int {
	n := len(v)
	x := make([]int, n+1)
	sum := make([]int64, n+1)
	dp := make([]int, n+1)
	even := make(map[int]int)
	odd := make(map[int]int)
	for i := 1; i <= n; i++ {
		x[i] = x[i-1] ^ v[i-1]
		sum[i] = sum[i-1] + int64(v[i-1])

		if i&1 == 1 {
			if j, ok := even[x[i]]; ok {
				dp[i] = j
			}
			odd[x[i]] = i
		} else {
			if j, ok := odd[x[i]]; ok {
				dp[i] = j
			}
			even[x[i]] = i
		}
	}

	ans := make([]int, len(queries))
	for i, cur := range queries {
		l, r := cur[0], cur[1]
		ans[i] = -1
		if x[r]^x[l-1] != 0 {
			continue
		}
		if sum[r]-sum[l-1] == 0 {
			// all 0
			ans[i] = 0
			continue
		}
		if (r-l+1)%2 == 1 {
			ans[i] = 1
			continue
		}
		if v[l-1] == 0 || v[r-1] == 0 {
			ans[i] = 1
			continue
		}
		// 这里还有长度的考虑
		if dp[r] >= l {
			ans[i] = 2
		}
	}

	return ans
}
