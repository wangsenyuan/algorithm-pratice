package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	line1, _ := reader.ReadString('\n')
	line2, _ := reader.ReadString('\n')
	line1 = strings.Trim(line1, "\n")
	line2 = strings.Trim(line2, "\n")

	line3 := replace(line1, line2)

	fmt.Println(line3)
}

func replace(source, rep string) string {
	result := strings.Replace(source, rep, "", -1)
	if len(result) == len(source) {
		return result
	}

	return replace(result, rep)
}
