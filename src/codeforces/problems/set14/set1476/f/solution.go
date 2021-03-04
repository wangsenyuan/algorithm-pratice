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

		n := readNum(reader)
		P := readNNums(reader, n)
		res := solve(n, P)

		if len(res) == 0 {
			buf.WriteString("NO\n")
			continue
		}
		buf.WriteString("YES\n")
		buf.WriteString(res)
		buf.WriteByte('\n')
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

const INF = 1 << 30

func solve(n int, P []int) string {
	L := make([]int, n)
	R := make([]int, n)

	for i := 0; i < n; i++ {
		L[i] = max(1, i-P[i]+1)
		R[i] = min(n, i+P[i]+1)
	}

	sx := make([]int, 2*n)

	for i := n; i < 2*n; i++ {
		sx[i] = R[i-n]
	}

	for i := n - 1; i > 0; i-- {
		sx[i] = max(sx[i<<1], sx[(i<<1)|1])
	}

	getMax := func(l, r int) int {
		var res int
		l += n
		r += n
		for l < r {
			if l&1 == 1 {
				res = max(res, sx[l])
				l++
			}
			if r&1 == 1 {
				r--
				res = max(res, sx[r])
			}
			l >>= 1
			r >>= 1
		}
		return res
	}

	nn := n + 1
	sn := make([]int, 2*nn)
	for i := 0; i < len(sn); i++ {
		sn[i] = INF
	}

	update := func(p int, v int) {
		p += nn
		sn[p] = min(sn[p], v)
		for p > 1 {
			sn[p>>1] = min(sn[p], sn[p^1])
			p >>= 1
		}
	}

	get := func(l, r int) int {
		res := INF
		l += nn
		r += nn
		for l < r {
			if l&1 == 1 {
				res = min(res, sn[l])
				l++
			}
			if r&1 == 1 {
				r--
				res = min(res, sn[r])
			}
			l >>= 1
			r >>= 1
		}
		return res
	}

	dp := make([]int, n+1)
	par := make([]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = -INF
	}
	dp[0] = 0
	par[0] = -1

	update(0, 0)

	for i := 1; i <= n; i++ {
		j := i - 1
		k := L[j] - 1
		// find the min m that dp[m] >= k
		m := get(k, n+1)
		if m < INF {
			nval := max(getMax(m, i-1), i-1)
			if nval > dp[i] {
				dp[i] = nval
				par[i] = m
			}
		}
		if dp[j] >= i && max(dp[j], R[j]) > dp[i] {
			dp[i] = max(dp[j], R[j])
			par[i] = -1
		}
		if dp[j] > dp[i] {
			dp[i] = dp[j]
			par[i] = -1
		}
		update(dp[i], i)
	}

	if dp[n] != n {
		return ""
	}
	res := make([]byte, n)
	cur := n
	for cur > 0 {
		if par[cur] < 0 {
			res[cur-1] = 'R'
			cur--
			continue
		}
		diff := cur - par[cur]
		res[cur-1] = 'L'
		cur--
		for diff > 1 {
			res[cur-1] = 'R'
			cur--
			diff--
		}
	}
	return string(res)
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
