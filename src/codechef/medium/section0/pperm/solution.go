package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
)

const MOD = 1000000007
const N = 5000000

var answer []int64

func init() {
	sieve := make([]bool, N+1)
	answer = make([]int64, N+1)
	answer[1] = 1
	var primeCnt int64
	for i := 2; i <= N; i++ {
		if !sieve[i] {
			primeCnt++
			for j := 2 * i; j <= N; j += i {
				sieve[j] = true
			}
		}
		answer[i] = (answer[i-1] * (primeCnt + 1)) % MOD
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

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	t := readNum(scanner)
	var buf bytes.Buffer

	for t > 0 {
		n := readNum(scanner)
		res := solve(n)
		buf.WriteString(strconv.Itoa(int(res)))
		buf.WriteByte('\n')
		t--
	}

	fmt.Print(buf.String())
}

func solve(n int) int64 {
	if n < 0 || n > N {
		panic(fmt.Sprintf("wrong input %d", n))
	}
	return answer[n]
}
