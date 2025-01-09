package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	res := process(reader)

	fmt.Println(res)
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

func process(reader *bufio.Reader) int {
	n := readNum(reader)
	a := readNNums(reader, n)
	b := readNNums(reader, n)
	return solve(a, b)
}

func solve(a []int, b []int) int {
	n := len(a)

	calc := func(arr []int) [][]int {
		res := make([][]int, 4)
		for i := range 4 {
			res[i] = make([]int, n+1)
		}

		for i := 0; i < n; i++ {
			res[0][i] = i * arr[i]
		}

		for i := n - 1; i >= 0; i-- {
			res[1][i] = res[0][i] + res[1][i+1]
			res[2][i] = (n-1-i)*arr[i] + res[2][i+1]
			res[3][i] = arr[i] + res[3][i+1]
		}

		return res
	}

	x := calc(a)
	y := calc(b)

	var ans int
	var sum int

	for i := 0; i < n; i++ {
		tmp := sum
		if i&1 == 0 {
			time := 2 * i
			sum += time * a[i]
			sum += (time + 1) * b[i]
			tmp += (time-i)*x[3][i] + x[1][i]
			time += n - i
			tmp += time*y[3][i] + y[2][i]
		} else {
			time := 2 * i
			sum += time * b[i]
			sum += (time + 1) * a[i]
			tmp += (time-i)*y[3][i] + y[1][i]
			time += n - i
			tmp += time*x[3][i] + x[2][i]
		}
		ans = max(ans, tmp)
	}

	return max(ans, sum)
}
