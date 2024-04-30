package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)
	p := readNNums(reader, n)
	res := solve(p)

	fmt.Println(res)
}

func readString(r *bufio.Reader) string {
	s, _ := r.ReadBytes('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return string(s[:i])
		}
	}
	return string(s)
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

func solve(p []int) int {
	n := len(p)
	id := make([]int, n)
	for i := 0; i < n; i++ {
		p[i]--
		id[i] = i
	}
	dp := make([]bool, n)
	dp[0] = true
	x := p[0]
	for i := 1; i < n; i++ {
		if x < p[i] {
			dp[i] = true
			x = p[i]
		}
	}
	sort.Slice(id, func(i, j int) bool {
		return p[id[i]] > p[id[j]]
	})

	l := n
	tr := make(fenwick, n+10)

	cnt := make([]int, n)

	for _, i := range id {
		if tr.pre(i) == 1 {
			// 在它前面只有一个更大的值，删掉它后，贡献 + 1
			cnt[l]++
		}
		tr.update(i, 1)
		l = min(l, i)
	}

	ans := []int{p[0], cnt[0] - 1}

	for i := 1; i < n; i++ {
		tmp := cnt[i]
		if dp[i] {
			tmp--
		}
		if tmp > ans[1] || tmp == ans[1] && p[i] < ans[0] {
			ans[0] = p[i]
			ans[1] = tmp
		}
	}

	return ans[0] + 1
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
