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
		n := readNum(reader)
		A := readNNums(reader, n)
		cnt, res := solve(A)
		buf.WriteString(fmt.Sprintf("%d\n", cnt))
		if cnt > 0 {
			for _, step := range res {
				buf.WriteString(fmt.Sprintf("%d %d\n", step[0], step[1]))
			}
		}
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

func solve(A []int) (int, [][]int) {
	a := -1

	n := len(A)

	for i := 0; i < n; i++ {
		if A[i]&1 == 1 {
			a = i
		}
	}
	if a < 0 {
		return -1, nil
	}

	cnt1, res1 := change(0, A, a)
	cnt2, res2 := change(1, A, a)
	if cnt1 < cnt2 {
		return cnt1, res1
	}
	return cnt2, res2
}

func change(first int, B []int, odd int) (int, [][]int) {
	n := len(B)
	A := make([]int, n)
	copy(A, B)
	// try find index that is same pairty with first, and at the even positions
	a := -1
	for i := 0; i < n; i++ {
		A[i] &= 1
		if A[i] == 1 {
			if i&1 == 0 && first == 1 {
				//10pattern
				a = i
			} else if i&1 == 1 && first == 0 {
				//01pattern
				a = i
			}
		}
	}
	var res [][]int
	if a < 0 {
		res = append(res, []int{1, odd + 1})
		A[0] ^= A[odd]
		a = 0
	}

	for i := 0; i < n; i++ {
		if i&1 == 0 {
			if A[i] != first {
				res = append(res, []int{i + 1, a + 1})
			}
		} else {
			if A[i] == first {
				res = append(res, []int{i + 1, a + 1})
			}
		}
	}
	return len(res), res
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
