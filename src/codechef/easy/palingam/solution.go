package main

import (
	"bufio"
	"fmt"
	"os"
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

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var t int
	scanner.Scan()
	readInt(scanner.Bytes(), 0, &t)

	for i := 0; i < t; i++ {
		scanner.Scan()
		a := scanner.Text()
		scanner.Scan()
		b := scanner.Text()
		res := solve(a, b)
		if res {
			fmt.Println("A")
		} else {
			fmt.Println("B")
		}
	}
}

func solve(a, b string) bool {
	acnt := make([]int, 26)
	bcnt := make([]int, 26)

	for i := 0; i < len(a); i++ {
		acnt[a[i]-'a']++
		bcnt[b[i]-'a']++
	}

	bWon := true
	for i := 0; i < 26; i++ {
		if acnt[i] > 0 && bcnt[i] == 0 {
			bWon = false
			break
		}
	}
	if bWon {
		return false
	}

	for i := 0; i < 26; i++ {
		if acnt[i] > 1 && bcnt[i] == 0 {
			return true
		}
		if acnt[i] > 0 && bcnt[i] == 0 {
			j := 0
			for j < 26 {
				if i != j {
					if bcnt[j] > 0 && acnt[j] == 0 {
						break
					}
				}
				j++
			}
			if j == 26 {
				// B can't find a unique letter after A pick i
				return true
			}
		}
	}

	return false
}
