package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, x := readTwoNums(reader)
	a := readNNums(reader, n)

	res := solve(a, x)

	fmt.Println(res)
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

func solve(a []int, t int) int {
	n := len(a)
	// L是升序的
	L := make([]int, n)
	// R是降序排列的
	R := make([]int, n)

	help := make([]int, n)

	var res int
	merge := func(l int, mid int, r int) {
		// 计算中间的部分
		for i, j := l, mid+1; i <= mid; i++ {
			x := R[i]
			for j <= r && x+L[j] < t {
				j++
			}
			res += j - (mid + 1)
		}

		var sum int
		for i := l; i <= mid; i++ {
			sum += a[i]
		}
		for i, j, k := l, mid+1, l; k <= r; k++ {
			if j > r || i <= mid && L[i] <= sum+L[j] {
				help[k] = L[i]
				i++
			} else {
				help[k] = sum + L[j]
				j++
			}
		}

		sum = 0
		for i := mid + 1; i <= r; i++ {
			sum += a[i]
		}

		copy(L[l:r+1], help[l:r+1])
		for i, j, k := mid, r, r; k >= l; k-- {
			if j == mid || i >= l && sum+R[i] <= R[j] {
				help[k] = sum + R[i]
				i--
			} else {
				help[k] = R[j]
				j--
			}
		}
		copy(R[l:r+1], help[l:r+1])
	}

	var dfs func(l int, r int)

	dfs = func(l int, r int) {
		if l == r {
			if a[l] < t {
				res++
			}
			L[l] = a[l]
			R[l] = a[l]
			return
		}
		mid := (l + r) / 2
		dfs(l, mid)
		dfs(mid+1, r)
		merge(l, mid, r)
	}

	dfs(0, n-1)

	return res
}
