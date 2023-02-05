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
		Q := make([][]int, 1)
		Q[0] = readNNums(reader, 2)
		res := solve(A, Q)
		for _, ans := range res {
			buf.WriteString(fmt.Sprintf("%d %d\n", ans[0], ans[1]))
		}
	}

	fmt.Print(buf.String())
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

func solve(A []int, Q [][]int) [][]int {
	n := len(A)
	sum := make([]int64, n+1)
	xor := make([]int64, n+1)

	for i, x := range A {
		sum[i+1] = sum[i] + int64(x)
		xor[i+1] = xor[i] ^ int64(x)
	}

	find := func(q []int) []int {
		l, r := q[0], q[1]
		l--
		exp := (sum[r] - sum[l]) - (xor[r] ^ xor[l])
		ans := []int{l + 1, r}

		for i, j := l, l; j < r; j++ {
			tmp := (sum[j+1] - sum[i]) - (xor[j+1] ^ xor[i])
			if tmp < exp {
				continue
			}
			for i < j && (sum[j+1]-sum[i+1])-(xor[j+1]^xor[i+1]) == exp {
				i++
			}
			// tmp == exp
			if ans[1]-ans[0]+1 > j-i+1 {
				ans = []int{i + 1, j + 1}
			}
		}
		//for i := l; i < r; i++ {
		//	j := search(i, n, func(j int) bool {
		//		return sum[j]-sum[i]-(xor[j]^xor[i]) >= exp
		//	})
		//	if j > r || sum[j]-sum[i]-(xor[j]^xor[i]) != exp {
		//		continue
		//	}
		//	if ans[0] < 0 || ans[1]-ans[0] > j-(i+1) {
		//		ans = []int{i + 1, j}
		//	}
		//}

		return ans
	}

	ans := make([][]int, len(Q))

	for i, cur := range Q {
		ans[i] = find(cur)
	}

	return ans
}

func search(l, r int, f func(int) bool) int {
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
