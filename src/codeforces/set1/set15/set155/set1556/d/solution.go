package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n, k := readTwoNums(reader)

	query := func(op string, i int, j int) int {
		fmt.Printf("%s %d %d\n", op, i, j)
		return readNum(reader)
	}

	res := solve(n, k, query)

	fmt.Println("finish ", res)
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

func solve(n int, k int, query func(string, int, int) int) int {
	// a | b + a & b = a + b

	get := func(i int, j int) int {
		x := query("or", i+1, j+1)
		y := query("and", i+1, j+1)
		return x + y
	}

	x := get(0, 1)
	y := get(0, 2)
	z := get(1, 2)
	// a[0] + a[1] = x
	// a[0] + a[2] = y
	// a[1] + a[2] = z
	// 2 * a[1] = x - y + z
	a := make([]int, n)
	a[1] = (x - y + z) / 2
	a[0] = x - a[1]
	a[2] = z - a[1]
	for i := 3; i < n; i++ {
		x := get(0, i)
		a[i] = x - a[0]
	}
	sort.Ints(a)
	return a[k-1]
}
