package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	var buf bytes.Buffer

	for tc > 0 {
		tc--
		n := readNum(reader)
		A := readNNums(reader, n)
		ok, res := solve(A)
		if !ok {
			buf.WriteString("-1\n")
		} else {
			buf.WriteString(fmt.Sprintf("%d\n", len(res)))
			for _, op := range res {
				buf.WriteString(fmt.Sprintf("%d %d %d\n", op[0], op[1], op[2]))
			}
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
		if s[i] == '\n' {
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

func solve(A []int) (bool, [][]int) {
	// A[i] = A[j] - A[k]
	// i < j < k
	// 最后的两个数是不能变化的
	n := len(A)

	if sort.IntsAreSorted(A) {
		return true, nil
	}
	// 如果 A[n-2] > A[n-1] => -1
	// A[n-2] <= A[n-1]
	// 如果 A[n-2] >= 0, 始终选择 n - 2 和 n- 1, 即可
	// A[n-2] < 0
	// 假设i满足条件了，也就是 A[i] <= A[i+1] <= A[i+2] ...<= A[n-1]
	// 现在 A[i-1] > A[i],需要使得 A[i-1] <= A[i], 如果A[i] 是某种操作得到的，那ok
	if A[n-2] > A[n-1] {
		return false, nil
	}
	// x = A[n-2] - A[n-1] <= A[n-2]
	// A[n-1] >= 0
	var res [][]int

	if A[n-1] >= 0 {
		for i := n - 3; i >= 0; i-- {
			res = append(res, []int{i + 1, n - 1, n})
		}
	} else {
		// can't use the previous way
		//  A[n-2] <= A[n-1] < 0
		// 假设 A[i] > 0 ( > A[i+1] of cause)
		// but no way to change A[i] < A[i+1]
		return false, nil
	}

	return true, res
}
