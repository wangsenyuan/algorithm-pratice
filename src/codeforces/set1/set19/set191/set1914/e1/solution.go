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
		a := readNNums(reader, n)
		b := readNNums(reader, n)
		res := solve(a, b)
		buf.WriteString(fmt.Sprintf("%d\n", res))
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

const inf = 1 << 60

func solve(a []int, b []int) int {
	// 对于alice和bob来说，策略似乎是一样的，就是尽量的摧毁对方最多的那个
	// 且要保证自己的score是更大的
	// 假设i，j a[i] > 0, b[i] > 0
	//	摧毁i后， 对于alice来说, a[i-1] - b[i]
	//  摧毁j后， a[j-1] - b[j]

	n := len(a)
	N := 1 << n

	var play func(turn int, flag int) int

	play = func(turn int, flag int) int {
		if flag == N-1 {
			return 0
		}
		if turn == 0 {
			// alice
			var res int = -inf
			for i := 0; i < n; i++ {
				if (flag>>i)&1 == 0 {
					tmp := a[i] - 1
					tmp += play(1, flag|(1<<i))
					res = max(res, tmp)
				}
			}
			return res
		}
		// bob
		var res int = inf
		for i := 0; i < n; i++ {
			if (flag>>i)&1 == 0 {
				tmp := -(b[i] - 1)
				tmp += play(0, flag|(1<<i))
				res = min(res, tmp)
			}
		}
		return res
	}

	return play(0, 0)
}
