package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

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

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)
	pairs := make([][]int, n)

	for i := range n {
		pairs[i] = readNNums(reader, 2)
	}

	res := solve(pairs)

	fmt.Println(res)
}

const mod = 998244353

func add(a, b int) int {
	a += b
	if a >= mod {
		a -= mod
	}
	return a
}

func sub(a, b int) int {
	return add(a, mod-b)
}

func mul(a, b int) int {
	return a * b % mod
}

func solve(pairs [][]int) int {
	n := len(pairs)

	F := make([]int, n+1)
	F[0] = 1
	for i := 1; i <= n; i++ {
		F[i] = mul(i, F[i-1])
	}

	res := F[n]

	for u := 0; u < 2; u++ {
		slices.SortFunc(pairs, func(a, b []int) int {
			if a[u] != b[u] {
				return a[u] - b[u]
			}
			return a[1^u] - b[1^u]
		})
		bad := 1
		for i := 0; i < n; {
			j := i
			for i < n && pairs[i][u] == pairs[j][u] {
				i++
			}
			bad = mul(bad, F[i-j])
		}
		res = sub(res, bad)
	}

	slices.SortFunc(pairs, func(a, b []int) int {
		if a[0] != b[0] {
			return a[0] - b[0]
		}
		return a[1] - b[1]
	})

	found := false

	for i := 1; i < n; i++ {
		if pairs[i-1][1] > pairs[i][1] {
			found = true
			break
		}
	}

	if !found {
		bad3 := 1

		for i := 0; i < n; {
			j := i
			for i < n && pairs[i][0] == pairs[j][0] && pairs[i][1] == pairs[j][1] {
				i++
			}

			bad3 = mul(bad3, F[i-j])
		}

		res = add(res, bad3)
	}

	return res
}
