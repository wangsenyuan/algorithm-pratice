package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var buf bytes.Buffer
	tc := readNum(reader)

	for tc > 0 {
		tc--
		m, n := readTwoNums(reader)
		red := make([][]int, m)
		for i := 0; i < m; i++ {
			red[i] = readNNums(reader, 2)
		}
		blue := make([][]int, n)
		for i := 0; i < n; i++ {
			blue[i] = readNNums(reader, 2)
		}
		res := solve(red, blue)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}

	fmt.Print(buf.String())
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

const eps = 1e-10

func solve(red, blue [][]int) int {
	points := make([]Point, 0, len(red)+len(blue))
	for i := 0; i < len(red); i++ {
		points = append(points, Point{red[i][0], red[i][1], 0})
	}

	for i := 0; i < len(blue); i++ {
		points = append(points, Point{blue[i][0], blue[i][1], 1})
	}

	vp := make([]Pair, 2*(len(red)+len(blue)-1))
	cnt := make([]int, 2)
	check := func(idx int) int {
		var p int
		for i := 0; i < len(points); i++ {
			if i == idx {
				continue
			}
			dx := points[i].x - points[idx].x
			dy := points[i].y - points[idx].y
			u := math.Atan2(float64(dy), float64(dx))
			vp[p] = Pair{u, i}
			p++
			vp[p] = Pair{u + 2*math.Pi, i}
			p++
		}

		sort.Sort(Pairs(vp))

		ans := len(points)

		color := points[idx].color
		cnt[0] = 0
		cnt[1] = 0

		n := len(points)
		for i, j := 0, 0; i < n-1; i++ {
			for j < p && vp[j].first-vp[i].first <= math.Pi-eps {
				cnt[points[vp[j].second].color]++
				j++
			}
			// idx inside, vp[i].second inside.
			cnt[color]++
			ans = min3(ans, cnt[0]+len(blue)-cnt[1], cnt[1]+len(red)-cnt[0])

			// idx inside, vp[i].second outside.
			cnt[points[vp[i].second].color]--
			ans = min3(ans, cnt[0]+len(blue)-cnt[1], cnt[1]+len(red)-cnt[0])

			// idx outside, vp[i].second outside.
			cnt[color]--
			ans = min3(ans, cnt[0]+len(blue)-cnt[1], cnt[1]+len(red)-cnt[0])
			// idx outside, vp[i].second inside.
			cnt[points[vp[i].second].color]++
			ans = min3(ans, cnt[0]+len(blue)-cnt[1], cnt[1]+len(red)-cnt[0])
			// Move vp[i].second outside.
			cnt[points[vp[i].second].color]--
		}

		return ans
	}
	ans := len(points)
	for i := 0; i < len(points); i++ {
		ans = min(ans, check(i))
	}
	return ans
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func min3(a, b, c int) int {
	if a <= b && a <= c {
		return a
	}
	if b <= a && b <= c {
		return b
	}
	return c
}

type Pair struct {
	first  float64
	second int
}

type Pairs []Pair

func (ps Pairs) Len() int {
	return len(ps)
}

func (ps Pairs) Less(i, j int) bool {
	return ps[i].first < ps[j].first
}

func (ps Pairs) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}

type Point struct {
	x, y  int
	color int
}
