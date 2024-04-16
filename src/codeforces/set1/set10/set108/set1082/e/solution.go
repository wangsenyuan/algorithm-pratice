package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, c := readTwoNums(reader)

	a := readNNums(reader, n)

	res := solve(c, a)

	fmt.Println(res)
}

func readString(r *bufio.Reader) string {
	s, _ := r.ReadBytes('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return string(s[:i])
		}
	}
	return string(s)
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

func solve(c int, nums []int) int {
	n := len(nums)
	pref := make([]int, n+1)
	var x int
	for i, num := range nums {
		pref[i+1] = pref[i]
		if num == c {
			pref[i+1]++
		}
		x = max(x, num)
	}

	get := func(l int, r int) int {
		return pref[r] - pref[l]
	}
	last := make([]int, x+1)
	segs := make([][]int, x+1)
	for i := 0; i <= x; i++ {
		segs[i] = make([]int, 0, 1)
		last[i] = -1
	}

	for i, num := range nums {
		segs[num] = append(segs[num], -get(last[num]+1, i))
		segs[num] = append(segs[num], 1)
		last[num] = i
	}

	for i := 1; i <= x; i++ {
		segs[i] = append(segs[i], -get(last[i]+1, n))
	}

	calc := func(x int) int {
		res := -(1 << 60)
		var bal int
		for _, v := range segs[x] {
			bal = max(0, bal+v)
			res = max(res, bal)
		}

		return res
	}

	var best int
	for i := 1; i <= x; i++ {
		if i == c {
			continue
		}
		best = max(best, calc(i))
	}

	return best + get(0, n)
}

func solve1(c int, nums []int) int {
	n := len(nums)

	suf := make([]int, n+1)
	for i := n - 1; i >= 0; i-- {
		suf[i] = suf[i+1]
		if nums[i] == c {
			suf[i]++
		}
	}

	pos := make(map[int][]int)

	pref := make([]int, n+1)
	for i, num := range nums {
		pref[i+1] = pref[i]
		if nums[i] == c {
			pref[i+1]++
		}
		if v, ok := pos[num]; ok {
			pos[num] = append(v, i)
		} else {
			pos[num] = []int{i}
		}
	}

	count := func(vs []int, l int, r int) int {
		// 在l和r中间有多少个元素
		return r - l + 1 + pref[vs[l]] + suf[vs[r]+1]
	}

	L := make(map[int]int)
	R := make(map[int]int)

	diff := make([]int, n+1)

	find := func(r int) int {
		vs := pos[nums[r]]

		i, j := L[nums[r]], R[nums[r]]
		for vs[j] < r {
			j++
		}

		if diff[r+1]-diff[vs[i]] >= j-i {
			i = j
		}

		L[nums[r]] = i
		R[nums[r]] = j

		return count(vs, i, j)
	}

	best := suf[0]

	for r := 0; r < n; r++ {
		diff[r+1] = diff[r]
		if nums[r] == c {
			diff[r+1]++
			continue
		}
		f := find(r)
		best = max(best, f)
	}

	return best
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func search(n int, f func(int) bool) int {
	l, r := 0, n
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
