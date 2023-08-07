package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	var buf bytes.Buffer

	for tc > 0 {
		tc--
		d := readNum(reader)
		res := solve(d)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}

	fmt.Print(buf.String())
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

func solve(d int) int {
	var primes []int
	for x := d + 1; ; x++ {
		ok := true
		for i := 2; i <= x/i; i++ {
			if x%i == 0 {
				ok = false
				break
			}
		}
		if ok {
			primes = append(primes, x)
			break
		}
	}
	for x := primes[0] + d; ; x++ {
		ok := true
		for i := 2; i <= x/i; i++ {
			if x%i == 0 {
				ok = false
				break
			}
		}
		if ok {
			primes = append(primes, x)
			break
		}
	}

	ans := min(primes[0]*primes[1], primes[0]*primes[0]*primes[0])
	return ans
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
