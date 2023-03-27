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
	var buf bytes.Buffer
	tc := readNum(reader)

	for tc > 0 {
		tc--
		n, k := readTwoNums(reader)
		pts := make([][]int, n)
		for i := 0; i < n; i++ {
			pts[i] = readNNums(reader, 2)
		}
		res := solve(k, pts)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}

	fmt.Print(buf.String())
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' {
			return s[:i]
		}
	}
	return s
}

func readNInt64s(reader *bufio.Reader, n int) []int64 {
	res := make([]int64, n)
	s, _ := reader.ReadBytes('\n')

	var pos int

	for i := 0; i < n; i++ {
		pos = readInt64(s, pos, &res[i]) + 1
	}

	return res
}

func readInt64(bytes []byte, from int, val *int64) int {
	i := from
	var sign int64 = 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	var tmp int64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int64(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
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

const INF = 1e18

func solve(K int, points [][]int) int64 {
	ycoords := getSortedYCoords(points)

	sort.Sort(Values(points))
	n := len(points)

	lct := NewFT(n)
	rct := NewFT(n)
	lsum := NewFT(n)
	rsum := NewFT(n)

	check := func(x int, lt int) int64 {
		rt := K - lt
		L := lct.lowerBound(int64(lt))
		R := rct.lowerBound(int64(rt))
		if L == n || R == n {
			return INF
		}
		tmp := lsum.get(L+1) + rsum.get(R+1)
		tmp += int64(x) * int64(lt-rt)

		return tmp
	}

	leftPos := make([]int, n)
	rightPos := make([]int, n)
	leftValues := make([][]int, n)
	rightValues := make([][]int, n)
	var ans int64 = INF

	for _, Y := range ycoords {
		lct.reset()
		rct.reset()
		lsum.reset()
		rsum.reset()

		for i := 0; i < n; i++ {
			x, y := points[i][0], points[i][1]
			leftValues[i] = []int{abs(Y-y) - x, i}
			rightValues[i] = []int{abs(Y-y) + x, i}
		}

		sort.Sort(Values(leftValues))
		sort.Sort(Values(rightValues))

		for i := 0; i < n; i++ {
			leftPos[leftValues[i][1]] = i
			rightPos[rightValues[i][1]] = i

			rct.update(i, 1)
			rsum.update(i, int64(rightValues[i][0]))
		}

		for i := 0; i < n; i++ {
			u := rightPos[i]
			v := leftPos[i]
			x, y := points[i][0], points[i][1]
			rval := int64(abs(Y-y) + x)
			lval := int64(abs(Y-y) - x)

			rct.update(u, -1)
			rsum.update(u, -rval)
			lct.update(v, 1)
			lsum.update(v, lval)

			if i > 0 && x == points[i-1][0] {
				continue
			}
			lo, hi := 0, K
			for hi-lo > 2 {
				mid := (hi + lo) / 2
				if check(x, mid) > check(x, mid+1) {
					lo = mid
				} else {
					hi = mid + 1
				}
			}
			for j := lo; j <= hi; j++ {
				ans = min(ans, check(x, j))
			}
		}
	}

	return ans
}

type Values [][]int

func (v Values) Len() int {
	return len(v)
}

func (v Values) Less(i, j int) bool {
	return v[i][0] < v[j][0] || v[i][0] == v[j][0] && v[i][1] < v[j][1]
}

func (v Values) Swap(i, j int) {
	v[i], v[j] = v[j], v[i]
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func min(a, b int64) int64 {
	if a <= b {
		return a
	}
	return b
}

func getSortedYCoords(pts [][]int) []int {
	ycoords := make(map[int]int)

	for _, pt := range pts {
		ycoords[pt[1]]++
	}
	var arr []int
	for y := range ycoords {
		arr = append(arr, y)
	}
	sort.Ints(arr)
	return arr
}

type FT struct {
	arr []int64
}

func NewFT(n int) *FT {
	arr := make([]int64, n)
	return &FT{arr}
}

func (ft *FT) update(pos int, diff int64) {
	for ; pos < len(ft.arr); pos |= (pos + 1) {
		ft.arr[pos] += diff
	}
}

func (ft *FT) get(pos int) int64 {
	var res int64
	for ; pos > 0; pos &= (pos - 1) {
		res += ft.arr[pos-1]
	}
	return res
}

func (ft *FT) reset() {
	for i := 0; i < len(ft.arr); i++ {
		ft.arr[i] = 0
	}
}

func (ft *FT) lowerBound(sum int64) int {
	if sum <= 0 {
		return -1
	}
	var pos int
	for pw := 1 << 12; pw > 0; pw >>= 1 {
		if pos+pw <= len(ft.arr) && ft.arr[pos+pw-1] < sum {
			pos += pw
			sum -= ft.arr[pos-1]
		}
	}
	return pos
}
