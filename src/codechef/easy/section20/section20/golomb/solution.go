package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	for tc > 0 {
		tc--
		var l, r uint64
		s, _ := reader.ReadBytes('\n')
		pos := readUint64(s, 0, &l)
		readUint64(s, pos + 1, &r)
		fmt.Println(solve(int64(l), int64(r)))
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


const SIZE = 2000000
const MOD = 1000000007
var g [SIZE]int64
var count [SIZE]int64
var sum [SIZE]int64

func init() {
	g[1] = 1
	g[2] = 2
	g[3] = 3

	for i, cur := 3, 4; i < SIZE; i++ {
		for j := int64(0); j < g[i] && cur < SIZE; j++ {
			g[cur] = int64(i)
			cur++
		}
	}

	for i := 1; i < SIZE; i++ {
		count[i] = count[i-1] + g[i]
	}

	for i := 1; i < SIZE; i++{
		I := int64(i)
		sum[i] = (sum[i-1] + ( (I * I) % MOD * g[i]) % MOD) % MOD
	}
}

func solve(l, r int64) int {
	return (query(r) - query(l - 1) + MOD) % MOD
}

func query(n int64) int {
	i := search(SIZE, func(i int) bool {
		return count[i] >= n
	})
	if i == 0 {
		return 0
	}
	I := int64(i)

	ans := (sum[i-1] + ((I * I) % MOD * (n - count[i - 1]) )% MOD) % MOD

	return int(ans)
}

func search(n int,fn func(i int) bool ) int {
	var left, right = 0, n
	for left < right {
		mid := (left + right) / 2
		if fn(mid) {
			right = mid
		} else {
			left =mid + 1
		}
	}
	return right
}
