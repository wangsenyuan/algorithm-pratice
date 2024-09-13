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
	a := readNNums(reader, m)
	res := solve(n, a)
	var buf bytes.Buffer

	for _, row := range res {
		buf.WriteString(fmt.Sprintf("%d %d\n", row[0], row[1]))
	}

	fmt.Print(buf.String())
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

func solve(n int, a []int) [][]int {
	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = []int{i, i}
	}
	m := len(a)
	pos := make([]int, n)
	set := make(fenwick, n+m+10)
	for i := n - 1; i >= 0; i-- {
		pos[i] = n - 1 - i
		set.update(pos[i], 1)
	}

	for j, i := range a {
		// 看i右边有多少个数字
		i--
		res[i][0] = 0
		res[i][1] = max(res[i][1], set.query(pos[i]+1, n+m))
		set.update(pos[i], -1)
		pos[i] = j + n
		set.update(pos[i], 1)
	}

	for i := 0; i < n; i++ {
		res[i][1] = max(res[i][1], set.query(pos[i]+1, n+m))
		res[i][0]++
		res[i][1]++
	}

	return res
}

func solve1(n int, a []int) [][]int {
	res := make([][]int, n)
	pos := make([]int, n)
	for i := 0; i < n; i++ {
		res[i] = []int{i, i}
		pos[i] = -1
	}
	m := len(a)

	set := make(fenwick, n+10)
	occ := make(fenwick, m+10)

	for id, i := range a {
		i--
		if pos[i] < 0 {
			// 这个是第一次出现, 要看它后边后多少个数，被移动到了前面
			res[i][1] = max(res[i][1], i+set.query(i+1, n))
			set.update(i, 1)
		} else {
			// 这个是第二次出现了, 上次到现在出现了多少个数
			cnt := occ.query(pos[i]+1, m)
			res[i][1] = max(res[i][1], cnt)
			// 从上次的位置移除掉
			occ.update(pos[i], -1)
		}

		pos[i] = id
		occ.update(id, 1)
		res[i][0] = 0
	}

	for i := 0; i < n; i++ {
		if pos[i] < 0 {
			// 这个数字没有出现过
			res[i][1] += set.query(i+1, n)
		} else {
			// 还需要再处理一遍
			cnt := occ.query(pos[i]+1, m)
			res[i][1] = max(res[i][1], cnt)
		}
		res[i][0]++
		res[i][1]++
	}

	return res
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
