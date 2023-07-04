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

	tc := readNum(reader)

	var buf bytes.Buffer
	for tc > 0 {
		tc--
		pts := make([][]int, 4)
		for i := 0; i < 4; i++ {
			pts[i] = readNNums(reader, 2)
		}
		res := solve(pts)
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

func solve(pts [][]int) int64 {
	id := []int{0, 1, 2, 3}

	var ans int64 = math.MaxInt64

	for {
		var cur int64
		x1 := getSeg(pts[id[0]][0], pts[id[3]][0])
		x2 := getSeg(pts[id[1]][0], pts[id[2]][0])
		cur += length(x1) + length(x2)
		xs := getOpt(x1, x2)
		y1 := getSeg(pts[id[0]][1], pts[id[1]][1])
		y2 := getSeg(pts[id[2]][1], pts[id[3]][1])
		cur += length(y1) + length(y2)
		ys := getOpt(y1, y2)
		is := int64(min(xs[1], ys[1])) - int64(max(xs[0], ys[0]))
		if is < 0 {
			cur += 2 * (-is)
		}
		if cur < ans {
			ans = cur
		}
		if !nextPermutation(id) {
			break
		}
	}

	return ans
}

func length(seg []int) int64 {
	return int64(seg[1] - seg[0])
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	return a + b - min(a, b)
}

func getSeg(a int, b int) []int {
	return []int{min(a, b), max(a, b)}
}

func getOpt(a, b []int) []int {
	return []int{
		max(max(a[0]-b[1], b[0]-a[1]), 0),
		max(max(b[1]-a[0], a[1]-b[0]), 0),
	}
}

func solve1(pts [][]int) int64 {
	var ans int64 = math.MaxInt64

	fp := make([][]int, 4)
	for i := 0; i < 4; i++ {
		fp[i] = make([]int, 2)
	}

	assign := func(arr ...int) {
		for i := 0; i < len(arr); i++ {
			fp[i/2][i%2] = arr[i]
		}
	}

	compare := func(i, j int) bool {
		return fp[i][0] < fp[j][0] || fp[i][0] == fp[j][0] && fp[i][1] < fp[j][1]
	}

	id := []int{0, 1, 2, 3}

	for st := 0; st < 2; st++ {
		for idx1 := 0; idx1 < 4; idx1++ {
			for idx2 := 0; idx2 < 4; idx2++ {
				for idy1 := 0; idy1 < 4; idy1++ {
					x1 := pts[idx1][0]
					x2 := pts[idx2][0]
					y1 := pts[idy1][1]
					if x1 > x2 {
						continue
					}
					for k := -1; k <= 1; k += 2 {
						y2 := y1 + k*(x1-x2)
						assign(x1, y1, x2, y1, x2, y2, x1, y2)
						sort.Slice(fp, compare)
						for j := 0; j < 4; j++ {
							id[j] = j
						}
						for {
							var cur int64
							for j := 0; j < 4; j++ {
								cur += distance(fp[id[j]], pts[j])
							}
							if cur < ans {
								ans = cur
							}
							if !nextPermutation(id) {
								break
							}
						}
					}
				}
			}
		}

		for i := 0; i < 4; i++ {
			pts[i][0], pts[i][1] = pts[i][1], pts[i][0]
		}
	}

	return ans
}

func nextPermutation(arr []int) bool {
	n := len(arr)
	i := n - 1
	for i > 0 && arr[i-1] > arr[i] {
		i--
	}
	//arr[i-1] < arr[i]
	if i == 0 {
		//unable to permute
		return false
	}
	i--
	j := i + 1
	for k := j; k < n; k++ {
		if arr[k] > arr[i] {
			j = k
		}
	}
	arr[i], arr[j] = arr[j], arr[i]

	for a, b := i+1, n-1; a < b; a, b = a+1, b-1 {
		arr[a], arr[b] = arr[b], arr[a]
	}
	return true
}

func distance(a, b []int) int64 {
	dx := abs(a[0] - b[0])
	dy := abs(a[1] - b[1])
	return int64(dx) + int64(dy)
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
