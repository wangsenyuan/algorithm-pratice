package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

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

	n, m := readTwoNums(reader)
	r := readNNums(reader, n)

	res := solve(m, r)

	fmt.Println(res)
}

const inf = 1 << 30

func solve(m int, r []int) int {
	// n := len(r)
	score := make([]int, m+2)

	var cnt int

	add := make([]int, m+2)

	process := func() {
		for i := 0; i <= m; i++ {
			if i > 0 {
				add[i] += add[i-1]
			}
			score[i] += add[i]
		}
		clear(add)

		for i := m; i > 0; i-- {
			score[i] = max(score[i], score[i-1])
		}
	}

	for _, v := range r {
		if v == 0 {
			process()
			cnt++
			continue
		}
		if abs(v) > cnt {
			// not able to gain this score
			continue
		}
		if v > 0 {
			add[v]++
		} else {
			add[0]++
			add[cnt+v+1]--
		}
	}
	process()

	return slices.Max(score)
}

func solve1(m int, r []int) int {
	fp := make([]int, m+1)

	for i := 0; i <= m; i++ {
		fp[i] = -inf
	}
	fp[0] = 0
	pos := make([]int, m+1)
	neg := make([]int, m+1)
	n := len(r)
	var cnt int
	for i := 0; i <= n; i++ {
		if i == n || r[i] == 0 {
			for k := 1; k <= cnt; k++ {
				pos[k] += pos[k-1]
				neg[k] += neg[k-1]
			}

			for j := cnt; j > 0; j-- {
				fp[j] = max(fp[j], fp[j-1]) + pos[j] + neg[cnt-j]
			}

			fp[0] += neg[cnt]

			cnt++
			clear(pos)
			clear(neg)
		} else {
			if r[i] > 0 {
				pos[r[i]]++
			} else {
				neg[abs(r[i])]++
			}
		}
	}

	return slices.Max(fp)
}

func abs(num int) int {
	return max(num, -num)
}
