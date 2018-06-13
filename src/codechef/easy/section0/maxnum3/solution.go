package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func readInt(bytes []byte, from int, val *int) int {
	i := from
	tmp := 0
	for i < len(bytes) && bytes[i] == ' ' {
		i++
	}
	sign := 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
}

func readOneNum(scanner *bufio.Scanner) (a int) {
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

func toIntArray(s []byte) []int {
	n := len(s)
	res := make([]int, n)
	for i := 0; i < n; i++ {
		res[i] = int(s[i] - '0')
	}
	return res
}

func intsArrayToString(s []int) string {
	var buf bytes.Buffer

	for i := 0; i < len(s); i++ {
		buf.WriteByte(byte(s[i] + '0'))
	}

	return buf.String()
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	t := readOneNum(scanner)

	for t > 0 {
		scanner.Scan()
		s := toIntArray(scanner.Bytes())
		can, res := solve(s)
		if !can {
			fmt.Println(-1)
		} else {
			fmt.Printf("%s\n", intsArrayToString(res))
		}
		t--
	}
}

func solve(s []int) (bool, []int) {
	n := len(s)

	var sum int

	for i := 0; i < n; i++ {
		sum = (sum + s[i]) % 3
	}

	res := make([]int, n-1)
	pos := -1
	for i := 0; i < n; i++ {
		if s[i]%3 == sum {
			if i == n-1 {
				if s[n-2]%2 == 0 {
					pos = i
				}
			} else {
				if s[n-1]%2 == 0 {
					pos = i
					if s[i] < s[i+1] {
						break
					}
				}
			}
		}
	}

	if pos < 0 {
		return false, nil
	}

	copy(res[:pos], s[:pos])
	copy(res[pos:], s[pos+1:])
	return true, res
}
