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

func fillNNums(bs []byte, n int, res []int) {
	x := 0
	for i := 0; i < n; i++ {
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, m := readTwoNums(reader)

	A := readNNums(reader, n)

	queries := make([][]int, m)

	for i := 0; i < m; i++ {
		queries[i] = readNNums(reader, 2)
	}

	res := solve(A, queries)

	for i := 0; i < m; i++ {
		fmt.Println(res[i])
	}
}

func solve(A []int, quries [][]int) []int64 {
	n := len(A)
	B := make([]int, n)
	copy(B, A)

	sort.Ints(B)

	var m int

	for i := 1; i <= len(B); i++ {
		if i == len(B) || B[i] > B[i-1] {
			B[m] = B[i-1]
			m++
		}
	}

	B = B[:m]
	for i := 0; i < n; i++ {
		// change A[i] to its index
		A[i] = sort.SearchInts(B, A[i]) + 1
	}

	ans := make([]int64, len(quries))

	prefix := make([][]Pair, n+2)
	suffix := make([][]Pair, n+2)
	for i := 0; i <= n+1; i++ {
		prefix[i] = make([]Pair, 0, 10)
		suffix[i] = make([]Pair, 0, 10)
	}

	for i := 1; i <= len(quries); i++ {
		l, r := quries[i-1][0], quries[i-1][1]
		if l > r {
			l, r = r, l
		}
		if A[l-1] != A[r-1] {
			ans[i-1]++
		}

		prefix[l-1] = append(prefix[l-1], Pair{A[l-1] + 1, -i})
		suffix[l+1] = append(suffix[l+1], Pair{A[l-1] - 1, -i})
		prefix[r-1] = append(prefix[r-1], Pair{A[r-1] + 1, -i})
		suffix[r+1] = append(suffix[r+1], Pair{A[r-1] - 1, -i})

		prefix[l-1] = append(prefix[l-1], Pair{A[r-1] + 1, i})
		suffix[l+1] = append(suffix[l+1], Pair{A[r-1] - 1, i})
		prefix[r-1] = append(prefix[r-1], Pair{A[l-1] + 1, i})
		suffix[r+1] = append(suffix[r+1], Pair{A[l-1] - 1, i})
	}

	fwt := make([]int, m+1)
	var add int64

	for i := 1; i <= n; i++ {
		add += int64(i - 1)
		for j := A[i-1]; j > 0; j = j & (j - 1) {
			add -= int64(fwt[j])
		}

		for j := A[i-1]; j <= m; j = (j | (j - 1)) + 1 {
			fwt[j]++
		}
	}

	for i := 0; i <= m; i++ {
		fwt[i] = 0
	}

	for i := 1; i <= n; i++ {
		for j := A[i-1]; j > 0; j = j & (j - 1) {
			fwt[j]++
		}

		for q := 0; q < len(prefix[i]); q++ {
			idx := abs(prefix[i][q].second)
			mu := prefix[i][q].second / idx

			for j := prefix[i][q].first; j <= m; j = (j | (j - 1)) + 1 {
				ans[idx-1] += int64(mu * fwt[j])
			}
		}
	}

	for i := 0; i <= m; i++ {
		fwt[i] = 0
	}

	for i := n; i > 0; i-- {
		for j := A[i-1]; j <= m; j = (j | (j - 1)) + 1 {
			fwt[j]++
		}

		for q := 0; q < len(suffix[i]); q++ {
			idx := abs(suffix[i][q].second)
			mu := suffix[i][q].second / idx

			for j := suffix[i][q].first; j > 0; j = j & (j - 1) {
				ans[idx-1] += int64(mu * fwt[j])
			}
		}
	}

	for i := 0; i < len(quries); i++ {
		ans[i] += add
	}

	return ans
}

type Pair struct {
	first, second int
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
