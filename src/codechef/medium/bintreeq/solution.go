package main

import (
	"bufio"
	"fmt"
	"os"
)

func readInt(bytes []byte, from int, val *int) int {
	i := from
	tmp := 0
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp
	return i
}

func readNum(scanner *bufio.Scanner) (a int) {
	scanner.Scan()
	readInt(scanner.Bytes(), 0, &a)
	return
}

func readTwoNums(scanner *bufio.Scanner) (a int, b int) {
	scanner.Scan()
	x := readInt(scanner.Bytes(), 0, &a)
	readInt(scanner.Bytes(), x+1, &b)
	return
}

func readNNums(scanner *bufio.Scanner, n int) []int {
	res := make([]int, n)
	x := -1
	scanner.Scan()
	for i := 0; i < n; i++ {
		x = readInt(scanner.Bytes(), x+1, &res[i])
	}
	return res
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	t := readNum(scanner)

	for t > 0 {
		tmp := readNNums(scanner, 3)
		n, u, v := tmp[0], tmp[1], tmp[2]
		fmt.Println(solve(n, u, v))

		t--
	}
}

func solve(n, u, v int) int {
	L := lca(u, v)

	if L == 0 {
		panic(fmt.Sprintf("failed to get LCA of %d, %d", u, v))
	}

	// 0 means p * 2, 1 means p * 2 + 1
	find := func(u, L int) []int {
		res := make([]int, 32)
		var i int
		for u != L {
			res[i] = u & 1
			u >>= 1
			i++
		}
		for j, k := 0, i-1; j < k; j, k = j+1, k-1 {
			res[j], res[k] = res[k], res[j]
		}
		return res[:i]
	}

	pu := find(u, L)
	pv := find(v, L)

	apply := func(x int, p []int) int64 {
		y := int64(x)
		for i := 0; i < len(p); i++ {
			y = y<<1 + int64(p[i])
		}
		return y
	}

	check := func(x int) bool {
		w := apply(x, pu)
		t := apply(x, pv)
		if w > int64(n) {
			return false
		}
		if t > int64(n) {
			return false
		}
		return true
	}

	lo, hi := 1, n+1

	for lo < hi {
		mid := lo + (hi-lo)>>1
		if check(mid) {
			lo = mid + 1
		} else {
			hi = mid
		}
	}

	return lo - 1
}

func lca(u, v int) int {
	for u != v {
		if u > v {
			u >>= 1
		} else {
			v >>= 1
		}
	}
	return u
}
