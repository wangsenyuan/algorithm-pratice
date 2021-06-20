package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n := readNum(reader)
	A := readNNums(reader, n)
	res := solve(n, A)
	fmt.Printf("%d\n", len(res))
	if len(res) > 0 {
		fmt.Printf("%d\n", res[0])
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

const X = 2000

func solve(n int, A []int) []int {
	dp := make([]bool, 1+n*X)
	fp := make([]bool, 1+n*X)
	var sum int
	dp[0] = true
	for _, cur := range A {
		for x := 0; x <= sum; x++ {
			fp[x] = false
		}
		for x := 0; x <= sum; x++ {
			if dp[x] {
				if x+cur < len(fp) {
					fp[x+cur] = true
				}
				if x-cur >= 0 {
					fp[x-cur] = true
				} else {
					fp[cur-x] = true
				}
			}
		}
		copy(dp, fp)
		sum += cur
	}
	if !dp[0] {
		return nil
	}

	g := A[0]

	for i := 0; i < n; i++ {
		if A[i]%2 == 1 {
			return []int{i + 1}
		}
		g = gcd(g, A[i])
	}

	for i := 0; i < n; i++ {
		if (A[i]/g)%2 == 1 {
			return []int{i + 1}
		}
	}
	return nil
}

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}
