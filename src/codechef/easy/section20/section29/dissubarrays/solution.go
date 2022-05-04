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
		ok, B, C := solve(A)

		if !ok {
			buf.WriteString("NO\n")
			continue
		}
		buf.WriteString("YES\n")

		for i := 0; i < n; i++ {
			buf.WriteString(fmt.Sprintf("%d ", B[i]))
		}
		buf.WriteByte('\n')

		for i := 0; i < n; i++ {
			buf.WriteString(fmt.Sprintf("%d ", C[i]))
		}
		buf.WriteByte('\n')
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

func solve(A []int) (bool, []int, []int) {
	cnt := make(map[int]int)
	for _, num := range A {
		cnt[num]++
	}
	n := len(A)

	for _, v := range cnt {
		if v > n/2 {
			return false, nil, nil
		}
	}

	if n%2 == 0 && len(cnt) == 2 {
		// both half
		return false, nil, nil
	}

	sort.Ints(A)

	k := n / 2

	C := make([]int, n)
	copy(C, A[k:])
	copy(C[(n-k):], A)
	return true, A, C
}
