package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	var buf bytes.Buffer

	for tc > 0 {
		tc--
		n, m, k := readThreeNums(reader)
		points := make([][]int, k)
		for i := 0; i < k; i++ {
			points[i] = readNNums(reader, 2)
		}
		plots, res := solve(n, m, points)
		buf.WriteString(fmt.Sprintf("%d\n", plots))
		for i := 0; i < k; i++ {
			buf.WriteString(fmt.Sprintf("%d ", res[i]))
		}
		buf.WriteByte('\n')
	}

	fmt.Print(buf.String())
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

func solve(n int, m int, fountains [][]int) (int, []int) {
	k := len(fountains)

	points := make([]Point, k+1)

	for i := 0; i < k; i++ {
		points[i] = Point{i, fountains[i][0], fountains[i][1]}
	}
	points[k] = Point{k, 0, 0}
	k++

	sort.Slice(points, func(i, j int) bool {
		return points[i].x > points[j].x || points[i].x == points[j].x && points[i].y < points[j].y
	})

	ans := make([]int, k+1)
	tot := make([]int, k+1)
	cur := make([]int, k+1)
	lst := make([]int, k+1)

	for i := 0; i < k; i++ {
		cur[i] = m + 1
		lst[i] = n
	}

	for i := 1; i <= k; i++ {
		tot[i] = tot[i-1]
		cur[i] = cur[i-1]
		lst[i] = lst[i-1]
		e := points[i-1]
		if e.y < cur[i] {
			ans[e.id] = 1
			tot[i] += (cur[i] - 1) * (lst[i] - e.x)
			cur[i] = e.y
			lst[i] = e.x
		}
	}

	for i := 1; i <= k; i++ {
		e := points[i-1]
		if ans[e.id] == 0 {
			continue
		}
		t := tot[i-1]
		cr := cur[i-1]
		lt := lst[i-1]
		for j := i + 1; j <= k; j++ {
			ee := points[j-1]
			if cr > ee.y {
				t += (cr - 1) * (lt - ee.x)
				cr = ee.y
				lt = ee.x
			}
			if ans[ee.id] == 1 {
				ans[e.id] = t - tot[j]
				break
			}
		}
	}

	return tot[k], ans[:k-1]
}

type Point struct {
	id   int
	x, y int
}
