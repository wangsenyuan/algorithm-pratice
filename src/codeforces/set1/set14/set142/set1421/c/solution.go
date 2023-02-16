package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	s, _ := reader.ReadString('\n')
	res := solve(s)
	fmt.Print(res)
}

func solve(s string) string {
	var n int
	for n < len(s) && (s[n] >= 'a' && s[n] <= 'z') {
		n++
	}

	var res bytes.Buffer
	res.WriteString("3\n")
	res.WriteString("L 2\n")
	m := n + 1
	res.WriteString("R 2\n")
	m += n - 1
	res.WriteString(fmt.Sprintf("R %d\n", m-1))
	return res.String()
}
