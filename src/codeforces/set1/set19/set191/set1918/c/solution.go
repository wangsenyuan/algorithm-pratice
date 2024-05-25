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
		a, b, r := readThreeNums(reader)
		res := solve(a, b, r)

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

const H = 63

func solve(a, b, r int) int {

	process := func(u, v int, i int, x int) int {
		// 希望u在满足条件的情况下取最大值，v在满足条件的情况取最小值
		var res int
		for j := i - 1; j >= 0; j-- {
			f := (u >> j) & 1
			s := (v >> j) & 1
			if f == s {
				// 无论怎么处理，这个地方都没有贡献
				// 放置0
				continue
			}
			if f == 1 {
				// 这里可以取0
				res += 1 << j
				continue
			}
			// f == 0
			if x|(1<<j) <= r {
				res += 1 << j
				x |= 1 << j
				continue
			}
			// 这里只能取0
			res -= 1 << j
		}

		return res
	}

	for i := H - 1; i >= 0; i-- {
		if (a>>i)&1 != (b>>i)&1 {
			res := 1 << i
			u, v := -res, -res

			if (b>>i)&1 == 1 || 1<<i <= r {
				var x int
				if (b>>i)&1 == 0 {
					x = 1 << i
				}
				// a这边获得最大值的情况下，希望 1 << i出现在b那边
				u = process(a, b, i, x)
			}

			if (a>>i)&1 == 1 || (1<<i) <= r {
				var x int
				if (a>>i)&1 == 0 {
					x = 1 << i
				}
				v = process(b, a, i, x)
			}

			tmp := max(u, v)
			return res - tmp
		}
	}

	return 0
}
