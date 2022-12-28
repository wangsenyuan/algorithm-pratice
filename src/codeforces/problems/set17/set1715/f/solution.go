package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, m := readTwoNums(reader)

	query := func(poly [][]int) float64 {
		fmt.Printf("? %d\n", len(poly))
		for _, p := range poly {
			fmt.Printf("%d %d\n", p[0], p[1])
		}
		return readFloat(reader)
	}

	x, y := solve(n, m, query)

	fmt.Printf("! %.10f %.10f\n", x, y)
}

func readFloat(reader *bufio.Reader) float64 {
	s := readString(reader)
	r, _ := strconv.ParseFloat(s, 64)
	return r
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

func normalize(s string) string {

	for i := len(s); i > 0; i-- {
		if s[i-1] >= 'a' && s[i-1] <= 'z' {
			return s[:i]
		}
	}
	return ""
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

func solve(n int, m int, query func([][]int) float64) (float64, float64) {

	var poly [][]int

	poly = append(poly, []int{0, m + 1})
	poly = append(poly, []int{0, 0})

	for i := 1; i < n; i++ {
		poly = append(poly, []int{i, m})
		poly = append(poly, []int{i, 0})
	}

	poly = append(poly, []int{n, m})
	poly = append(poly, []int{n, m + 1})

	y := float64(m)*query(poly) - 0.5

	poly = poly[:0]

	poly = append(poly, []int{n + 1, 0})
	poly = append(poly, []int{0, 0})

	for i := 1; i < m; i++ {
		poly = append(poly, []int{n, i})
		poly = append(poly, []int{0, i})
	}
	poly = append(poly, []int{n, m})
	poly = append(poly, []int{n + 1, m})

	x := float64(n)*query(poly) - 0.5

	return x, y
}
