package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	res := process(reader)

	fmt.Println(res)
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

func process(reader *bufio.Reader) int {
	n := readNum(reader)
	points := make([][]int, n)
	for i := range n {
		points[i] = readNNums(reader, 2)
	}
	return solve(points)
}

func solve(points [][]int) int {
	ys := make(map[int]int)
	xs := make(map[int]int)
	for _, p := range points {
		ys[p[1]]++
		xs[p[0]]++
	}

	var arr []int
	for y := range ys {
		arr = append(arr, y)
	}
	sort.Ints(arr)
	for i, y := range arr {
		ys[y] = i
	}
	var arr2 []int
	for x := range xs {
		arr2 = append(arr2, x)
	}
	sort.Ints(arr2)
	for i, x := range arr2 {
		xs[x] = i
	}

	at := make([][]int, len(arr))
	for _, p := range points {
		x, y := xs[p[0]], ys[p[1]]
		at[y] = append(at[y], x)
	}

	n := len(arr2)
	freq := make(fenwick, n+10)

	var res int

	for i := len(arr) - 1; i >= 0; i-- {
		sort.Ints(at[i])
		var sum int
		// 如果当前位置是r
		for j := 0; j < len(at[i]); j++ {

			var pref int
			if j == 0 {
				pref = freq.getRange(0, at[i][j]-1)
			} else {
				pref = freq.getRange(at[i][j-1]+1, at[i][j]-1)
			}
			pref++

			sum += pref

			var suf int
			if j+1 < len(at[i]) {
				suf = freq.getRange(at[i][j]+1, at[i][j+1]-1)
			} else {
				suf = freq.getRange(at[i][j]+1, n)
			}
			suf++
			res += sum * suf
		}

		for j := 0; j < len(at[i]); j++ {
			x := at[i][j]
			if freq.getRange(x, x) == 0 {
				// 不能重复计算
				freq.update(x, 1)
			}
		}
	}

	return res
}

type fenwick []int

func (tr fenwick) update(i int, v int) {
	i++
	for i < len(tr) {
		tr[i] += v
		i += i & -i
	}
}

func (tr fenwick) get(i int) int {
	i++
	var res int
	for i > 0 {
		res += tr[i]
		i -= i & -i
	}
	return res
}

func (tr fenwick) getRange(l int, r int) int {
	if l > r {
		return 0
	}
	res := tr.get(r)
	if l > 0 {
		res -= tr.get(l - 1)
	}
	return res
}
