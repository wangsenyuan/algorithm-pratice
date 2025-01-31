package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	if res {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
}

func process(reader *bufio.Reader) bool {
	n, m, k := readThreeNums(reader)
	doors := make([][]int, m)
	for i := range m {
		doors[i] = readNNums(reader, 2)
	}
	thursday := make([]string, k)
	for i := range k {
		thursday[i] = readString(reader)
	}
	friday := make([]string, k)
	for i := range k {
		friday[i] = readString(reader)
	}
	return solve(n, k, doors, thursday, friday)
}

func solve(n int, k int, doors [][]int, thursday []string, friday []string) bool {
	// 必须知道A是否能够到到房间B
	// 除此之外，还需要确定，它们使用的key是一致的
	a1, f1, k1 := prepare(n, k, doors, thursday)
	a2, f2, k2 := prepare(n, k, doors, friday)

	var find func(f []int, x int) int

	find = func(f []int, x int) int {
		if f[x] != x {
			f[x] = find(f, f[x])
		}
		return f[x]
	}

	for i := 0; i < n; i++ {
		j := find(f1, i)
		if i == j {
			if find(f2, i) != j {
				return false
			}
			if len(k1[j]) != len(k2[j]) {
				return false
			}
			for x := range k1[j] {
				if !k2[j][x] {
					return false
				}
			}
		}
	}

	id := make(map[string]int)
	for i, x := range a1 {
		id[x.name] = i
	}

	for _, x := range a2 {
		y := a1[id[x.name]]
		u, v := x.room, y.room
		if find(f1, u) != find(f2, v) {
			return false
		}
	}
	return true
}

type person struct {
	name string
	room int
	keys []int
}

func readPerson(s string) person {
	buf := []byte(s)
	var j int
	for j < len(buf) && buf[j] != ' ' {
		j++
	}
	name := string(buf[:j])
	var room, m int
	j = readInt(buf, j+1, &room)
	j = readInt(buf, j+1, &m)
	keys := make([]int, m)
	for u := range m {
		j = readInt(buf, j+1, &keys[u])
		keys[u]--
	}
	return person{name, room - 1, keys}
}

func prepare(n int, k int, doors [][]int, people []string) (arr []person, fa []int, keys []map[int]bool) {
	arr = make([]person, k)
	id := make(map[string]int)
	for i := range k {
		p := readPerson(people[i])
		id[p.name] = i
		arr[i] = p
	}

	keys = make([]map[int]bool, n)
	for i := range n {
		keys[i] = make(map[int]bool)
	}

	for _, cur := range arr {
		r := cur.room
		for _, j := range cur.keys {
			// 在房间r中有钥匙j
			keys[r][j] = true
		}
	}

	fa = make([]int, n)
	for i := 0; i < n; i++ {
		fa[i] = i
	}

	var find func(x int) int

	find = func(x int) int {
		if fa[x] != x {
			fa[x] = find(fa[x])
		}
		return fa[x]
	}

	merge := func(u int, v int) {
		if u > v {
			u, v = v, u
		}
		fa[v] = u
		for j := range keys[v] {
			keys[u][j] = true
		}
	}

	for {
		ok := false
		for i, door := range doors {
			u, v := door[0], door[1]
			u--
			v--
			u = find(u)
			v = find(v)
			if u != v && (keys[u][i] || keys[v][i]) {
				ok = true
				merge(u, v)
			}
		}
		// 这里至少连接了两个房间，所以最多循环n次
		if !ok {
			break
		}
	}

	return
}
