package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	args := os.Args
	if len(args) <= 1 {
		log.Fatal("Usage: duplicatecode file")
	}

	fileName := args[1]

	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}

	str := string(content)
	mp := make(map[string]int)
	for strings.Contains(str, "\"businessCode\":") {
		i := strings.Index(str, "\"businessCode\":") + len("\"businessCode\":")
		// skip the first \"
		j := i + 1
		for j < len(str) && str[j] != '"' {
			j++
		}
		if j == len(str) {
			break
		}

		bizCode := str[i+1 : j]
		mp[bizCode]++
		fmt.Println(bizCode)
		str = strings.Replace(str, "\"businessCode\":", "", 1)
	}

	fmt.Println("---duplicate-----")
	for k, v := range mp {
		if v > 1 {
			fmt.Println(k)
		}
	}
}
