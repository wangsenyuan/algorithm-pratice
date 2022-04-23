package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, m := readTwoNums(reader)

	A := readNNums(reader, n)
	posts := make([]string, m)

	for i := 0; i < m; i++ {
		posts[i] = readString(reader)
	}

	res := solve(A, posts)

	for _, cur := range res {
		fmt.Println(cur)
	}
}

func solve(A []int, posts []string) []string {
	special := make(map[int]bool)

	for _, id := range A {
		special[id] = true
	}

	var la []Post
	var lb []Post

	for _, post := range posts {
		cur := ParseAsPost(post)
		if special[cur.friend] {
			la = append(la, cur)
		} else {
			lb = append(lb, cur)
		}
	}

	sort.Sort(Posts(la))
	sort.Sort(Posts(lb))

	var res []string

	for _, l := range la {
		res = append(res, l.content)
	}

	for _, l := range lb {
		res = append(res, l.content)
	}
	return res
}

func readUint64(bytes []byte, from int, val *uint64) int {
	i := from

	var tmp uint64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + uint64(bytes[i]-'0')
		i++
	}
	*val = tmp

	return i
}

func readInt(bytes []byte, from int, val *int) int {
	i := from
	sign := 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	tmp := 0
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' {
			return s[:i]
		}
	}
	return s
}

func readNum(reader *bufio.Reader) (a int) {
	bs, _ := reader.ReadBytes('\n')
	readInt(bs, 0, &a)
	return
}

func readTwoNums(reader *bufio.Reader) (a int, b int) {
	res := readNNums(reader, 2)
	a, b = res[0], res[1]
	return
}

func readThreeNums(reader *bufio.Reader) (a int, b int, c int) {
	res := readNNums(reader, 3)
	a, b, c = res[0], res[1], res[2]
	return
}

func readNNums(reader *bufio.Reader, n int) []int {
	res := make([]int, n)
	x := 0
	bs, _ := reader.ReadBytes('\n')
	for i := 0; i < n; i++ {
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') && bs[x] != '-' {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
	return res
}

type Post struct {
	friend  int
	pop     int
	content string
}

func ParseAsPost(s string) Post {
	buf := []byte(s)
	var f, p int
	pos := readInt(buf, 0, &f)
	pos = readInt(buf, pos+1, &p)
	content := s[pos+1:]
	return Post{f, p, content}
}

type Posts []Post

func (ps Posts) Len() int {
	return len(ps)
}

func (ps Posts) Less(i, j int) bool {
	return ps[i].pop > ps[j].pop
}

func (ps Posts) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}
