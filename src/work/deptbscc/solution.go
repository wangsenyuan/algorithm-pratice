package main

import (
	"bufio"
	"os"
	"io"
	"strings"
	"fmt"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		text, rs := reader.ReadString('\n')
		if rs == io.EOF {
			break
		}
		parts := strings.Split(strings.Trim(text, "\n"), "\t")
		old, new := parts[0], parts[1]
		fmt.Printf("UPDATE department_information SET business_code = '%s' WHERE ent_code = 'EC17052920MESD8G' AND business_code = '%s';\n ", new, old)
	}
}
