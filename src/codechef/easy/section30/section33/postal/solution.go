package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)
	var buf bytes.Buffer
	for tc > 0 {
		tc--
		n := readNum(reader)
		pos := readNNums(reader, n)
		dirs := readNNums(reader, n)
		q := readNum(reader)
		questions := make([][]int, q)
		for i := 0; i < q; i++ {
			questions[i] = readNNums(reader, 2)
		}

		res := solve(pos, dirs, questions)
		for _, cur := range res {
			buf.WriteString(fmt.Sprintf("%d %d\n", cur[0], cur[1]))
		}
	}
	fmt.Print(buf.String())
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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
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

func solve(pos []int, initial_dirs []int, questions [][]int) [][]int {
	n := len(pos)

	final_pos := make([]int, n)
	sorted_pos := make([]int, n)

	ans := make([][]int, len(questions))

	for i, cur := range questions {
		I, T := cur[0], cur[1]
		for j := 0; j < n; j++ {
			if initial_dirs[j] == 0 {
				// go north
				final_pos[j] = pos[j] - T
			} else {
				final_pos[j] = pos[j] + T
			}
			sorted_pos[j] = final_pos[j]
		}
		sort.Ints(sorted_pos)
		var rev_count int
		for j := 0; j < n; j++ {
			if final_pos[j] == sorted_pos[I] {
				rev_count++
				continue
			}
			rank := search(n, func(k int) bool {
				return sorted_pos[k] >= final_pos[j]
			})

			if (I-j)*(I-rank) <= 0 {
				rev_count++
			}
		}
		rev_count--

		ans[i] = []int{sorted_pos[I], rev_count}
	}

	return ans
}

func search(n int, f func(int) bool) int {
	l, r := 0, n
	for l < r {
		mid := (l + r) / 2
		if f(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return r
}
