package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)
	// tc := 1

	var buf bytes.Buffer
	for tc > 0 {
		tc--
		n, m := readTwoNums(reader)
		T := readNNums(reader, m)
		C := readNNums(reader, m)
		res := solve(n, T, C)
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

func readNInt64s(reader *bufio.Reader, n int) []int64 {
	res := make([]int64, n)
	s, _ := reader.ReadBytes('\n')

	var pos int

	for i := 0; i < n; i++ {
		pos = readInt64(s, pos, &res[i]) + 1
	}

	return res
}

func readInt64(bytes []byte, from int, val *int64) int {
	i := from
	var sign int64 = 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	var tmp int64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int64(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
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

	for len(bs) == 0 || bs[0] == '\n' || bs[0] == '\r' {
		bs, _ = reader.ReadBytes('\n')
	}

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

const INF = 1e9

func solve(n int, T []int, C []int) int {
	dp := make([]int, 5)
	for i := 0; i < 5; i++ {
		dp[i] = -INF
	}
	dp[2] = 0

	fp := make([]int, 5)

	var ptr int

	for i := 1; i <= n; i++ {
		for j := 0; j < 5; j++ {
			fp[j] = -INF
		}
		for dif := 0; dif < 5; dif++ {
			if dif > 0 {
				fp[dif] = max(fp[dif], dp[dif-1])
			}
			if dif < 4 {
				fp[dif] = max(fp[dif], dp[dif+1])
			}
		}

		if ptr < len(T) && T[ptr] == i {
			if C[ptr] == 1 {
				fp[4]++
				fp[3]++
				fp[2] = max_of_array(fp[:3])
				fp[1] = -INF
				fp[0] = -INF
			} else {
				fp[0]++
				fp[1]++
				fp[2] = max_of_array(fp[2:])
				fp[3] = -INF
				fp[4] = -INF
			}

			ptr++
		}
		copy(dp, fp)
	}

	return max_of_array(dp)
}

func max_of_array(arr []int) int {
	res := arr[0]
	for i := 1; i < len(arr); i++ {
		res = max(res, arr[i])
	}
	return res
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
