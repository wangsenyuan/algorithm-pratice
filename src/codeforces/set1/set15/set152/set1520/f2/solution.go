package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n, t := readTwoNums(reader)
	ask := func(l int, r int) int {
		fmt.Printf("? %d %d\n", l, r)
		return readNum(reader)
	}

	out := func(pos int) {
		fmt.Printf("! %d\n", pos)
	}

	play := func(f func(k int)) {
		f(readNum(reader))
	}

	solve(n, t, play, ask, out)
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

func solve(n int, t int, play func(f func(k int)), ask func(int, int) int, out func(int)) {
	type pair struct {
		first  int
		second int
	}
	cache := make(map[pair]int)

	var dfs func(pos int, l int, r int)

	dfs = func(pos int, l int, r int) {
		cache[pair{l, r}]--
		if l != r {
			mid := (l + r) / 2
			if pos <= mid {
				dfs(pos, l, mid)
			} else {
				dfs(pos, mid+1, r)
			}
		}
	}

	for t > 0 {
		t--
		play(func(k int) {
			l, r := 0, n-1
			for l < r {
				mid := (l + r) / 2
				if _, ok := cache[pair{l, mid}]; !ok {
					cache[pair{l, mid}] = mid - l + 1 - ask(l+1, mid+1)
				}
				if k <= cache[pair{l, mid}] {
					r = mid
				} else {
					k -= cache[pair{l, mid}]
					l = mid + 1
				}
			}
			out(l + 1)
			dfs(l, 0, n-1)
		})
	}
}
