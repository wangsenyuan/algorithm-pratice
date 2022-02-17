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
		pos := readInt(s, 0, &n)
		var m uint64
		readUint64(s, pos+1, &m)
		A := readNNums(reader, n)
		res := solve(int64(m), A)

		buf.WriteString(fmt.Sprintf("%d\n", res))
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

const MAX_N = 100000

func solve(M int64, A []int) int {
	cnt := make([]int64, MAX_N+1)
	sum := make([]int64, MAX_N+1)

	for _, num := range A {
		cnt[num]++
		sum[num] += int64(num)
	}

	for i := 1; i <= MAX_N; i++ {
		cnt[i] += cnt[i-1]
		sum[i] += sum[i-1]
	}

	getSum := func(l, r int) int64 {
		return sum[r] - sum[l]
	}

	getCnt := func(l, r int) int64 {
		return cnt[r] - cnt[l]
	}

	for i := MAX_N; i > 1; i-- {
		var cur int64
		for j := 0; j < MAX_N; j += i {
			a := getSum(j, min(MAX_N, j+i-1))
			b := getCnt(j, min(MAX_N, j+i-1))
			cur += a - b*int64(j)
		}

		if M >= cur && (M-cur)%int64(i) == 0 {
			return i
		}
	}

	return 1
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
