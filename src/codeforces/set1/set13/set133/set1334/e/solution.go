package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n := readNum(reader)
	a := readNNums(reader, n)
	cost := readNNums(reader, n)
	m := readNum(reader)
	b := readNNums(reader, m)
	ok, res := solve(a, cost, b)

	if !ok {
		fmt.Println("NO")
	} else {
		fmt.Printf("YES\n%d\n", res)
	}
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

func solve(a []int, cost []int, b []int) (bool, int) {
	n := len(a)
	m := len(b)
	pos := make([][]int, n+1)
	negSum := make([]int, n+1)

	for i := 0; i < n; i++ {
		if pos[a[i]] == nil {
			pos[a[i]] = make([]int, 0, 1)
		}
		pos[a[i]] = append(pos[a[i]], i)
		negSum[i+1] = negSum[i]
		if cost[i] < 0 {
			negSum[i+1] += cost[i]
		}
	}

	tr := make(fenwick, n+1)
	fp := make([]int, n+1)
	for i := 0; i <= n; i++ {
		fp[i] = 1e18
	}
	var sum int
	for i, v := range a {
		if v > b[0] {
			// it has to be deleted
			if cost[i] > 0 {
				tr.update(i, cost[i])
			}
		} else if v == b[0] {
			// 前面的都删除掉
			fp[i] = sum
		}
		sum += cost[i]
	}

	pid := b[0] + 1

	for bi := 1; bi < m; bi++ {
		p := pos[b[bi-1]]
		var j int
		var mn int = 1e18
		pre := -1
		for _, i := range pos[b[bi]] {
			if pre >= 0 {
				mn += tr.query(pre, i-1) + negSum[i] - negSum[pre]
			}
			for j < len(p) && p[j] < i {
				// p[j] 在i的前面
				res := fp[p[j]] + tr.query(p[j]+1, i-1) + negSum[i] - negSum[p[j]+1]
				mn = min(mn, res)
				j++
			}
			fp[i] = mn
			pre = i
		}
		for pid <= b[bi] {
			for _, i := range pos[pid] {
				if cost[i] > 0 {
					tr.update(i, -cost[i])
				}
			}
			pid++
		}
	}
	var ans int = 1e18
	sum = 0
	for i := n - 1; i >= 0; i-- {
		v := a[i]
		if v == b[m-1] {
			ans = min(ans, fp[i]+sum)
		}
		if v > b[m-1] || cost[i] < 0 {
			sum += cost[i]
		}
	}

	if ans > 1e17 {
		return false, 0
	}

	return true, ans
}

type fenwick []int

func (f fenwick) update(i int, val int) {
	for i++; i < len(f); i += i & -i {
		f[i] += val
	}
}
func (f fenwick) pre(i int) (res int) {
	for i++; i > 0; i &= i - 1 {
		res += f[i]
	}
	return
}
func (f fenwick) query(l, r int) int {
	return f.pre(r) - f.pre(l-1)
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
