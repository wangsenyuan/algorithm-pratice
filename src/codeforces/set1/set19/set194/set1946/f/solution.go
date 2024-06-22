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
		n, m := readTwoNums(reader)
		a := readNNums(reader, n)
		queries := make([][]int, m)
		for i := 0; i < m; i++ {
			queries[i] = readNNums(reader, 2)
		}
		res := solve(a, queries)
		for i := 0; i < m; i++ {
			buf.WriteString(fmt.Sprintf("%d ", res[i]))
		}
		buf.WriteByte('\n')
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

func solve(a []int, queries [][]int) []int {
	reverse(a)

	n := len(a)
	pos := make([]int, n+1)
	for i, num := range a {
		pos[num] = i
	}

	at := make([][]pair, n)

	for i, cur := range queries {
		l, r := cur[0]-1, cur[1]-1
		l, r = r, l
		l = n - 1 - l
		r = n - 1 - r
		at[r] = append(at[r], pair{l, i})
	}
	dp := make([]int, n+1)
	tr := make(fenwick, n+10)

	ans := make([]int, len(queries))

	for r := 0; r < n; r++ {
		x := a[r]
		// dp[u] 表示从x开始的，且在r的前面的，能够形成z % y = 0, y % x = 0 的所有序列的计数
		dp[x] = 1
		for y := x; y <= n; y += x {
			if pos[y] > pos[x] {
				// y 在x的后面
				continue
			}
			for z := 2 * y; z <= n; z += y {
				if pos[z] > pos[y] {
					continue
				}
				dp[z] += dp[y]
			}
		}
		for y := x; y <= n; y += x {
			tr.add(pos[y], dp[y])
			dp[y] = 0
		}

		for _, cur := range at[r] {
			l, i := cur.first, cur.second
			ans[i] = tr.query(l, r)
		}
	}

	return ans
}

type pair struct {
	first  int
	second int
}

func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

type fenwick []int

func (tr fenwick) add(p int, v int) {
	p++
	for p < len(tr) {
		tr[p] += v
		p += p & -p
	}
}

func (tr fenwick) pre(p int) int {
	var res int
	p++
	for p > 0 {
		res += tr[p]
		p -= p & -p
	}
	return res
}

func (tr fenwick) query(l int, r int) int {
	res := tr.pre(r)
	if l > 0 {
		res -= tr.pre(l - 1)
	}
	return res
}
