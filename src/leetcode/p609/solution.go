package main

import (
	"strings"
	"fmt"
)

func findDuplicate(paths []string) [][]string {
	contentMap := make(map[string][]string)

	for _, path := range paths {

		strs := strings.Split(path, " ")
		dir := strs[0]

		for i := 1; i < len(strs); i++ {
			str := strs[i]
			a := strings.Index(str, "(")
			b := strings.Index(str, ")")
			content := str[a+1: b]

			file := dir + "/" + str[:a]
			if files, found := contentMap[content]; found {
				contentMap[content] = append(files, file)
			} else {
				contentMap[content] = []string{file}
			}
		}
	}

	var res [][]string
	for _, files := range contentMap {
		if len(files) > 1 {
			res = append(res, files)
		}
	}
	return res
}

func main() {
	files := []string{"root/a 1.txt(abcd) 2.txt(efgh)", "root/c 3.txt(abcd)", "root/c/d 4.txt(efgh)", "root 4.txt(efgh)"}
	fmt.Println(findDuplicate(files))
}
