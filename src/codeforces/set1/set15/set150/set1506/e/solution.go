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
		n := readNum(reader)
		q := readNNums(reader, n)
		res := solve(q)
		for _, r := range res {
			for j := 0; j < n; j++ {
				buf.WriteString(fmt.Sprintf("%d ", r[j]))
			}
			buf.WriteByte('\n')
		}
	}
	fmt.Print(buf.String())
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')

	for i := 0; i < len(s); i++ {
		if s[i] == '\r' || s[i] == '\n' {
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

func solve(q []int) [][]int {
	res := [][]int{
		getMin(q),
		getMax(q),
	}
	return res
}

func getMax(q []int) []int {
	n := len(q)
	pos := make([]int, n+1)

	tr := make(SegTree, 2*(n+1))
	for i := 0; i <= n; i++ {
		tr.update(i, i)
	}
	var arr []int

	for i := 0; i < n; i++ {
		if pos[q[i]] == 0 {
			pos[q[i]] = i + 1
			arr = append(arr, q[i])
			tr.update(q[i], 0)
		}
	}

	res := make([]int, n)

	for i, j := 0, 0; i < len(arr); i++ {
		res[j] = arr[i]
		j++
		next := n
		if i+1 < len(arr) {
			next = pos[arr[i+1]] - 1
		}
		for j < next {
			v := tr.get(1, arr[i])
			res[j] = v
			tr.update(v, 0)
			j++
		}
	}

	return res
}

func getMin(q []int) []int {
	n := len(q)
	pos := make([]int, n+1)

	for i := 0; i < n; i++ {
		if pos[q[i]] == 0 {
			pos[q[i]] = i + 1
		}
	}

	var free []int
	var arr []int
	for i := 1; i <= n; i++ {
		if pos[i] == 0 {
			free = append(free, i)
		} else {
			arr = append(arr, i)
		}
	}

	res := make([]int, n+1)

	for i := 0; i < len(arr); i++ {
		next := n + 1
		if i+1 < len(arr) {
			next = pos[arr[i+1]]
		}
		res[pos[arr[i]]] = arr[i]
		for j := pos[arr[i]] + 1; j < next; j++ {
			res[j] = free[0]
			free = free[1:]
		}
	}

	return res[1:]
}

type SegTree []int

func (tr SegTree) update(i int, v int) {
	i += len(tr) / 2
	tr[i] = v
	for i > 1 {
		tr[i>>1] = max(tr[i], tr[i^1])
		i >>= 1
	}
}

func (tr SegTree) get(l int, r int) int {
	l += len(tr) / 2
	r += len(tr) / 2
	var res int
	for l < r {
		if l&1 == 1 {
			res = max(res, tr[l])
			l++
		}
		if r&1 == 1 {
			r--
			res = max(res, tr[r])
		}
		l >>= 1
		r >>= 1
	}
	return res
}
