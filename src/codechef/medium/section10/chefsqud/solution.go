package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

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
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') {
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

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	for tc > 0 {
		tc--
		n := readNum(reader)
		A := readNNums(reader, n)
		fmt.Println(solve(n, A))
	}
}

func solve(n int, A []int) int64 {
	B := make([]int, n)
	copy(B, A)

	sort.Ints(B)

	ii := make(map[int]int)
	var p int = 1
	for i := 1; i < n; i++ {
		if B[i] == B[i-1] {
			continue
		}
		ii[B[i]] = p
		p++
	}

	arr := make([]int64, p+1)

	update := func(pos int, v int64) {
		pos++
		for pos <= p {
			arr[pos] += v
			pos += pos & -pos
		}
	}

	get := func(pos int) int64 {
		var res int64
		pos++
		for pos > 0 {
			res += arr[pos]
			pos -= pos & -pos
		}

		return res
	}

	getRange := func(left, right int) int64 {
		res := get(right)
		if left > 0 {
			res -= get(left - 1)
		}
		return res
	}

	reset := func() {
		for i := 0; i < len(arr); i++ {
			arr[i] = 0
		}
	}

	count := func(mid int64) int64 {
		reset()
		var cur, cnt int64
		var l int
		var r int = -1

		for l < n && r < n {
			if r < n-1 {
				pos := ii[A[r+1]]
				next := getRange(pos+1, p-1)

				for cur+next <= mid && r < n-1 {
					r++
					cur += next
					update(pos, 1)
					if r < n-1 {
						pos = ii[A[r+1]]
						next = getRange(pos+1, p-1)
					}
				}
			}

			cnt += int64(r - l + 1)
			pos := ii[A[l]]
			l++
			cur -= get(pos - 1)
			update(pos, -1)
		}

		return cnt
	}
	N := int64(n)
	var l int64
	var r int64 = 1 + N*(N-1)/2
	medianPos := (1 + N*(N+1)/2) / 2
	var ans int64
	for l <= r {
		mid := (l + r) / 2
		cnt := count(mid)
		if cnt >= medianPos {
			ans = mid
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return ans
}
