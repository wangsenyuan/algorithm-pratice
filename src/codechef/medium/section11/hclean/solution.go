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
		n, d := readTwoNums(reader)

		S := readNNums(reader, n)

		res := solve(n, d, S)

		if len(res) == 0 {
			fmt.Println("-2")
			continue
		}
		for i := 0; i < len(res); i++ {
			print(res[i], n)
		}
	}

}

func print(x int, n int) {
	buf := make([]byte, 0, 2*n)

	for i := 0; i < n; i++ {
		if x&(1<<uint(i)) == 0 {
			buf = append(buf, '-')
		}
		buf = append(buf, '1')
		buf = append(buf, ' ')
	}
	fmt.Printf("%s\n", string(buf))
}

func solve(n, D int, S []int) []int {
	if D < 4 {
		return nil
	}

	var loop func(n int) []int
	loop = func(n int) []int {
		if n == 1 {
			return []int{0, 1}
		}
		tmp := loop(n - 1)

		res := make([]int, 1<<uint(n))

		for i := 0; i < 1<<uint(n-1); i++ {
			res[i] = tmp[i]
		}
		for i := 0; i < 1<<uint(n-1); i++ {
			res[i+(1<<uint(n-1))] = 1<<uint(n-1) | tmp[1<<uint(n-1)-i-1]
		}
		return res
	}

	res := loop(n)

	var X int

	for i := 0; i < n; i++ {
		if S[i] == 1 {
			X |= 1 << uint(i)
		}
	}

	var k int

	for k < len(res) && res[k] != X {
		k++
	}

	ans := make([]int, len(res))
	copy(ans, res[k:])
	copy(ans[len(res)-k:], res)

	return ans
}
