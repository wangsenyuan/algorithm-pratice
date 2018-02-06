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

func readNum(scanner *bufio.Scanner) (a int) {
	scanner.Scan()
	readInt(scanner.Bytes(), 0, &a)
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

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	t := readNum(scanner)

	for i := 0; i < t; i++ {
		n := readNum(scanner)
		nums := readNNums(scanner, n)
		res := solve(n, nums)
		fmt.Println(res)
	}
}

const MOD = 1000000007

func solve(n int, nums []int) int64 {
	count := make(map[int]int64)

	for i, j := 0, 0; i <= n; i++ {
		if i < n && nums[i] == nums[j] {
			continue
		}
		d := int64(i - j)
		tmp := (d * (d + 1) / 2) % MOD
		count[nums[j]] = (count[nums[j]] + tmp) % MOD
		j = i
	}

	ans := int64(1)

	for _, v := range count {
		ans = (ans * v) % MOD
	}

	return ans
}
