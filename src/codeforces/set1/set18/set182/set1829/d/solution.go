package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)

	defer out.Flush()

	var tc int
	fmt.Fscan(in, &tc)

	for tc > 0 {
		tc--
		var n, expect int
		fmt.Fscan(in, &n, &expect)
		res := solve(n, expect)

		if res {
			fmt.Fprintln(out, "YES")
		} else {
			fmt.Fprintln(out, "NO")
		}
	}
}

func solve(n int, expect int) bool {
	if n == expect {
		return true
	}
	if n < expect || n%3 != 0 {
		return false
	}

	m := n / 3
	return solve(m, expect) || solve(m*2, expect)
}
