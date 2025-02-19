package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n := readNum(reader)
	res := solve(n)
	var buf bytes.Buffer
	for _, row := range res {
		for _, x := range row {
			buf.WriteString(fmt.Sprintf("%d ", x))
		}
		buf.WriteByte('\n')
	}
	buf.WriteTo(os.Stdout)
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

func solve(n int) [][]int {
	a := make([][]int, n)
	for i := 0; i < n; i++ {
		a[i] = make([]int, n)
	}
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1; j++ {
			a[i][j] = (i + j) % (n - 1)
			if a[i][j] == 0 {
				a[i][j] = n - 1
			}
		}
		// need to swap (i, i) to (i, n - 1)
		if i == 0 {
			a[i][n-1] = n - 1
			a[n-1][i] = n - 1
		} else {
			a[i][n-1] = a[i][i]
			a[n-1][i] = a[i][i]
			a[i][i] = 0
		}
	}
	a[0][0] = 0
	return a
}
