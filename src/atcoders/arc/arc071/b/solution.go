package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, m := readTwoNums(reader)
	A := readNNums(reader, n)
	B := readNNums(reader, m)
	res := solve(A, B)

	fmt.Println(res)
}

func readString(r *bufio.Reader) string {
	s, _ := r.ReadString('\n')
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

const mod = 1e9 + 7

func add(a, b int) int {
	if a < 0 {
		a += mod
	}
	if b < 0 {
		b += mod
	}
	a += b
	if a >= mod {
		a -= mod
	}
	return a
}

func sub(a, b int) int {
	if b < 0 {
		b += mod
	}
	return add(a, mod-b)
}

func mul(a, b int) int {
	if a < 0 {
		a += mod
	}
	if b < 0 {
		b += mod
	}
	return int(int64(a) * int64(b) % mod)
}

func solve(xs []int, ys []int) int {
	// 对于 i和j,前面的 i * j 个点都可以和它组成一个矩形
	// 先划定一条线 y = x[i], 如何能计算出它左面所有的面积和
	// 然后再移动这条线到下一个个点，此时所有左边的面积增加 len * w
	var h int
	var ty int
	for i, y := range ys {
		// total = yi - y0 + yi - y1 ... + yi - yi-1
		ty = add(ty, sub(mul(i, y), h))
		h = add(h, y)
	}
	var w int
	var tx int
	for i, x := range xs {
		tmp := sub(mul(i, x), w)
		tx = add(tx, tmp)
		w = add(w, x)
	}

	return mul(tx, ty)
}
