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
	b := readNNums(reader, n)
	res := solve(a, b)
	fmt.Println(res)
}

func readString(r *bufio.Reader) string {
	s, _ := r.ReadString('\n')
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

const H = 18

func solve(a []int, b []int) int {
	f := func(a, p, q []int) (res int) {
		if p == nil || q == nil {
			return
		}
		for j := 0; j < 18; j++ {
			p1, q1 := 0, 0
			for _, i := range p {
				p1 += a[i] >> j & 1
			}
			for _, i := range q {
				q1 += a[i] >> j & 1
			}
			res += (p1*(len(q)-q1) + (len(p)-p1)*q1) << j
		}
		return
	}

	var ans int

	var dfs func([]int, int)
	dfs = func(idx []int, j int) {
		if len(idx) <= 1 {
			return
		}
		if j < 0 {
			ans += f(a, idx, idx) / 2
			return
		}
		p := [4][]int{}
		for _, i := range idx {
			m := a[i]>>j&1<<1 | b[i]>>j&1
			p[m] = append(p[m], i)
		}
		ans += f(a, p[0], p[1])
		ans += f(a, p[2], p[3])
		ans += f(b, p[0], p[2])
		ans += f(b, p[1], p[3])
		dfs(append(p[0], p[3]...), j-1)
		dfs(append(p[1], p[2]...), j-1)
	}
	n := len(a)

	idx := make([]int, n)

	for i := range idx {
		idx[i] = i
	}

	dfs(idx, H-1)

	return ans
}
