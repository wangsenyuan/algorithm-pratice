package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := 1

	var buf bytes.Buffer

	for tc > 0 {
		tc--
		n, m := readTwoNums(reader)
		Q := make([][]int, m)
		for i := 0; i < m; i++ {
			s, _ := reader.ReadBytes('\n')
			if s[0] == '1' {
				Q[i] = make([]int, 3)
				var l, r int
				pos := readInt(s, 2, &l)
				readInt(s, pos+1, &r)
				Q[i] = []int{1, l, r}
			} else {
				var x int
				readInt(s, 2, &x)
				Q[i] = []int{2, x}
			}
		}
		res := solve(n, Q)

		for _, ans := range res {
			buf.WriteString(fmt.Sprintf("%d\n", ans))
		}
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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
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

func solve(n int, Q [][]int) []int64 {
	sz := n
	arr := make([]int, 16)
	if n > 15 {
		sz = 14
		for i := n - 13; i <= n; i++ {
			arr[i-(n-13)+1] = i
		}
	} else {
		for i := 1; i <= sz; i++ {
			arr[i] = i
		}
	}

	fact := make([]int64, 16)
	fact[0] = 1
	for i := 1; i <= 15; i++ {
		fact[i] = int64(i) * fact[i-1]
	}

	tmp := make([]int, 16)
	vis := make([]bool, 16)
	var change func(id int, x int64)

	change = func(id int, x int64) {
		if id > sz {
			return
		}
		var cnt int64
		lf := fact[sz-id]
		val := x / lf
		for i := 1; i <= sz; i++ {
			if !vis[i] {
				if val == cnt {
					tmp[id] = arr[i]
					vis[i] = true
					change(id+1, x%lf)
					break
				}
				cnt++
			}
		}
	}

	calc := func(x int64) int64 {
		return x * (x + 1) / 2
	}

	var offset int64

	change(1, offset)

	var ans []int64

	for _, cur := range Q {
		if cur[0] == 1 {
			l, r := cur[1], cur[2]
			var res int64
			if r >= arr[1] {
				for j := max(arr[1], l); j <= r; j++ {
					res += int64(tmp[j-arr[1]+1])
				}
				r = arr[1] - 1
			}
			if l <= r {
				res += calc(int64(r)) - calc(int64(l)-1)
			}
			ans = append(ans, res)
		} else {
			offset += int64(cur[1])
			for i := 0; i < len(vis); i++ {
				vis[i] = false
			}
			change(1, offset)
		}
	}

	return ans
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
