package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var buf bytes.Buffer
	tc := readNum(reader)

	for tc > 0 {
		tc--

		y, x1, x2 := readThreeNums(reader)

		res := solve(x1, x2, y)

		buf.WriteString(fmt.Sprintf("%d\n", res))
	}

	fmt.Print(buf.String())
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

func solve(x1, x2, y int) int {
	y = abs(y)
	if y == 1 {
		if int64(x1)*int64(x2) <= 0 {
			return abs(x1) + abs(x2) + 1
		}
		return abs(x2-x1) + 1
	}

	if y == 0 {
		if x1 < 0 && x2 > 0 {
			return 3
		}
		if x1 == 0 && x2 > 0 || x1 < 0 && x2 == 0 {
			return 2
		}
		return 1
	}

	var factors []int

	for i := 2; i <= y/i; i++ {
		if y%i == 0 {
			factors = append(factors, i)
			for y%i == 0 {
				y /= i
			}
		}
	}

	if y > 1 {
		factors = append(factors, y)
	}
	if int64(x1)*int64(x2) < 0 {
		return countCoprime(factors, abs(x1)) + countCoprime(factors, abs(x2))
	}

	if x1 <= 0 && x2 <= 0 {
		x1, x2 = -x2, -x1
	}
	return countCoprime(factors, x2) - countCoprime(factors, max(0, x1-1))
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func countCoprime(fact []int, num int) int {
	n := len(fact)
	t := 1 << n

	var ans int

	for mask := 0; mask < t; mask++ {
		var cnt int
		val := 1
		for i := 0; i < n; i++ {
			if (mask>>i)&1 == 1 {
				val *= fact[i]
				cnt++
			}
		}

		if cnt&1 == 1 {
			ans -= num / val
		} else {
			ans += num / val
		}
	}
	return ans
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
