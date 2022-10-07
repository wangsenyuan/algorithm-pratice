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
		n := readNum(reader)
		A := readNNums(reader, n)
		res := solve(A)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}

	fmt.Print(buf.String())
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

func solve(A []int) int {
	n := len(A)
	// a, b, a, b,
	cnt := make([]map[int]int, 2)
	for i := 0; i < 2; i++ {
		cnt[i] = make(map[int]int)
	}

	for i := 0; i < n; i++ {
		a := A[i]
		cnt[i&1][a]++
	}
	t := createPairs(cnt[1])

	sort.Slice(t, func(i, j int) bool {
		return t[i].second > t[j].second
	})

	rk := make(map[int]int)

	for i := 0; i < len(t); i++ {
		rk[t[i].first] = i + 1
	}

	ex := make(map[int]map[int]int)

	add := func(x int, y int, v int) {
		if len(ex[x]) == 0 {
			ex[x] = make(map[int]int)
		}
		ex[x][y] += v
	}

	mex := make(map[int]int)

	x := A[0]
	y := A[1]
	tmp := 2

	for i := 2; i < n; i++ {
		nx, ny := A[i], A[i-1]
		if i&1 == 1 {
			nx = A[i-1]
			ny = A[i]
		}
		if x == nx && y == ny {
			tmp++
		} else {
			if cnt[0][y] > 0 && cnt[1][x] > 0 {
				add(y, x, tmp>>1)
			}
			tmp = 2
			x = nx
			y = ny
		}
	}

	if cnt[0][y] > 0 && cnt[1][x] > 0 {
		add(y, x, tmp>>1)
	}

	ans := n

	for x, ys := range ex {
		var t2 []int

		for y, v := range ys {
			t2 = append(t2, rk[y])
			ans = min(ans, n-cnt[0][x]-cnt[1][y]+v)
		}
		if cnt[1][x] > 0 {
			t2 = append(t2, rk[x])
		}

		sort.Ints(t2)

		var flag bool

		for i := 0; i < len(t2); i++ {
			if t2[i] != i+1 {
				mex[x] = i + 1
				flag = true
				break
			}
		}

		if !flag {
			mex[x] = t2[len(t2)-1] + 1
		}
	}

	for x, y := range cnt[0] {
		if _, ok := ex[x]; ok {
			if mex[x] > len(t) {
				continue
			}
			ans = min(ans, n-y-t[mex[x]-1].second)
		} else {
			if t[0].first == x {
				if len(t) > 2 {
					ans = min(ans, n-y-t[1].second)
				}
			} else {
				ans = min(ans, n-y-t[0].second)
			}
		}
	}

	return ans
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

type Pair struct {
	first  int
	second int
}

func createPairs(m map[int]int) []Pair {
	res := make([]Pair, len(m))
	var id int
	for k, v := range m {
		res[id] = Pair{k, v}
		id++
	}
	return res
}
