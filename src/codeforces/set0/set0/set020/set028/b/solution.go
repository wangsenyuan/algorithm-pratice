package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)
	b := readNNums(reader, n)
	d := readNNums(reader, n)

	res := solve(d, b)

	if res {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
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

func solve(d []int, b []int) bool {
	n := len(d)
	// 也就是说b[i]和i要能连接起来
	set := NewDSU(n + 1)
	for i := 1; i <= n; i++ {
		v := d[i-1]
		if v != 0 {
			if i-v > 0 {
				set.Union(i, i-v)
			}
			if i+v <= n {
				set.Union(i, i+v)
			}
		}
	}

	for i := 1; i <= n; i++ {
		x := b[i-1]
		if set.Find(i) != set.Find(x) {
			return false
		}
	}
	return true
}

type DSU struct {
	arr []int
	cnt []int
}

func NewDSU(n int) *DSU {
	d := new(DSU)
	d.arr = make([]int, n)
	d.cnt = make([]int, n)
	for i := 0; i < n; i++ {
		d.arr[i] = i
		d.cnt[i]++
	}
	return d
}

func (d *DSU) Find(u int) int {
	if d.arr[u] != u {
		d.arr[u] = d.Find(d.arr[u])
	}
	return d.arr[u]
}

func (d *DSU) Union(a, b int) bool {
	a = d.Find(a)
	b = d.Find(b)
	if a == b {
		return false
	}
	if d.cnt[a] < d.cnt[b] {
		a, b = b, a
	}
	d.cnt[a] += d.cnt[b]
	d.arr[b] = a
	return true
}

func (d *DSU) Reset() {
	for i := 0; i < len(d.arr); i++ {
		d.arr[i] = i
		d.cnt[i] = 1
	}
}
