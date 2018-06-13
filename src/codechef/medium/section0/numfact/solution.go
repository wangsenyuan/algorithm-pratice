package main

import (
	"bufio"
	"fmt"
	"os"
)

const MAX_A = 1000000

var primes []int

func init() {
	nums := make([]bool, MAX_A)
	primes = make([]int, 78499)
	var j int
	for i := 2; i < MAX_A; i++ {
		if !nums[i] {
			primes[j] = i
			j++
			for k := 2 * i; k < MAX_A; k += i {
				nums[k] = true
			}
		}
	}
	//fmt.Printf("[debug] total %d primes", j)
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

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	t := readNum(scanner)

	for t > 0 {
		n := readNum(scanner)
		A := readNNums(scanner, n)
		fmt.Println(solve(n, A))
		t--
	}
}

func solve(n int, A []int) int64 {
	cnt := make(map[int]int)

	for _, a := range A {
		b := a
		for _, p := range primes {
			if b < p {
				break
			}
			for b%p == 0 {
				cnt[p]++
				b /= p
			}
		}
	}

	res := int64(1)

	for _, v := range cnt {
		res = res * int64(v+1)
	}

	return res
}
