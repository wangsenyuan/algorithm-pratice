package main

import (
	"bufio"
	"bytes"
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
	var buf bytes.Buffer
	for tc > 0 {
		tc--
		n, k := readTwoNums(reader)
		res := solve(n, k)
		for i := 0; i < len(res); i++ {
			buf.WriteString(fmt.Sprintf("%d %d\n", res[i][0], res[i][1]))
		}
	}
	fmt.Print(buf.String())
}

func solve(n, k int) [][]int {
	N := 1 << uint(n)

	x := make([]int, N)
	for i := 0; i < N; i++ {
		x[i] = i ^ (i >> 1)
	}
	res := make([][]int, 0, k)
	ans := make([]int, 0, k)

	if k%2 == 1 {
		y := make([]int, 0, (k+1)/2)
		z := make([]int, 0, k-cap(y))
		for i := 0; i < (k+1)/2; i++ {
			y = append(y, x[i])
		}
		for i := (k + 1) / 2; i < k; i++ {
			z = append(z, x[i])
		}

		for i := 0; i < k/2; i++ {
			ans = append(ans, y[i])
			ans = append(ans, z[i])
		}
		ans = append(ans, y[k/2])
	} else {
		for i := 0; i < k; i++ {
			if i%2 == 0 {
				ans = append(ans, x[i/2])
			} else {
				ans = append(ans, 1<<uint(n-1)+x[i/2])
			}
		}
	}
	for i := 0; i < len(ans); i++ {
		res = append(res, []int{ans[i], ans[(i+1)%len(ans)]})
	}
	return res
}

func digitOneCount(num int) int {
	var res int
	for num > 0 {
		res++
		num -= num & -num
	}
	return res
}
