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
		s, _ := reader.ReadBytes('\n')
		var n int
		pos := readInt(s, 0, &n) + 1
		var K uint64
		readUint64(s, pos, &K)
		A := readNNums(reader, n)
		res := solve(int64(K), A)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}

	fmt.Print(buf.String())
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

func solve(K int64, A []int) int {
	n := len(A)

	tot := int64(n) * int64(n+1) / 2

	K = tot - K + 1

	cnt := make([]int, n+1)

	count := func(m int) int64 {
		if m == 0 {
			return tot
		}
		for i := 0; i <= n; i++ {
			cnt[i] = 0
		}
		var mex int
		var r int
		for r < n && mex < m {
			cnt[A[r]]++
			for mex < m && cnt[mex] > 0 {
				mex++
			}
			r++
		}
		if mex < m {
			return 0
		}
		res := int64(n - r + 1)
		for l := 0; l < n; l++ {
			cnt[A[l]]--
			for r < n && cnt[A[l]] == 0 && A[l] < m {
				cnt[A[r]]++
				r++
			}
			if r == n && A[l] < m && cnt[A[l]] == 0 {
				break
			}
			res += int64(n - r + 1)
		}

		return res
	}

	l, r := 0, n+1

	for l < r {
		mid := (l + r) / 2
		cnt := count(mid)
		if cnt < K {
			// less mex, more count
			r = mid
		} else {
			l = mid + 1
		}
	}
	return r - 1
}
