package main

import (
	"bufio"
	"fmt"
	"os"
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

func main() {
	reader := bufio.NewReader(os.Stdin)

	S, _ := reader.ReadString('\n')
	T, _ := reader.ReadString('\n')

	solver := NewSolver(S[:len(S)-1], T[:len(T)-1])

	q := readNum(reader)

	for q > 0 {
		q--
		n := readNum(reader)
		fmt.Println(solver.Query(n))
	}
}

type Solver struct {
	count []int
	sz    int
	c1    int
	c3    int
}

func NewSolver(S, T string) Solver {
	temp := T + "#" + S + S
	pi := kmp(temp)

	count := make([]int, 2*len(S))

	var c1, c2 int
	for i, j := len(T)+1, 0; i < len(temp); i, j = i+1, j+1 {
		if i == len(T)+1 {
			if pi[i] == len(T) {
				c1++
				c2++
				count[j] = 1
			}
			continue
		}
		count[j] = count[j-1]
		if pi[i] == len(T) {
			c2++
			if i <= len(S)+len(T) {
				c1++
			}
			count[j]++
		}
	}

	c3 := c2 - 2*c1

	return Solver{count, len(S), c1, c3}
}

func (solver Solver) Query(n int) int64 {
	quot := n / solver.sz
	rem := n % solver.sz

	var ans int64
	if quot > 0 {
		ans = int64(quot-1)*int64(solver.c3) + int64(quot)*int64(solver.c1)

		if rem != 0 {
			ans += int64(solver.count[solver.sz+rem-1] - solver.count[solver.sz-1])
		}
	} else {
		ans = int64(solver.count[rem-1])
	}

	return ans
}

func kmp(s string) []int {
	n := len(s)
	lps := make([]int, n)

	for i := 1; i < n; i++ {
		j := lps[i-1]

		for j > 0 && s[i] != s[j] {
			j = lps[j-1]
		}

		lps[i] = j
		if s[i] == s[j] {
			lps[i]++
		}
	}

	return lps
}
