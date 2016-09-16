package main

import "fmt"

func main() {
	board := [][]byte{
		[]byte{'o', 'a', 'a', 'n'},
		[]byte{'e', 't', 'a', 'e'},
		[]byte{'i', 'h', 'k', 'r'},
		[]byte{'i', 'f', 'l', 'v'},
	}

	words := []string{"oath", "pea", "eat", "rain"}

	res := findWords(board, words)

	fmt.Println(res)
}

var dx = [4]int{-1, 0, 0, 1}
var dy = [4]int{0, -1, 1, 0}

func findWords(board [][]byte, words []string) []string {
	t := &Trie{make([]*Trie, 26), false}

	for _, word := range words {
		t.Add([]byte(word), 0)
	}

	checked := make([][]bool, len(board))

	for i := range checked {
		checked[i] = make([]bool, len(board[i]))
	}

	var res []string
	added := make(map[string]bool)

	var search func(i, j int, path string)

	search = func(i, j int, path string) {
		if !t.HasPrefix([]byte(path), 0) {
			return
		}
		next := path + string(board[i][j])

		if t.Search([]byte(next), 0) && !added[next] {
			res = append(res, next)
			added[next] = true
		}

		checked[i][j] = true

		for k := range dx {
			x, y := dx[k]+i, dy[k]+j
			if x < 0 || x >= len(board) || y < 0 || y >= len(board[i]) {
				continue
			}

			if checked[x][y] {
				continue
			}
			search(x, y, next)
		}

		checked[i][j] = false
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			search(i, j, "")
		}
	}

	return res
}

type Trie struct {
	children []*Trie
	end      bool
}

func (t *Trie) Add(s []byte, i int) {
	if i == len(s) {
		t.end = true
		return
	}

	if t.children[s[i]-'a'] == nil {
		t.children[s[i]-'a'] = &Trie{make([]*Trie, 26), false}
	}
	t.children[s[i]-'a'].Add(s, i+1)
}

func (t Trie) HasPrefix(s []byte, i int) bool {
	if i == len(s) {
		return true
	}

	if t.children[s[i]-'a'] == nil {
		return false
	}

	return t.children[s[i]-'a'].HasPrefix(s, i+1)
}

func (t Trie) Search(s []byte, i int) bool {
	if i == len(s) {
		return t.end
	}

	if t.children[s[i]-'a'] == nil {
		return false
	}

	return t.children[s[i]-'a'].Search(s, i+1)
}
