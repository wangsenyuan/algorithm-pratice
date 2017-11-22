package main

import (
	"bufio"
	"fmt"
	"os"
)

func readInt(bytes []byte, from int, val *int) int {
	i := from
	tmp := 0
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp
	return i
}

func readIntPair(bytes []byte, a *int, b *int) {
	i := readInt(bytes, 0, a)
	readInt(bytes, i+1, b)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var t int
	scanner.Scan()
	readInt(scanner.Bytes(), 0, &t)

	for i := 0; i < t; i++ {
		var n int
		scanner.Scan()
		readInt(scanner.Bytes(), 0, &n)
		pos := make([][]int, n)
		for j := 0; j < n; j++ {
			var a, b int
			scanner.Scan()
			readIntPair(scanner.Bytes(), &a, &b)
			pos[j] = []int{a, b}
		}

		var m int
		scanner.Scan()
		readInt(scanner.Bytes(), 0, &m)
		strips := make([][]int, m)
		for j := 0; j < m; j++ {
			var a, b int
			scanner.Scan()
			readIntPair(scanner.Bytes(), &a, &b)
			strips[j] = []int{a, b}
		}
		res := solve(n, pos, m, strips)
		fmt.Println(res)
	}
}

const MAX_SIDE_LENGTH = 2500000
const INF = int64(1000000000000000000)

func solve(n int, pos [][]int, m int, strips [][]int) int64 {

	f := make([]int64, MAX_SIDE_LENGTH)

	for l := 0; l < MAX_SIDE_LENGTH; l++ {
		f[l] = INF
	}

	f[0] = 0

	for l := 0; l < MAX_SIDE_LENGTH; l++ {
		for i := 0; i < m; i++ {
			a, b := strips[i][0], int64(strips[i][1])

			if l+a < MAX_SIDE_LENGTH && f[l]+b < f[l+a] {
				f[l+a] = f[l] + b
			}
		}
	}

	for l := MAX_SIDE_LENGTH - 2; l >= 0; l-- {
		if f[l] > f[l+1] {
			f[l] = f[l+1]
		}
	}
	var res int64

	for i := 1; i <= n; i++ {
		var a, b int64
		if i < n {
			a = int64(pos[i][0] - pos[i-1][0])
			b = int64(pos[i][1] - pos[i-1][1])
		} else {
			a = int64(pos[0][0] - pos[n-1][0])
			b = int64(pos[0][1] - pos[n-1][1])
		}
		length := a*a + b*b
		left, right := 0, MAX_SIDE_LENGTH
		for left < right {
			mid := left + (right-left)/2
			if int64(mid)*int64(mid) < length {
				left = mid + 1
			} else {
				right = mid
			}
		}

		res += f[left]
	}

	return res
}
