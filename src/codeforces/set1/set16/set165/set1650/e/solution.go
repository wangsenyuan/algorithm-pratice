package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var buf bytes.Buffer

	tc := readNum(reader)

	for tc > 0 {
		tc--
		readString(reader)
		n, d := readTwoNums(reader)
		A := readNNums(reader, n)
		res := solve(d, A)
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

func normalize(s string) string {

	for i := len(s); i > 0; i-- {
		if s[i-1] >= 'a' && s[i-1] <= 'z' {
			return s[:i]
		}
	}
	return ""
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

func solve(d int, A []int) int {
	n := len(A)
	arr := make([]int, n+1)
	copy(arr[1:], A)
	n++
	// when change A[i] to some place,
	// 1. 需要把造成最小u的A[i]移动
	// 2. 如果至少有两对  i, j, k, l, i < j < k < l， 则无法获得更好的结果
	// 3. 如果存在 i, j, k, A[j] - A[i] = u, A[k] - A[j] = u
	//   那么只能移动 A[j] 到 [0...i) 或者 （k, n)区间，获得比u更好的结果
	// 4. 如果存在 i, j, A[j] - A[i] = u, 那么可以有两个选择，i或者j
	u := arr[1] - arr[0]

	for i := 2; i < n; i++ {
		u = min(u, arr[i]-arr[i-1])
	}

	var idx []int
	// 1 not included
	for i := 1; i < n; i++ {
		if arr[i]-arr[i-1] == u {
			idx = append(idx, i)
		}
		if len(idx) > 2 {
			return u - 1
		}
	}

	check := func(id int, u int) bool {
		var ok bool
		for i := 1; i < n; i++ {
			// move id to get a better result >= u
			if i == id {
				// but only can't at position id
				continue
			}
			j := i - 1
			if j == id {
				j--
			}

			dist := arr[i] - arr[j]
			if dist < u {
				return false
			}
			if dist/2 >= u {
				ok = true
			}
		}

		if ok {
			return true
		}

		// not ok yet
		dist := d - arr[n-1]

		if id == n-1 {
			dist = d - arr[n-2]
		}
		// put it a d
		return dist >= u
	}

	move := func(id int) int {
		// place one between [l...r] to get the best distance
		l, r := 1, d+1
		for l < r {
			mid := (l + r) / 2
			if !check(id, mid) {
				r = mid
			} else {
				l = mid + 1
			}
		}
		return r - 1
	}

	if len(idx) == 2 {
		if idx[1]-1 != idx[0] {
			return u - 1
		}
		// j-1, j, j + 1
		// have to move j
		j := idx[0]
		u = max(u, move(j))
	} else {
		// len(idx) == 1
		j := idx[0]
		i := j - 1
		if i > 0 {
			u = max(u, move(i))
		}
		u = max(u, move(j))
	}

	return u - 1
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
