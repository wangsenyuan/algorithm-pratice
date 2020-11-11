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
		s, _ := reader.ReadString('\n')
		cnt, arr := solve(n, s)
		buf.WriteString(fmt.Sprintf("%d\n", cnt))
		for i := 0; i < n; i++ {
			buf.WriteString(fmt.Sprintf("%d ", arr[i]))
		}
		buf.WriteByte('\n')
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

func solve(n int, s string) (int, []int) {
	set := make([]int, n)
	pos := make([][]int, 2)
	for i := 0; i < 2; i++ {
		pos[i] = make([]int, 0, n)
	}

	put := func(p int, i int, j int) int {
		ans := len(pos[i]) + len(pos[j])

		if len(pos[i]) == 0 {
			pos[j] = append(pos[j], ans)
		} else {
			n := len(pos[i])
			ans = pos[i][n-1]
			pos[i] = pos[i][:n-1]
			pos[j] = append(pos[j], ans)
		}

		return ans
	}

	for i := 0; i < n; i++ {
		if s[i] == '0' {
			set[i] = put(i, 1, 0) + 1
		} else {
			set[i] = put(i, 0, 1) + 1
		}
	}

	res := len(pos[0]) + len(pos[1])

	return res, set
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
