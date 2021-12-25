package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		scanner.Scan()
		in := scanner.Text()
		if len(in) == 2 && in[0] == '4' && in[1] == '2' {
			break
		}
		fmt.Println(in)
	}
}
