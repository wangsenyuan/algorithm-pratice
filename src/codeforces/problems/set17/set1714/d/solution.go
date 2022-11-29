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

		T := readString(reader)

		n := readNum(reader)
		S := make([]string, n)

		for i := 0; i < n; i++ {
			S[i] = readString(reader)
		}

		res := solve(T, S)

		if len(res) == 0 {
			buf.WriteString("-1\n")
		} else {
			buf.WriteString(fmt.Sprintf("%d\n", len(res)))
			for _, cur := range res {
				buf.WriteString(fmt.Sprintf("%d %d\n", cur[0], cur[1]))
			}
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

const INF = 1 << 30

func solve(T string, S []string) [][]int {
	n := len(T)

	dp := make([]Pair, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = Pair{-1, -1}
	}
	N := n + 1

	arr := make([]int, 2*N)

	for i := 0; i < len(arr); i++ {
		arr[i] = INF
	}

	set := func(p int, v int) {
		p += N
		arr[p] = v
		for p > 1 {
			arr[p>>1] = min(arr[p], arr[p^1])
			p >>= 1
		}
	}

	get := func(l int, r int) int {
		l += N
		r += N
		res := INF
		for l < r {
			if l&1 == 1 {
				res = min(res, arr[l])
				l++
			}
			if r&1 == 1 {
				r--
				res = min(res, arr[r])
			}
			l >>= 1
			r >>= 1
		}
		return res
	}

	set(0, 0)

	for i := 1; i <= n; i++ {
		for j, s := range S {
			if i-len(s) >= 0 && T[i-len(s):i] == s {
				// j is the string index
				k := i - len(s) + 1
				x := get(k-1, i)
				if x < INF {
					if dp[i].first < 0 || dp[i].first > x+1 {
						dp[i] = Pair{x + 1, j}
						set(i, x+1)
					}
				}
			}
		}
	}

	if dp[n].first < 0 {
		return nil
	}

	var res [][]int
	cur := dp[n].first

	for i := n; i > 0; i-- {
		if dp[i].first == cur {
			j := dp[i].second
			res = append(res, []int{j + 1, 1 + i - len(S[j])})
			cur--
		}
	}

	reverse(res)

	return res
}

func reverse(arr [][]int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

type Pair struct {
	first  int
	second int
}
