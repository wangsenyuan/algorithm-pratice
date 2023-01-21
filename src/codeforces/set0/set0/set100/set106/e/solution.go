package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n := readNum(reader)
	points := make([][]int, n)
	for i := 0; i < n; i++ {
		points[i] = readNNums(reader, 3)
	}
	res := solve(points)

	fmt.Printf("%.10f %.10f %.10f\n", res[0], res[1], res[2])
}

func readNFloats(reader *bufio.Reader, n int) []float64 {
	s, _ := reader.ReadBytes('\n')
	res := make([]float64, n)
	var pos int
	for i := 0; i < n; i++ {
		pos = readFloat64(s, pos, &res[i]) + 1
	}
	return res
}

func readFloat64(bytes []byte, from int, val *float64) int {
	i := from
	var sign float64 = 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	var real int64

	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		real = real*10 + int64(bytes[i]-'0')
		i++
	}

	if i == len(bytes) || bytes[i] != '.' {
		*val = float64(real) * sign
		return i
	}

	// bytes[i] == '.'
	i++

	var fraq float64
	var base float64 = 0.1
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		fraq += base * float64(bytes[i]-'0')
		base /= 10
		i++
	}

	*val = (float64(real) + fraq) * sign

	return i
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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
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

func solve(points [][]int) []float64 {
	var ox, oy, oz float64
	var steps = 0.5
	for i := 0; i < 60; i++ {

		for t := 0; t < 300; t++ {
			var who int
			var dist float64
			for j := 0; j < len(points); j++ {
				tmp := distance(points[j], ox, oy, oz)
				if tmp > dist {
					dist = tmp
					who = j
				}
			}

			ox += (float64(points[who][0]) - ox) * steps
			oy += (float64(points[who][1]) - oy) * steps
			oz += (float64(points[who][2]) - oz) * steps
		}
		steps /= 2
	}
	return []float64{ox, oy, oz}
}

func distance(p []int, x, y, z float64) float64 {
	dx := float64(p[0]) - x
	dy := float64(p[1]) - y
	dz := float64(p[2]) - z
	return math.Sqrt(dx*dx + dy*dy + dz*dz)
}
