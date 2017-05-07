package main

import (
	"flag"
	"io/ioutil"
	"os"
	"fmt"
	"path/filepath"
	"strings"
)

func main() {
	dir := flag.String("d", "", "directory to start")
	flag.Parse()
	if dir == nil || len(*dir) == 0 {
		panic("main -d={where to start}")
	}
	err := filepath.Walk(*dir, visit)
	fmt.Printf("filepath.Walk() returned %v\n", err)

}

func visit(path string, f os.FileInfo, err error) error {
	if f.IsDir() || !strings.Contains(path, "com/maycur/aux/accounting") || !strings.HasSuffix(path, ".java") {
		//fmt.Println(path)
		return nil
	}

	//fmt.Println(path + "-------------")

	bytes, err := ioutil.ReadFile(path)

	if err != nil {
		return err
	}

	content := string(bytes)
	fmt.Println(content)
	//fmt.Println(path + "-------------")
	return nil
}
