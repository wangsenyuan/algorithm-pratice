package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func readInt(bytes []byte, from int, val *int) int {
	i := from
	sign := 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	tmp := 0
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
}

func readNum(scanner *bufio.Scanner) (a int) {
	scanner.Scan()
	readInt(scanner.Bytes(), 0, &a)
	return
}

func readTwoNums(scanner *bufio.Scanner) (a int, b int) {
	res := readNNums(scanner, 2)
	a, b = res[0], res[1]
	return
}

func readNNums(scanner *bufio.Scanner, n int) []int {
	res := make([]int, n)
	x := 0
	scanner.Scan()
	for i := 0; i < n; i++ {
		for x < len(scanner.Bytes()) && scanner.Bytes()[x] == ' ' {
			x++
		}
		x = readInt(scanner.Bytes(), x, &res[i])
	}
	return res
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	tc := readNum(scanner)

	for x := 1; x <= tc; x++ {
		n := readNum(scanner)
		var buf bytes.Buffer
		for i := 0; i < n; i++ {
			scanner.Scan()
			buf.WriteString(scanner.Text())
		}
		S := buf.String()
		buf.Reset()
		m := readNum(scanner)
		animals := make([][]string, m)
		for i := 0; i < m; i++ {
			scanner.Scan()
			line := scanner.Text()
			var cnt int
			for j := 0; j < len(line); j++ {
				if unicode.IsSpace(rune(line[j])) {
					cnt++
				}
				if cnt == 2 {
					line = line[j+1:]
					break
				}
			}
			animals[i] = strings.Split(line, " ")
		}
		fmt.Printf("Case #%d:\n", x)
		res := solve(S, animals)
		for _, ans := range res {
			fmt.Printf("%.7f\n", ans)
		}
	}
}

func solve(S string, animals [][]string) []float64 {
	root := ParseAsTree(normalize(S))
	res := make([]float64, len(animals))
	for i := 0; i < len(animals); i++ {
		res[i] = root.Decide(animals[i])
	}

	return res
}

type Tree struct {
	weight  float64
	feature string
	leaf    bool
	yes     *Tree
	no      *Tree
}

func normalize(S string) string {
	var buf bytes.Buffer

	var pos int

	for pos < len(S) {
		// skip redudant spaces
		if !unicode.IsSpace(rune(S[pos])) {
			buf.WriteByte(S[pos])
		}
		pos++
	}

	return buf.String()
}

func (root *Tree) Decide(animal []string) float64 {
	res := 1.0

	features := make(map[string]bool)
	for i := 0; i < len(animal); i++ {
		features[animal[i]] = true
	}

	node := root

	for {
		res *= node.weight
		if node.leaf {
			break
		}
		if features[node.feature] {
			node = node.yes
		} else {
			node = node.no
		}
	}

	return res
}

func ParseAsTree(S string) *Tree {
	tree := new(Tree)
	weight, pos := readWeight(S)
	tree.weight = weight
	if S[pos] == ')' {
		// read a leaf
		tree.leaf = true
		return tree
	}
	// read feature
	x := pos
	for pos < len(S) && S[pos] != '(' {
		pos++
	}
	tree.feature = S[x:pos]

	x = pos
	y := pair(S, pos)
	tree.yes = ParseAsTree(S[x : y+1])
	pos = y + 1
	x = pos
	y = pair(S, pos)
	tree.no = ParseAsTree(S[x : y+1])

	return tree
}

func pair(S string, from int) int {
	var level int
	for i := from; i < len(S); i++ {
		if S[i] == '(' {
			level++
		} else if S[i] == ')' {
			level--
		}
		if level == 0 {
			return i
		}
	}
	return -1
}

func readWeight(S string) (float64, int) {
	var pos int
	pos++
	var d int
	for pos < len(S) && unicode.IsDigit(rune(S[pos])) {
		d = d*10 + int(S[pos]-'0')
		pos++
	}
	if pos == len(S) || S[pos] == ')' {
		return float64(d), pos
	}
	//S[pos] == '.'
	var f int
	pos++
	base := 1.0

	for pos < len(S) && unicode.IsDigit(rune(S[pos])) {
		f = f*10 + int(S[pos]-'0')
		base *= 0.1
		pos++
	}
	res := float64(d) + float64(f)*base

	return res, pos
}
