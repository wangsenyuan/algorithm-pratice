package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
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

func readTwoNum(scanner *bufio.Scanner) (a int, b int) {
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

	n, k := readTwoNum(scanner)

	twts := make([]int, n)

	round := 1
	cnt := 0

	var buf bytes.Buffer

	for i := 0; i < k; i++ {
		scanner.Scan()
		bs := scanner.Bytes()

		if bs[2] == 'I' {
			// click
			var x int
			readInt(bs, 6, &x)
			if twts[x-1] == round {
				twts[x-1]--
				cnt--
			} else {
				twts[x-1] = round
				cnt++
			}
		} else {
			cnt = 0
			round++
		}
		buf.WriteString(strconv.Itoa(cnt))
		if i < k-1 {
			buf.WriteByte('\n')
		}
	}

	fmt.Println(buf.String())
}
