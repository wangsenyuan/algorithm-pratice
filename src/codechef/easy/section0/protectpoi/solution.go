package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func readInt(bytes []byte, from int, val *int) int {
	i := from
	tmp := 0
	sign := 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
}

func readNum(scanner *bufio.Scanner) (a int) {
	scanner.Scan()
	readInt(scanner.Bytes(), 0, &a)
	return
}

func readTwo(scanner *bufio.Scanner) (a int, b int) {
	scanner.Scan()
	x := readInt(scanner.Bytes(), 0, &a)
	readInt(scanner.Bytes(), x+1, &b)
	return
}

func readThree(scanner *bufio.Scanner) (a int, b int, c int) {
	scanner.Scan()
	x := readInt(scanner.Bytes(), 0, &a)
	x = readInt(scanner.Bytes(), x+1, &b)
	readInt(scanner.Bytes(), x+1, &c)
	return
}

func readFour(scanner *bufio.Scanner) (a int, b int, c int, d int) {
	scanner.Scan()
	x := readInt(scanner.Bytes(), 0, &a)
	x = readInt(scanner.Bytes(), x+1, &b)
	x = readInt(scanner.Bytes(), x+1, &c)
	x = readInt(scanner.Bytes(), x+1, &d)
	return
}

func readNums(scanner *bufio.Scanner, n int) []int {
	res := make([]int, n)
	x := -1
	scanner.Scan()
	bs := scanner.Bytes()
	for i := 0; i < n; i++ {
		x = readInt(bs, x+1, &res[i])
	}
	return res
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	t := readNum(scanner)
	for i := 0; i < t; i++ {
		n, k, m := readThree(scanner)
		snakes := make([][]int, m)
		for j := 0; j < m; j++ {
			snakes[j] = readNums(scanner, 4)
		}
		fmt.Println(solve(n, k, m, snakes))
	}
}

func solve(n int, K int, m int, snakes [][]int) int {
	hs := make([][]int, m)
	vs := make([][]int, m)
	j, k := 0, 0
	for i := 0; i < m; i++ {
		a, b, c, d := snakes[i][0], snakes[i][1], snakes[i][2], snakes[i][3]
		if a > c {
			a, c = c, a
		}
		if b > d {
			b, d = d, b
		}
		if a == c {
			hs[j] = []int{b, d}
			j++
			vs[k] = []int{a, a}
			k++
		} else {
			vs[k] = []int{a, c}
			k++
			hs[j] = []int{b, b}
			j++
		}
	}

	h := cover(n, K, hs)
	if h < 0 {
		return -1
	}
	v := cover(n, K, vs)
	if v < 0 {
		return -1
	}
	return h + v
}

func cover(n int, k int, snakes [][]int) int {
	m := len(snakes)

	xs := make([]S, m)
	for i := 0; i < m; i++ {
		x := S{snakes[i][0], snakes[i][1]}
		xs[i] = x
	}

	sort.Sort(SS(xs))

	var ans int
	start := (n-k)/2 + 1
	end := start + k
	i := 0
	tmp := -1
	for start < end && i < m {
		if xs[i].v <= start {
			if xs[i].w >= start {
				tmp = max(tmp, xs[i].w)
			}
			i++
		} else {
			if tmp >= start {
				ans++
				start = tmp + 1
				tmp = -1
			} else {
				return -1
			}
		}
	}
	if tmp >= start {
		start = tmp + 1
		ans++
	}
	if start < end {
		return -1
	}
	return ans
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

type S struct {
	v int
	w int
}

type SS []S

func (this SS) Len() int {
	return len(this)
}
func (this SS) Less(i, j int) bool {
	return this[i].v < this[j].v
}
func (this SS) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}
