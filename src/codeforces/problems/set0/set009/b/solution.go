package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, vb, vs := readThreeNums(reader)
	X := readNNums(reader, n)
	room := readNNums(reader, 2)

	fmt.Println(solve(n, vb, vs, X, room))
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
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') {
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

func solve(n int, vb int, vs int, X []int, exam []int) int {
	time := make([]float64, n)
	v1 := float64(vs)
	v2 := float64(vb)
	var best = math.MaxFloat64
	for i := 0; i < n; i++ {
		t1 := float64(X[i]) / v2
		t2 := distance(exam[0], exam[1], X[i], 0) / v1
		time[i] = t1 + t2
		if i > 0 && time[i] < best {
			best = time[i]
		}
	}
	var pos int

	var dist = math.MaxFloat64
	for i := 0; i < n; i++ {
		if math.Abs(time[i]-best) < 1e-9 {
			tmp := distance(exam[0], exam[1], X[i], 0)
			if tmp < dist {
				dist = tmp
				pos = i
			}
		}
	}

	return pos + 1
}

func distance(a, b, c, d int) float64 {
	dx := float64(c - a)
	dy := float64(d - b)
	return math.Sqrt(dx*dx + dy*dy)
}
