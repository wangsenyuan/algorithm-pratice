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
		n, k := readTwoNums(reader)
		res := solve(n, k)
		buf.WriteString(fmt.Sprintf("%d\n", res))
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

func solve(n int, k int) int {
	f := func(x int) int {
		var res int
		for x > 0 {
			res += x % 10
			x /= 10
		}
		return res
	}

	s := func(x int) int {
		var res int
		for i := 0; i <= k; i++ {
			res += f(x + i)
		}
		return res
	}

	for x := 0; x < 100; x++ {
		if s(x) == n {
			return x
		}
	}

	g := func(w int, c int) int {
		// w 是个位数，c是个位数前面（10位开始，有多少个9）
		var sum int
		var i int
		for w+i < 10 && i <= k {
			sum += w + i
			sum += c * 9
			i++
		}
		for i <= k {
			sum += (w + i) % 10
			// 产生了c次进位，并增加了1
			sum += 1
			i++
		}
		if sum > n {
			return -1
		}
		sum = n - sum
		if sum%(k+1) != 0 {
			return -1
		}
		sum /= (k + 1)
		// 最高位只能是8
		if sum <= 8 {
			x := sum
			for c > 0 {
				x = x*10 + 9
				c--
			}
			x = x*10 + w
			return x
		}
		// 最好有个8
		sum -= 8
		y, x := sum/9, sum%9
		for y > 0 {
			x = x*10 + 9
			y--
		}
		x = x*10 + 8
		for c > 0 {
			x = x*10 + 9
			c--
		}
		x = x*10 + w

		return x
	}

	res := -1

	for w := 0; w < 10; w++ {
		for c := 0; c*9+w <= n; c++ {
			tmp := g(w, c)
			if tmp < 0 {
				continue
			}
			if res < 0 || tmp < res {
				res = tmp
			}
		}
	}

	return res
}
