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
		n := readNum(reader)
		A := readNNums(reader, n)
		res := solve(A)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}

	fmt.Print(buf.String())
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

func solve(A []int) int {
	n := len(A)
	const H = 30
	first := make([]int, H)
	last := make([]int, H)
	cnt := make([][]int, H)
	var or int
	for i := 0; i < H; i++ {
		first[i] = -1
		last[i] = -1
		cnt[i] = make([]int, n+1)

		for j := 0; j < n; j++ {
			cnt[i][j+1] = cnt[i][j]
			if (A[j]>>i)&1 == 1 {
				cnt[i][j+1]++
				last[i] = j
				if first[i] == -1 {
					first[i] = j
				}
			}
		}
		if last[i] >= 0 {
			or |= (1 << i)
		}
	}
	ans := -1
	for l := 0; l < n; l++ {
		r := n - 1
		for j := 0; j < H; j++ {
			// [l..r] |= or
			if first[j] < l {
				continue
			}
			r = min(r, last[j]-1)
		}

		if l <= r {
			ok := true
			for j := 0; j < H && ok; j++ {
				if last[j] < 0 {
					continue
				}
				ok = cnt[j][r+1]-cnt[j][l] > 0
			}
			if ok && (ans < 0 || ans < r-l+1) {
				ans = r - l + 1
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

func solve1(A []int) int {
	n := len(A)

	arr := make([]int, 2*n)
	copy(arr[n:], A)

	for i := n - 1; i > 0; i-- {
		arr[i] = arr[i*2] | arr[i*2+1]
	}

	get := func(l, r int) int {
		l += n
		r += n
		var res int
		for l < r {
			if l&1 == 1 {
				res |= arr[l]
				l++
			}
			if r&1 == 1 {
				r--
				res |= arr[r]
			}
			l >>= 1
			r >>= 1
		}
		return res
	}

	or := get(0, n)

	ans := -1

	var cur int
	for r := 0; r < n; r++ {
		cur |= A[r]
		if cur == or {
			l := search(0, r+1, func(i int) bool {
				return get(i, r+1) < or
			})
			l--
			// get(l, r + 1) == or
			var x int
			if l > 0 {
				x = get(0, l)
			}
			var y int
			if r+1 < n {
				y = get(r+1, n)
			}

			if (x | y) == or {
				// a candidate
				ll := search(0, l, func(i int) bool {
					return get(0, i)|y == or
				})
				if ans < 0 || ans < (r-ll+1) {
					ans = r - ll + 1
				}
			}
		}
	}

	return ans
}

func search(l int, r int, f func(int) bool) int {
	for l < r {
		mid := (l + r) / 2
		if f(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return r
}
