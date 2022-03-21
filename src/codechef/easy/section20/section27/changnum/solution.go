package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)
	var buf bytes.Buffer
	for tc > 0 {
		tc--
		n, k := readTwoNums(reader)
		A := readNNums(reader, n)
		B := readNNums(reader, n)
		ans1, ans2 := solve(A, B)
		if k == 1 {
			buf.WriteString(fmt.Sprintf("%d\n", ans1))
		} else {
			buf.WriteString(fmt.Sprintf("%d %d\n", ans1, ans2))
		}
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
		if s[i] == '\n' {
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

const MOD = 1000000007
const N = 100010

var fact []int64

func init() {
	fact = make([]int64, N)
	fact[0] = 1
	for i := 1; i < N; i++ {
		fact[i] = int64(i) * fact[i-1] % MOD
	}
}

func solve(A, B []int) (int, int) {
	// at most n steps will work
	n := len(A)

	set := make([]int, N)
	cnt := make([]int, N)

	for i := 1; i <= n; i++ {
		set[i] = i
		cnt[i]++
	}

	var find func(x int) int
	find = func(x int) int {
		if set[x] != x {
			set[x] = find(set[x])
		}
		return set[x]
	}

	union := func(x, y int) {
		px := find(x)
		py := find(y)
		if px != py {
			if cnt[px] < cnt[py] {
				px, py = py, px
			}
			set[py] = px
			cnt[px] += cnt[py]
		}
	}

	for i := 0; i < n; i++ {
		if A[i] != B[i] {
			union(A[i], B[i])
		}
	}

	var ans1 int
	var ans2 int64 = 1

	for i := 1; i <= n; i++ {
		if set[i] == i {
			ans1 += cnt[i] - 1
			ans2 *= fact[cnt[i]]
			ans2 %= MOD
		}
	}

	ans2 *= fact[ans1]
	ans2 %= MOD

	return ans1, int(ans2)
}
