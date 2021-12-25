package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	tc := readNum(scanner)

	for tc > 0 {
		tc--
		n := readNum(scanner)
		scanner.Scan()
		line := scanner.Text()
		nums, ops := parse(n, []byte(line))
		res := solve(n, nums, ops)
		fmt.Println(res)
	}
}

var f [][][]int64

const N = 200

func init() {
	f = make([][][]int64, N)
	for i := 0; i < N; i++ {
		f[i] = make([][]int64, N)
		for j := 0; j < N; j++ {
			f[i][j] = make([]int64, 2)
		}
	}
}

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

func readNum(scanner *bufio.Scanner) (a int) {
	scanner.Scan()
	readInt(scanner.Bytes(), 0, &a)
	return
}

func readTwoNums(scanner *bufio.Scanner) (a int, b int) {
	scanner.Scan()
	x := readInt(scanner.Bytes(), 0, &a)
	readInt(scanner.Bytes(), x+1, &b)
	return
}

func readNNums(scanner *bufio.Scanner, n int) []int {
	res := make([]int, n)
	x := -1
	scanner.Scan()
	for i := 0; i < n; i++ {
		x = readInt(scanner.Bytes(), x+1, &res[i])
	}
	return res
}

func parse(n int, s []byte) (nums []int, ops []byte) {
	nums = make([]int, n)
	ops = make([]byte, n-1)
	pos := readInt(s, 0, &nums[0])
	for i := 1; i < n; i++ {
		for pos < len(s) && s[pos] == ' ' {
			pos++
		}
		ops[i-1] = s[pos]
		pos++
		for pos < len(s) && s[pos] == ' ' {
			pos++
		}
		pos = readInt(s, pos, &nums[i])
	}
	return
}

func solve(n int, nums []int, ops []byte) int64 {
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			f[i][j][0] = math.MinInt32
			f[i][j][1] = math.MaxInt32
		}
		f[i][0][0] = int64(nums[i])
		f[i][0][1] = int64(nums[i])
	}
	//f[i][j] = max result in the range [i:j]

	for k := 1; k < n; k++ {
		for i := 0; i+k < n; i++ {
			//f[i][j]
			for kk := 0; kk < k; kk++ {
				if ops[i+kk] == '-' {
					f[i][k][0] = max(f[i][k][0], f[i][kk][0]-f[i+kk+1][k-kk-1][1])
					f[i][k][1] = min(f[i][k][1], f[i][kk][1]-f[i+kk+1][k-kk-1][0])
				} else {
					f[i][k][0] = max(f[i][k][0], f[i][kk][0]+f[i+kk+1][k-kk-1][0])
					f[i][k][1] = min(f[i][k][1], f[i][kk][1]+f[i+kk+1][k-kk-1][1])
				}
			}
		}
	}

	return f[0][n-1][0]
}

func max(a, b int64) int64 {
	if a >= b {
		return a
	}
	return b
}

func min(a, b int64) int64 {
	if a <= b {
		return a
	}
	return b
}
