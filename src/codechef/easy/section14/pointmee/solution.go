package main

import (
	"bufio"
	"bytes"
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
	var buf bytes.Buffer
	tc := readNum(reader)

	for tc > 0 {
		tc--
		n := readNum(reader)
		X := readNNums(reader, n)
		Y := readNNums(reader, n)
		res := solve(n, X, Y)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}
	fmt.Print(buf.String())
}

func solve(n int, X []int, Y []int) int {
	for i := 0; i < n; i++ {
		X[i] *= 2
		Y[i] *= 2
	}
	P := make([]Pair, 0, 6*n*n)

	add := func(x, y int) {
		P = append(P, Pair{x, y})
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j {
				continue
			}
			a := X[i]
			b := Y[i]
			c := X[j]
			d := Y[j]
			add(a, d)
			add(c, b+c-a)
			add(c, b-c+a)
			add(a+b-d, d)
			add(a-b+d, d)
			add((a+b+c-d)/2, (b+a-c+d)/2)
		}
	}

	ans := 2 * n

	for i := 0; i < len(P); i++ {
		cur := P[i]
		top := cur.second
		right := cur.first
		var cnt int
		for j := 0; j < n; j++ {
			if right == X[j] && top == Y[j] {
				continue
			}
			if right == X[j] || top == Y[j] {
				cnt++
			} else if abs(right-X[j]) == abs(top-Y[j]) {
				cnt++
			} else {
				cnt += 2
			}
		}
		ans = min(ans, cnt)
	}

	return ans
}

type Pair struct {
	first, second int
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
