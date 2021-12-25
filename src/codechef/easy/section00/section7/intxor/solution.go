package main

import (
	"bufio"
	"fmt"
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

func readUint64Num(scanner *bufio.Scanner) uint64 {
	var num uint64
	scanner.Scan()
	readUint64(scanner.Bytes(), 0, &num)
	return num
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	tc := readNum(scanner)

	for tc > 0 {
		tc--
		n := readNum(scanner)

		res := play(scanner, n)
		if res < 0 {
			break
		}
	}
}

func play(scanner *bufio.Scanner, n int) int {
	A := make([]uint64, n+1)

	fmt.Println("1 1 2 3")
	a := readUint64Num(scanner)
	fmt.Println("1 1 2 4")
	b := readUint64Num(scanner)
	fmt.Println("1 3 4 5")
	c := readUint64Num(scanner)
	fmt.Println("1 3 4 6")
	d := readUint64Num(scanner)
	A[5] = c ^ a ^ b
	A[6] = d ^ a ^ b

	i, j, k := 5, 6, 7
	var cnt int
	for k <= n {
		if cnt > 0 && cnt%2 == 0 {
			i += 2
			j += 2
		}
		fmt.Printf("1 %d %d %d\n", i, j, k)
		num := readUint64Num(scanner)
		A[k] = num ^ A[i] ^ A[j]
		k++
		cnt++
	}

	fmt.Printf("1 %d %d %d\n", n-1, n, 1)
	A[1] = readUint64Num(scanner) ^ A[n-1] ^ A[n]

	if n&1 == 1 {
		fmt.Printf("1 %d %d %d\n", n-2, n, 2)
		A[2] = readUint64Num(scanner) ^ A[n-2] ^ A[n]
	} else {
		fmt.Printf("1 %d %d %d\n", n-1, n, 2)
		A[2] = readUint64Num(scanner) ^ A[n-1] ^ A[n]
	}

	A[3] = a ^ A[1] ^ A[2]
	A[4] = b ^ A[1] ^ A[2]

	fmt.Printf("2 ")
	for i := 1; i <= n; i++ {
		if i < n {
			fmt.Printf("%d ", A[i])
		} else {
			fmt.Printf("%d\n", A[i])
		}
	}
	res := readNum(scanner)
	if res == -1 {
		return res
	}
	return 0
}
