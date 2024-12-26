package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	fmt.Println(res)
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

func process(reader *bufio.Reader) int {
	n, m := readTwoNums(reader)
	a := readNNums(reader, n)
	return solve(m, a)
}

func solve(m int, a []int) int {

	pre := make([]int, m+1)
	for i := 0; i <= m; i++ {
		pre[i] = -1
	}
	var res int
	n := len(a)

	tr := make(fenwick, n+10)

	for i, x := range a {
		if pre[x] < 0 {
			res += m - 1
		} else {
			// 上次出现到当前位置，出现的distinct 数字的个数
			res += tr.get(i) - tr.get(pre[x])
			tr.update(pre[x], -1)
		}
		pre[x] = i
		tr.update(i, 1)
	}

	return res
}

type fenwick []int

func (tr fenwick) update(p int, v int) {
	p++
	for p < len(tr) {
		tr[p] += v
		p += p & -p
	}
}

func (tr fenwick) get(p int) int {
	p++
	var res int
	for p > 0 {
		res += tr[p]
		p -= p & -p
	}
	return res
}
