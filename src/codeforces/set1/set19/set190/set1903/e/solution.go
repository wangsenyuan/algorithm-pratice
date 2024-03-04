package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	comp := func() int {
		return readNum(reader)
	}
	play := func(i int) {
		fmt.Printf("%d\n", i)
	}

	role := func(r string) {
		fmt.Printf("%s\n", r)
	}

	for tc > 0 {
		tc--
		n := readNum(reader)
		s := readNNums(reader, 2)
		pts := make([][]int, n)
		for i := 0; i < n; i++ {
			pts[i] = readNNums(reader, 2)
		}
		solve(s, pts, role, play, comp)
	}

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

func solve(s []int, pts [][]int, role func(string), play func(int), comp func() int) {
	v := (s[0] & 1) ^ (s[1] & 1)
	p := make([]map[int]bool, 2)
	for i := 0; i < 2; i++ {
		p[i] = make(map[int]bool)
	}
	for i, x := range pts {
		j := (x[0] & 1) ^ (x[1] & 1)
		p[j][i+1] = true
	}

	n := len(pts)

	if len(p[v]) >= len(p[v^1]) {
		role("First")
		v ^= 1
		for i := 0; i < n; i += 2 {
			if len(p[v]) > 0 {
				play(pick(p[v]))
			} else {
				play(pick(p[v^1]))
			}
			if i+1 < n {
				j := comp()
				if p[0][j] {
					delete(p[0], j)
				} else {
					delete(p[1], j)
				}
			}
		}
	} else {
		role("Second")
		for i := 0; i < n; i += 2 {
			j := comp()
			if p[0][j] {
				delete(p[0], j)
			} else {
				delete(p[1], j)
			}
			if i+1 < n {
				if len(p[v]) > 0 {
					play(pick(p[v]))
				} else {
					play(pick(p[v^1]))
				}
			}
		}
	}
}

func pick(r map[int]bool) int {
	for k := range r {
		delete(r, k)
		return k
	}
	return -1
}
