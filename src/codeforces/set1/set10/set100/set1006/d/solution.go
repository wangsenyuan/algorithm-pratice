package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	readString(reader)
	a := readString(reader)
	b := readString(reader)
	res := solve(a, b)
	fmt.Println(res)
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
}

func solve(a string, b string) int {
	n := len(a)
	var res int

	// 这里不对，只能改变a
	for i, j := 0, n-1; i <= j; i, j = i+1, j-1 {
		if i == j {
			if a[i] != b[i] {
				res++
			}
		} else {
			x, y := min(a[i], a[j]), max(a[i], a[j])
			u, v := min(b[i], b[j]), max(b[i], b[j])
			if x == y {
				if u != v {
					if x != u && x != v {
						// 只能将x\y修改成u, v才行
						res += 2
					} else {
						// 只需要修改一个x成不同的那个即可
						res++
					}
				}
				// else u = v, 只需要交换即可
			} else {
				// x < y
				if u == v {
					// 只需要让x=y即可
					res++
				} else {
					// u != v and x != y
					if x == u && y == v {
						continue
					}
					if x == u || x == v || y == u || y == v {
						res++
					} else {
						res += 2
					}
				}
			}
		}
	}
	return res
}
