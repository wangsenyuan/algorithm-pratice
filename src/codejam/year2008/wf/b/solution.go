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

func readTwoNums(scanner *bufio.Scanner) (a int, b int) {
	res := readNNums(scanner, 2)
	a, b = res[0], res[1]
	return
}

func readNNums(scanner *bufio.Scanner, n int) []int {
	res := make([]int, n)
	x := 0
	scanner.Scan()
	for i := 0; i < n; i++ {
		for x < len(scanner.Bytes()) && scanner.Bytes()[x] == ' ' {
			x++
		}
		x = readInt(scanner.Bytes(), x, &res[i])
	}
	return res
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	tc := readNum(scanner)

	for i := 1; i <= tc; i++ {
		W, H := readTwoNums(scanner)
		dx1, dy1 := readTwoNums(scanner)
		dx2, dy2 := readTwoNums(scanner)
		x, y := readTwoNums(scanner)
		fmt.Printf("Case #%d: %d\n", i, solve(W, H, dx1, dy1, dx2, dy2, x, y))
	}
}

func solve(W, H int, dx1 int, dy1 int, dx2 int, dy2 int, x, y int) int64 {
	if W < 100 && H < 100 || collinear(dx1, dy1, dx2, dy2) {
		return bfs(W, H, dx1, dy1, dx2, dy2, x, y)
	}

	inside := func(a, b int) bool {
		xx := x + a*dx1 + b*dx2
		yy := y + a*dy1 + b*dy2
		return xx >= 0 && xx < W && yy >= 0 && yy < H
	}

	b1, b2 := 0, 1000001

	var ans int64

	for a := 0; ; a++ {
		for !inside(a, b1) {
			b1++
			if b1 > b2 {
				return ans
			}
		}

		if inside(a, b2) {
			for inside(a, b2) {
				b2++
			}
			b2--
		} else {
			for !inside(a, b2) && b2 > b1 {
				b2--
				// fmt.Fprintf(os.Stderr, "[debug] %d %d %d\n", a, b1, b2)
			}
		}
		ans += int64(b2 - b1 + 1)
	}
}

func collinear(a, b, c, d int) bool {
	if a == 0 && b == 0 {
		return true
	}
	u := complex(float64(a), float64(b))
	v := complex(float64(c), float64(d))
	w := v / u
	return imag(w) == 0
}

func bfs(w, h int, dx1 int, dy1 int, dx2 int, dy2 int, x, y int) int64 {
	que := make([][]int, 0, 100)
	que = append(que, []int{x, y})
	var pos int
	vis := make(map[int64]bool)
	H := int64(h)
	for pos < len(que) {
		u, v := que[pos][0], que[pos][1]
		pos++
		k := int64(u)*H + int64(v)
		vis[k] = true
		a, b := u+dx1, v+dy1
		if a >= 0 && a < w && b >= 0 && b < h {
			k1 := int64(a)*H + int64(b)
			if !vis[k1] {
				que = append(que, []int{a, b})
				vis[k1] = true
			}
		}
		a, b = u+dx2, v+dy2
		if a >= 0 && a < w && b >= 0 && b < h {
			k1 := int64(a)*H + int64(b)
			if !vis[k1] {
				que = append(que, []int{a, b})
				vis[k1] = true
			}
		}
	}
	return int64(len(que))
}
