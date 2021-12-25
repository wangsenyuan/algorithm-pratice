package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	_, m, k := readThreeNums(reader)
	points := make([][]int, k)
	for i := 0; i < k; i++ {
		points[i] = readNNums(reader, 4)
	}
	fmt.Println(solve(m, k, points))
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

func solve(m int, k int, points [][]int) int64 {
	deliveries := make([]Delivery, k)
	in := make([][]int, m+2)
	out := make([][]int, m+2)
	rin := make([][]int, m+2)
	rout := make([][]int, m+2)

	for i := 0; i < len(in); i++ {
		in[i] = make([]int, 0, 10)
		out[i] = make([]int, 0, 10)
		rin[i] = make([]int, 0, 10)
		rout[i] = make([]int, 0, 10)
	}
	var cost int64
	for i := 0; i < k; i++ {
		cur := NewDelivery(points[i][0], points[i][1], points[i][2], points[i][3])
		deliveries[i] = cur
		cost += cur.value2()

		dist := abs(cur.x1-cur.x2) / 4
		l := max(0, cur.y1-dist)
		r := min(m+1, cur.y2+dist)
		in[l] = append(in[l], i)
		out[r] = append(out[r], i)
		rin[cur.y1] = append(rin[cur.y1], i)
		rout[cur.y2] = append(rout[cur.y2], i)
	}
	var incoming int
	var outgoing int
	var saving int64
	for _, j := range in[0] {
		incoming++
		cur := deliveries[j]
		saving += int64(abs(cur.x1-cur.x2)) - 4*int64(cur.y1)
	}

	var ans int64
	for i := 1; i <= m; i++ {
		saving += 4 * int64(incoming)
		saving -= 4 * int64(outgoing)

		for _, j := range in[i] {
			incoming++
			cur := deliveries[j]
			saving += int64(abs(cur.x1-cur.x2)) - 4*int64(abs(i-cur.y1))
		}

		if saving > ans {
			ans = saving
		}

		outgoing += len(rout[i])
		incoming -= len(rin[i])

		for _, j := range out[i] {
			outgoing--
			cur := deliveries[j]
			saving -= int64(abs(cur.x1-cur.x2)) - 4*int64(abs(i-cur.y2))
		}
	}
	return cost - ans
}

type Delivery struct {
	x1, y1, x2, y2 int
}

func NewDelivery(x1, y1, x2, y2 int) Delivery {
	if y1 > y2 {
		x1, x2 = x2, x1
		y1, y2 = y2, y1
	}
	return Delivery{x1, y1, x2, y2}
}

func (this Delivery) value2() int64 {
	return 2 * (int64(abs(this.y2-this.y1)) + int64(abs(this.x2-this.x1)))
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

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
