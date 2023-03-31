package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	ask := func(i, j, k int) int {
		fmt.Printf("? %d %d %d\n", i, j, k)
		return readNum(reader)
	}

	for tc > 0 {
		tc--
		n := readNum(reader)
		res := solve(n, ask)
		fmt.Printf("! %d %d\n", res[0], res[1])
	}
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

func max(a, b, c int) int {
	if a >= b && a >= c {
		return a
	}
	if b >= c && b >= a {
		return b
	}
	return c
}

func min(a, b, c int) int {
	if a <= b && a <= c {
		return a
	}
	if b <= a && b <= c {
		return b
	}

	return c
}

func solve(n int, ask func(int, int, int) int) []int {
	// ask(i, j, k) = max(A[i], A[j], A[k]) - min(A[i], A[j], A[k])
	// 如果 i, j, k中有0
	// ask(i, j, k) = max(A[i], A[j], A[k])
	// 先确定两个下标 a, b A[a] > or < A[b]
	// 4个数，里面必然有一个最大值，一个最小值, 且最大值 - 最小值至少出现2次
	// （0， 1， 2， 3） 有3个排列，那么就可以知道 A[i] & A[j] 是它们中间的最大值，或者最小值
	// 然后，开始和任意的i进行check
	// 如果差变大了，那么 i是新的最大值，或者最小值
	cache := make(map[int]int)
	query := func(i, j, k int) int {
		tmp := make([]int, 3)
		tmp[0] = max(i, j, k)
		tmp[1] = min(i, j, k)
		tmp[2] = i + j + k - tmp[0] - tmp[1]
		key := tmp[0]*1000000 + tmp[1]*1000 + tmp[2]
		if v, ok := cache[key]; ok {
			return v
		}
		cache[key] = ask(i+1, j+1, k+1)
		return cache[key]
	}

	find := func(a, b, c, d int) []int {
		X := [][]int{
			{a, b, c},
			{a, c, d},
			{a, b, d},
			{b, c, d},
		}
		ans := make(map[int][]int)
		var mx int
		for i, x := range X {
			sort.Ints(x)
			y := query(x[0], x[1], x[2])
			if len(ans[y]) == 0 {
				ans[y] = make([]int, 0, 1)
			}
			ans[y] = append(ans[y], i)
			if mx < y {
				mx = y
			}
		}

		for y, vs := range ans {
			if y == mx && len(vs) >= 2 {
				ii := X[vs[0]]
				jj := X[vs[1]]

				var tmp []int

				for i, j := 0, 0; i < 3 && j < 3; {
					if ii[i] == jj[j] {
						tmp = append(tmp, ii[i])
						i++
						j++
					} else if ii[i] < jj[j] {
						i++
					} else {
						j++
					}
				}
				return tmp
			}
		}

		return nil
	}
	// 4个一组，4次查询后， 可以找到 每组中的最大和最小值
	// 然后再4个一组， 相同于每次减少一半
	// 1024 + 512 + 256 + <= 2 * 1000 - 2
	// 如果呢不能整除4， 那么剩下1， 2， 3个

	arr := make([]int, n)

	for i := 0; i < n; i++ {
		arr[i] = i
	}

	for len(arr) >= 4 {
		var cur []int
		var p int
		for p+4 <= len(arr) {
			tmp := find(arr[p], arr[p+1], arr[p+2], arr[p+3])
			cur = append(cur, tmp...)
			p += 4
		}
		for p < len(arr) {
			cur = append(cur, arr[p])
			p++
		}
		arr = cur
	}

	if len(arr) == 3 {
		for i := 0; i < n; i++ {
			if i != arr[0] && i != arr[1] && i != arr[2] {
				arr = find(i, arr[0], arr[1], arr[2])
				break
			}
		}
	}

	arr[0]++
	arr[1]++

	return arr
}
