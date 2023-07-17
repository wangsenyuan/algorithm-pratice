package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
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

func main() {
	reader := bufio.NewReader(os.Stdin)

	//tc := readNum(scanner)
	tc := 1

	var buf bytes.Buffer

	for tc > 0 {
		tc--
		n := readNum(reader)
		stu := make([][]int, 2)
		for i := 0; i < 2; i++ {
			stu[i] = readNNums(reader, n)
		}

		res := solve(stu)

		buf.WriteString(fmt.Sprintf("%d\n", res))
	}

	fmt.Println(buf.String())
}

func solve(students [][]int) int64 {
	ans := make([]int64, 3)
	// ans[0] for the -2 column
	// ans[1] for -1 with 0 selected
	// ans[2] for -1 with 1 selected
	n := len(students[0])
	cur := make([]int64, 3)
	for i := 0; i < n; i++ {
		copy(cur, ans)
		ans[0] = max(cur...)
		ans[1] = max(cur[0], cur[2]) + int64(students[0][i])
		ans[2] = max(cur[0], cur[1]) + int64(students[1][i])
	}
	return max(ans...)
}

func max(arr ...int64) int64 {
	res := arr[0]
	for i := 1; i < len(arr); i++ {
		if res < arr[i] {
			res = arr[i]
		}
	}
	return res
}
