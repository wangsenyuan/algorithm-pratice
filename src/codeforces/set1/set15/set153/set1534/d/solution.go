package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n := readNum(reader)
	ask := func(r int) []int {
		fmt.Printf("? %d\n", r)
		return readNNums(reader, n)
	}

	res := solve(n, ask)
	var buf bytes.Buffer
	buf.WriteString("!\n")
	for _, x := range res {
		buf.WriteString(fmt.Sprintf("%d %d\n", x[0], x[1]))
	}
	fmt.Print(buf.String())
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

type pair struct {
	first  int
	second int
}

func solve(n int, ask func(int) []int) [][]int {

	g := make(map[pair]int)

	add := func(u, v int) {
		u, v = min(u, v), max(u, v)
		g[pair{u, v}]++
	}

	d := ask(1)

	cnt := make([]int, 2)
	for i := 0; i < n; i++ {
		cnt[d[i]&1]++
		if d[i] == 1 {
			add(0, i)
		}
	}
	var exp int
	if cnt[0] > cnt[1] {
		exp = 1
	}

	var arr []int
	for i := 0; i < n; i++ {
		if d[i]&1 == exp {
			arr = append(arr, i)
		}
	}

	for _, x := range arr {
		if x == 0 {
			continue
		}
		nd := ask(x + 1)
		for i := 0; i < n; i++ {
			if nd[i] == 1 {
				add(i, x)
			}
		}
	}
	
	// len(g) == n - 1
	var res [][]int

	for k := range g {
		u, v := k.first, k.second
		res = append(res, []int{u + 1, v + 1})
	}

	return res
}
