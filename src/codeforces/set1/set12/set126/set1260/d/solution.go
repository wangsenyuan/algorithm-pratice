package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	m, n, k, t := readFourNums(reader)
	a := readNNums(reader, m)
	traps := make([][]int, k)
	for i := 0; i < k; i++ {
		traps[i] = readNNums(reader, 3)
	}

	res := solve(a, n, traps, t)

	fmt.Println(res)
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

func readFourNums(reader *bufio.Reader) (a int, b int, c int, d int) {
	res := readNNums(reader, 4)
	a, b, c, d = res[0], res[1], res[2], res[3]
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

const inf = 1 << 60

func solve(soiders []int, n int, traps [][]int, t int) int {
	sort.Ints(soiders)
	at := make([]int, n+2)
	diff := make([]int, n+2)
	check := func(pos int) bool {
		w := soiders[pos]
		clear(at)
		clear(diff)

		for _, trap := range traps {
			l, r, d := trap[0], trap[1], trap[2]
			if d > w {
				at[l] = max(at[l], r)
			}
		}

		for i := 0; i <= n; i++ {
			if at[i] > 0 {
				diff[i]++
				diff[at[i]+1]--
			}
		}

		var res int

		for i := 1; i <= n; i++ {
			diff[i] += diff[i-1]
			if diff[i] > 0 {
				res++
			}
		}
		res *= 2
		res += n + 1
		return res <= t
	}

	res := sort.Search(len(soiders), check)
	return len(soiders) - res
}