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

	for tc > 0 {
		tc--
		n := readNum(reader)

		ask := func(a []int, b []int) int {
			var buf bytes.Buffer
			buf.WriteString(fmt.Sprintf("%d %d", len(a), len(b)))
			for _, x := range a {
				buf.WriteString(fmt.Sprintf(" %d", x))
			}
			for _, x := range b {
				buf.WriteString(fmt.Sprintf(" %d", x))
			}
			buf.WriteByte('\n')
			fmt.Print(buf.String())
			return readNum(reader)
		}

		res := solve(n, ask)
		fmt.Printf("-1 %d\n", res)
	}
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

func solve(n int, ask func([]int, []int) int) int {
	var res int

	for c := 0; c < 9; c++ {
		// 最远的两个点，它们的编号，必然是会相差的
		var a []int
		var b []int
		for i := 0; i < n; i++ {
			if (i>>c)&1 == 0 {
				a = append(a, i+1)
			} else {
				b = append(b, i+1)
			}
		}
		if len(a) == 0 || len(b) == 0 {
			continue
		}
		res = max(res, ask(a, b))
	}

	return res
}
