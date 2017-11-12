package main

import (
	"strings"
	"sort"
	"fmt"
)

type Dir struct {
	files   map[string]string
	subDirs map[string]*Dir
}

type FileSystem struct {
	root *Dir
}

func Constructor() FileSystem {
	return FileSystem{root: &Dir{make(map[string]string), make(map[string]*Dir)}}
}

func (this *FileSystem) Ls(path string) []string {
	dir := this.root

	if path != "/" {
		segs := strings.Split(path, "/")
		i := 1
		for ; i < len(segs)-1; i++ {
			dir = dir.subDirs[segs[i]]
		}

		if _, isFile := dir.files[segs[i]]; isFile {
			return []string{segs[i]}
		}
		dir = dir.subDirs[segs[i]]
	}

	var res []string
	for dirName, _ := range dir.subDirs {
		res = append(res, dirName)
	}

	for fileName, _ := range dir.files {
		res = append(res, fileName)
	}
	sort.Strings(res)
	return res
}

func (this *FileSystem) Mkdir(path string) {
	if path == "/" {
		return
	}
	segs := strings.Split(path, "/")

	dir := this.root
	for i := 1; i < len(segs); i++ {
		if dir.subDirs[segs[i]] == nil {
			subDir := &Dir{make(map[string]string), make(map[string]*Dir)}
			dir.subDirs[segs[i]] = subDir
			dir = subDir
		} else {
			dir = dir.subDirs[segs[i]]
		}
	}

}

func (this *FileSystem) AddContentToFile(filePath string, content string) {
	segs := strings.Split(filePath, "/")
	dirPath := strings.Join(segs[:len(segs)-1], "/")
	this.Mkdir(dirPath)
	fileName := segs[len(segs)-1]

	dir := this.root
	for i := 1; i < len(segs)-1; i++ {
		dir = dir.subDirs[segs[i]]
	}

	if file, found := dir.files[fileName]; found {
		dir.files[fileName] = file + content
	} else {
		dir.files[fileName] = content
	}
}

func (this *FileSystem) ReadContentFromFile(filePath string) string {
	segs := strings.Split(filePath, "/")[1:]

	dir := this.root
	for i := 0; i < len(segs)-1; i++ {
		if subDir, found := dir.subDirs[segs[i]]; found {
			dir = subDir
		} else {
			return ""
		}
	}
	if file, found := dir.files[segs[len(segs)-1]]; found {
		return file
	}
	return ""
}

func main() {
	fs := Constructor()
	fmt.Println(fs.Ls("/"))
	fs.Mkdir("/a/b/c")
	fmt.Println(fs.Ls("/a/b"))
	//fmt.Println(fs.Ls("/a/b/"))

	fs.AddContentToFile("/a/b/c/d", "hello")
	fmt.Println(fs.Ls("/"))
	fmt.Print(fs.ReadContentFromFile("/a/b/c/d"))
}
