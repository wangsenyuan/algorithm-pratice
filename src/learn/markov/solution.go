package main

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strings"
)

const NON_WORD = "#"

func main() {
	table := build(os.Stdin)
	gnerate(table, 1024, os.Stdout)
}

func build(r io.Reader) *Map {
	table := NewMap(4096)

	buf := bufio.NewScanner(r)
	pref := make([]string, 2)
	pref[0] = NON_WORD
	pref[1] = NON_WORD
	for buf.Scan() {
		line := buf.Text()
		ss := strings.Split(line, " ")
		for _, s := range ss {
			if len(s) == 0 {
				continue
			}
			table.AddWord(pref, s)
			pref[0] = pref[1]
			pref[1] = s
		}
	}
	table.AddWord(pref, NON_WORD)
	return &table
}

func gnerate(table *Map, num int, w io.Writer) {
	pref := make([]string, 2)
	pref[0] = NON_WORD
	pref[1] = NON_WORD

	buf := bufio.NewWriter(w)
	for i := 0; i < num; i++ {
		sp := table.Lookup(pref)
		var nm int
		var word string
		if sp == nil {
			fmt.Fprintf(os.Stderr, "no state found for %v", pref)
			break
		}
		for suf := sp.suffix; suf != nil; suf = suf.next {
			nm++
			if randNum(nm) == 0 {
				word = suf.word
			}
		}
		if word == NON_WORD {
			break
		}
		pref[0] = pref[1]
		pref[1] = word
		buf.WriteString(word)
		buf.WriteByte(' ')
	}
	buf.Flush()
}

func randNum(n int) int {
	return rand.Intn(n)
}
