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

	var buf bytes.Buffer

	for tc > 0 {
		tc--
		n, k, l := readThreeNums(reader)
		d := readNNums(reader, n)
		res := solve(n, k, l, d)

		if res {
			buf.WriteString("YES\n")
		} else {
			buf.WriteString("NO\n")
		}
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

func solve(n int, k int, l int, d []int) bool {
	// t % 2k = ...
	dp := make([]bool, 2*k)
	fp := make([]bool, 2*k)
	p := make([]int, 2*k)
	for i := 0; i < k; i++ {
		p[i] = i
	}
	p[k] = k
	for i := k + 1; i < 2*k; i++ {
		p[i] = p[i-1] - 1
	}

	// 在岸上的时候都是安全的
	for t := 0; t < 2*k; t++ {
		dp[t] = true
	}

	for i := 0; i < n; i++ {
		for t := 0; t < 2*k; t++ {
			if dp[t] {
				if t+1 < 2*k && d[i]+p[t+1] <= l {
					fp[t+1] = true
				}
				if t+1 == 2*k && d[i]+p[0] <= l {
					fp[0] = true
				}
			}
		}
		for t := 1; t < 2*k; t++ {
			if fp[t-1] && d[i]+p[t] <= l {
				fp[t] = true
			}
		}
		copy(dp, fp)
		clear(fp)
	}

	for j := 0; j < 2*k; j++ {
		if dp[j] {
			return true
		}
	}
	return false
}
