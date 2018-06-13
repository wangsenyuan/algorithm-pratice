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

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	t := readNum(scanner)

	for t > 0 {
		scanner.Scan()
		bs := scanner.Bytes()
		var spactAt int
		for spactAt < len(bs) && bs[spactAt] != ' ' {
			spactAt++
		}
		s := bs[:spactAt]
		var K int
		readInt(bs, spactAt+1, &K)
		res := solve(s, K)
		fmt.Println(res)
		t--
	}
}

func solve(s []byte, K int) int {
	var ans int

	for i := 0; i < len(s); i++ {
		occ := make([]int, 26)
		var cnt int
		var max int
		for j := i; j >= 0; j-- {
			if occ[int(s[j]-'a')] == 0 {
				cnt++
			}
			occ[int(s[j]-'a')]++
			if occ[int(s[j]-'a')] > max {
				max = occ[int(s[j]-'a')]
			}
			if cnt >= K && max*cnt == i-j+1 {
				ans++
			}
		}
	}

	return ans
}
