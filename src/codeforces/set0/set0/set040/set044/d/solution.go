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
	ps := make([][]int, n)

	for i := 0; i < n; i++ {
		ps[i] = readNNums(reader, 3)
	}

	res := solve(ps)

	fmt.Printf("%.9f\n", res)
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

func solve(planets [][]int) float64 {
	n := len(planets)

	ps := make([]Planet, n)

	for i := 0; i < n; i++ {
		ps[i] = Planet{planets[i][0], planets[i][1], planets[i][2]}
	}

	dist := make([]float64, n)

	for i := 1; i < n; i++ {
		dist[i] = ps[i].Dist(ps[0])
	}

	var res float64 = math.MaxFloat64

	for i := 1; i < n; i++ {
		for j := i + 1; j < n; j++ {
			res = math.Min(res, dist[i]+dist[j]+ps[i].Dist(ps[j]))
		}
	}

	return res / 2
}

type Planet struct {
	x int
	y int
	z int
}

func (this Planet) Dist(that Planet) float64 {
	dx := float64(this.x - that.x)
	dy := float64(this.y - that.y)
	dz := float64(this.z - that.z)
	return math.Sqrt(dx*dx + dy*dy + dz*dz)
}
