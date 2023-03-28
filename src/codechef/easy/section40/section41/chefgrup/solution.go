package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n := readNum(reader)
	nums := readNNums(reader, n)
	res := solve(nums)
	fmt.Println(res)
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

func solve(A []int) int {
	// 如果 gcd(a, b) > 1, 那么它们必须在一个group内
	// 然后看这样的group有多少个
	// union-find
	n := len(A)
	set := make([]int, n)
	cnt := make([]int, n)
	for i := 0; i < n; i++ {
		set[i] = i
		cnt[i]++
	}

	var find func(u int) int
	var union func(a, b int) bool

	find = func(u int) int {
		if set[u] != u {
			set[u] = find(set[u])
		}
		return set[u]
	}

	union = func(a, b int) bool {
		a = find(a)
		b = find(b)
		if a == b {
			return false
		}
		if cnt[a] < cnt[b] {
			a, b = b, a
		}
		cnt[a] += cnt[b]
		set[b] = a
		return true
	}

	mx := A[0]

	for i := 1; i < n; i++ {
		mx = max(mx, A[i])
	}
	pf := make([]int, mx+1)
	for i := 2; i <= mx; i++ {
		if pf[i] == 0 {
			pf[i] = i
			if mx/i < i {
				continue
			}
			for j := i * i; j <= mx; j += i {
				if pf[j] == 0 {
					pf[j] = i
				}
			}
		}
	}

	prev := make([]int, mx+1)

	for i, num := range A {
		tmp := num
		for tmp > 1 {
			x := pf[tmp]
			if prev[x] > 0 {
				union(prev[x]-1, i)
			}
			tmp /= x
		}
		tmp = num
		for tmp > 1 {
			x := pf[tmp]
			prev[x] = i + 1
			tmp /= x
		}
	}
	gids := make([]int, n)
	for i := 0; i < n; i++ {
		gids[find(i)]++
	}
	var res int

	for i := 0; i < n; i++ {
		if gids[i] > 0 {
			res++
		}
	}

	return res
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
