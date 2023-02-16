package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	L, R := readTwoNums(reader)

	n, E := solve(L, R)

	var buf bytes.Buffer

	buf.WriteString(fmt.Sprintf("YES\n%d %d\n", n, len(E)))

	for i := 0; i < len(E); i++ {
		buf.WriteString(fmt.Sprintf("%d %d %d\n", E[i][0], E[i][1], E[i][2]))
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

func solve(L, R int) (int, [][]int) {
	var edges [][]int

	add := func(u, v, w int) {
		edges = append(edges, []int{u, v, w})
	}

	var process func(l, r int) int

	process = func(l, r int) int {
		if l > 1 {
			k := process(1, r-l+1)
			add(k, k+1, l-1)
			return k + 1
		}

		if (r & -r) == r {
			k := ctz(r)
			add(1, 2, 1)
			for i := 3; i <= k+2; i++ {
				add(1, i, 1)
				for j := 2; j < i; j++ {
					add(j, i, 1<<(j-2))
				}
			}
			return k + 2
		}

		var k int

		for 1<<(k+1) <= r-1 {
			k++
		}

		process(1, 1<<k)

		add(1, k+3, 1)
		for i := 0; i <= k; i++ {
			if ((r-1)>>i)&1 == 1 {
				add(i+2, k+3, 1+((r-1)>>(i+1)<<(i+1)))
			}
		}
		return k + 3
	}

	k := process(L, R)

	return k, edges
}

func ctz(num int) int {
	if num == 0 {
		return -1
	}
	var i int
	for num&1 == 0 {
		num >>= 1
		i++
	}
	return i
}
