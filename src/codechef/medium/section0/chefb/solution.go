package main

import (
	"bufio"
	"os"
	"fmt"
	"sort"
)

var primes []int

func init() {
	sieve := make([]bool, 1000001)
	primes = make([]int, 1000000)

	i := 0
	for x := 2; x <= 1000000; x++ {
		if !sieve[x] {
			primes[i] = x
			i++
			if x < 1000 {
				for y := x * x; y <= 1000000; y += x {
					sieve[y] = true
				}
			}
		}
	}

	primes = primes[:i]
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

func solve(n int, A []int) int {
	cnt := make([]int, len(primes))

	for i := 0; i < n; i++ {
		x := A[i]
		for j := 0; x >= primes[j]*primes[j]; j++ {
			var k int
			for x%primes[j] == 0 {
				k++
				x /= primes[j]
			}
			if k > cnt[j] {
				cnt[j] = k
			}
		}
		if x > 1 {
			j := sort.SearchInts(primes, x)
			if cnt[j] == 0 {
				cnt[j] = 1
			}
		}
	}
	var res int
	for i := 0; i < len(primes); i++ {
		res += cnt[i]
	}

	return res
}

func main2() {
	var t int
	fmt.Scanf("%d", &t)

	for t > 0 {
		cnt := make(map[int]int)

		var n int
		fmt.Scanf("%d", &n)

		for n > 0 {
			var num int
			fmt.Scanf("%d", &num)

			for j := 0; primes[j]*primes[j] <= num; j++ {
				var k int
				for num%primes[j] == 0 {
					k++
					num /= primes[j]
				}
				if k > cnt[j] {
					cnt[j] = k
				}
			}
			if num > 1 {
				j := sort.SearchInts(primes, num)
				if cnt[j] == 0 {
					cnt[j] = 1
				}
			}

			n--
		}

		var res int
		for _, v := range cnt {
			res += v
		}

		fmt.Println(res)
		t--
	}
}
