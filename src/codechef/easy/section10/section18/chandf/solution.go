package main

import (
	"bufio"
	"fmt"
	"os"
)

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

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	for tc > 0 {
		tc--
		bs, _ := reader.ReadBytes('\n')
		var X, Y, L, R uint64
		var pos int
		pos = readUint64(bs, pos, &X)
		pos = readUint64(bs, pos+1, &Y)
		pos = readUint64(bs, pos+1, &L)
		pos = readUint64(bs, pos+1, &R)

		res := solve(X, Y, L, R)
		fmt.Println(res)
	}
}

func solve(X, Y, L, R uint64) uint64 {
	if L > R {
		L, R = R, L
	}

	res := F(X, Y, R)
	ans := R

	update(&res, &ans, X, Y, L)

	O := X | Y

	var tmp uint64
	var k int

	for i := 42; i >= 0; i-- {
		if R&(1<<uint64(i)) == L&(1<<uint64(i)) {
			tmp |= R & (1 << uint64(i))
		} else {
			k = i
			break
		}
	}

	for i := k; i >= 0; i-- {
		cur := tmp

		for j := k; j >= i+1; j-- {
			cur |= R & (1 << uint64(j))
		}
		// turn off i if set
		cur &= (^(1 << uint64(i)))

		for j := i - 1; j >= 0; j-- {
			cur |= O & (1 << uint64(j))
		}

		if cur >= L && cur <= R {
			update(&res, &ans, X, Y, cur)
		}
	}

	for i := k; i >= 0; i-- {
		cur := tmp

		for j := k; j >= i+1; j-- {
			cur |= L & (1 << uint64(j))
		}

		cur |= 1 << uint64(i)

		for j := i - 1; j >= 0; j-- {
			cur |= O & (1 << uint64(j))
		}

		if cur >= L && cur <= R {
			update(&res, &ans, X, Y, cur)
		}
	}

	return ans
}

func update(res *uint64, ans *uint64, x, y, z uint64) {
	tmp := F(x, y, z)
	if *res < tmp || *res == tmp && *ans > z {
		*res = tmp
		*ans = z
	}
}

func F(x, y, z uint64) uint64 {
	return (x & z) * (y & z)
}
