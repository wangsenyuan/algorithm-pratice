package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var buf bytes.Buffer

	tc := readNum(reader)

	for tc > 0 {
		tc--
		readNum(reader)
		s := readString(reader)
		s = strings.ReplaceAll(s, " ", "")
		_, res := solve(s)
		buf.WriteString(fmt.Sprintf("%d\n", len(res)))
		for _, swap := range res {
			buf.WriteString(fmt.Sprintf("%d %d\n", swap[0], swap[1]))
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

func solve(s string) (bool, [][]int) {
	n := len(s)

	var swaps [][]int

	A := make([]int, n+1)

	for i := 1; i <= n; i++ {
		A[i] = int(s[i-1] - '0')
	}

	swap := func(i, j int) {
		swaps = append(swaps, []int{i, j})
		A[i], A[j] = A[j], A[i]
		A[i+1], A[j+1] = A[j+1], A[i+1]
	}

	cnt := make([]int, 2)

	for i := 1; i <= n; i++ {
		cnt[A[i]]++
	}

	if max(cnt[0], cnt[1]) == n {
		return true, nil
	}

	placeZeros := func() {
		keep := 0
		for i := 2; i <= n; i++ {
			if A[i] == 1 || A[i-1] == 0 {
				continue
			}
			if A[i-1] == 1 {
				for A[keep] == 0 {
					keep++
				}
				if i == n {
					swap(n-3, n-1)
					swap(keep, n-2)
					continue
				}
				if keep+1 < i {
					swap(keep, i)
				} else {
					swap(i, i+2)
					swap(keep, i+2)
				}
			}
		}
	}

	if cnt[0] <= cnt[1] {
		placeZeros()
	} else {
		B := make([]int, n+1)
		for i := 1; i <= n; i++ {
			B[i] = 1 - A[n+1-i]
		}
		copy(A, B)
		placeZeros()

		for i := 0; i < len(swaps); i++ {
			a, b := swaps[i][0], swaps[i][1]
			swaps[i] = []int{n - b, n - a}
		}
	}
	return true, swaps
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
