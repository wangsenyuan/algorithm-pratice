package main

import (
	"bufio"
	"fmt"
	"math"
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
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
}

func readNum(scanner *bufio.Scanner) (a int) {
	scanner.Scan()
	readInt(scanner.Bytes(), 0, &a)
	return
}

func readTwoNums(scanner *bufio.Scanner) (a int, b int) {
	res := readNNums(scanner, 2)
	a, b = res[0], res[1]
	return
}

func readNNums(scanner *bufio.Scanner, n int) []int {
	res := make([]int, n)
	x := 0
	scanner.Scan()
	for i := 0; i < n; i++ {
		for x < len(scanner.Bytes()) && scanner.Bytes()[x] == ' ' {
			x++
		}
		x = readInt(scanner.Bytes(), x, &res[i])
	}
	return res
}

func fillNNums(scanner *bufio.Scanner, n int, res []int) {
	x := 0
	scanner.Scan()
	for i := 0; i < n; i++ {
		for x < len(scanner.Bytes()) && scanner.Bytes()[x] == ' ' {
			x++
		}
		x = readInt(scanner.Bytes(), x, &res[i])
	}
}

func readUint64(bytes []byte, from int, val *uint64) int {
	i := from

	var tmp uint64
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + uint64(bytes[i]-'0')
		i++
	}
	*val = tmp

	return i
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	tc := readNum(scanner)

	for tc > 0 {
		tc--
		k := readNum(scanner)
		scanner.Scan()
		bs := scanner.Bytes()
		pos := -1
		res := solve(k, func() uint64 {
			var num uint64
			pos = readUint64(bs, pos+1, &num)
			return num
		})
		fmt.Println(res)
	}
}

var pow2 [60]uint64

func init() {
	pow2[0] = 1
	for i := 1; i < 60; i++ {
		pow2[i] = pow2[i-1] * 2
	}
}

func solve(k int, fn func() uint64) uint64 {
	if k%2 == 0 {
		return 0
	}

	nums := make([]uint64, k)

	for i := 0; i < k; i++ {
		x := fn()
		nums[i] = x
	}

	//k is odd
	for i := 0; i < k; i++ {
		if nums[i]&1 == 0 {
			//could only use it as 1
			return 1
		}
	}
	//all k is odd
	var res uint64 = math.MaxUint64
	for i := 0; i < k; i++ {
		var pos uint
		for nums[i]&(1<<pos) != 0 {
			pos++
		}
		res = min(res, pow2[pos])
	}

	return res
}

func min(a, b uint64) uint64 {
	if a <= b {
		return a
	}
	return b
}
