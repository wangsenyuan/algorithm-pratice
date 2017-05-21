package main

import (
	"strings"
	"sort"
	"fmt"
)

type Dir struct {
	name      string
	fileNames []string
	files     map[string]string
	subDirs   map[string]*Dir
}

type FileSystem struct {
	root *Dir
}

func Constructor() FileSystem {
	return FileSystem{root: &Dir{"/", nil, nil, nil}}
}

func (this *FileSystem) Ls(path string) []string {
	dir := this.root
	segs := strings.Split(path, "/")[1:]
	i := 0
	for ; i < len(segs)-1; i++ {
		seg := segs[i]
		if len(dir.subDirs) == 0 {
			return nil
		}
		if subDir, found := dir.subDirs[seg]; found {
			dir = subDir
		} else {
			return nil
		}
	}

	if len(segs[i]) == 0 {
		//whole directory
		return dir.fileNames
	}

	if subDir, isDir := dir.subDirs[segs[i]]; isDir {
		return subDir.fileNames
	}

	res := make([]string, 0, len(dir.fileNames))

	for _, name := range dir.fileNames {
		if strings.HasPrefix(name, segs[i]) {
			res = append(res, name)
		}
	}
	return res
}

func (this *FileSystem) Mkdir(path string) {
	dir := this.root
	segs := strings.Split(path, "/")[1:]
	for i := 0; i < len(segs); i++ {
		if subDir, found := dir.subDirs[segs[i]]; found {
			dir = subDir
		} else {
			subDir = &Dir{segs[i], nil, nil, nil}
			if dir.subDirs == nil {
				dir.subDirs = make(map[string]*Dir)
			}
			if dir.fileNames == nil {
				dir.fileNames = make([]string, 0, 10)
			}
			dir.fileNames = append(dir.fileNames, segs[i])
			sort.Strings(dir.fileNames)
			dir.subDirs[segs[i]] = subDir
			dir = subDir
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
		if dir.files == nil {
			dir.files = make(map[string]string)
		}
		dir.files[fileName] = content
		if dir.fileNames == nil {
			dir.fileNames = make([]string, 0, 10)
		}
		dir.fileNames = append(dir.fileNames, fileName)
		sort.Strings(dir.fileNames)
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
	fs.AddContentToFile("/a/b/c/d", "hello")
	fmt.Println(fs.Ls("/"))
	fmt.Print(fs.ReadContentFromFile("/a/b/c/d"))
}
