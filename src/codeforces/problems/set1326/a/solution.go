package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	var tc int
	fmt.Scanf("%d", &tc)

	f := bufio.NewWriter(os.Stdout)

	defer f.Flush()

	for tc > 0 {
		tc--
		var n int
		fmt.Scanf("%d", &n)
		f.WriteString(solve(n) + "\n")
	}
}

func solve(n int) string {
	if n == 1 {
		return "-1"
	}

	var buf bytes.Buffer

	buf.WriteByte('2')

	for i := 0; i < n-1; i++ {
		buf.WriteByte('3')

	}

	return buf.String()
}
