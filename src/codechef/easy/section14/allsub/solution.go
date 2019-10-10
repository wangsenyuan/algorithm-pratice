package main

import (
	"bufio"
	"bytes"
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

func readInt64(bytes []byte, from int, val *int64) int {
	i := from
	var tmp int64
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + int64(bytes[i]-'0')
		i++
	}
	*val = tmp
	return i
}

func readNInt64Nums(scanner *bufio.Scanner, n int) []int64 {
	res := make([]int64, n)
	x := -1
	scanner.Scan()
	for i := 0; i < n; i++ {
		x = readInt64(scanner.Bytes(), x+1, &res[i])
	}
	return res
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	tc := readNum(scanner)

	for tc > 0 {
		tc--
		scanner.Scan()
		S := scanner.Text()
		scanner.Scan()
		R := scanner.Text()

		res := solve(S, R)

		if len(res) == 0 {
			fmt.Println("Impossible")
		} else {
			fmt.Println(res)
		}
	}
}

func solve(S, R string) string {
	if len(S) > len(R) {
		return ""
	}

	cnt := make([]int, 26)

	for i := 0; i < len(S); i++ {
		cnt[S[i]-'a']--
	}

	for i := 0; i < len(R); i++ {
		cnt[R[i]-'a']++
	}

	for i := 0; i < 26; i++ {
		if cnt[i] < 0 {
			return ""
		}
	}

	// there is an answer

	var buf bytes.Buffer

	for i := 0; i <= int(S[0]-'a'); i++ {
		for cnt[i] > 0 {
			buf.WriteByte(byte('a' + i))
			cnt[i]--
		}
	}

	buf.WriteString(S)

	for i := 0; i < 26; i++ {
		for cnt[i] > 0 {
			buf.WriteByte(byte('a' + i))
			cnt[i]--
		}
	}

	return buf.String()
}
