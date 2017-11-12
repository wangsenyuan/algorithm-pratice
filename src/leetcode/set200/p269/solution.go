package main

import "fmt"

func main() {
	words := []string{"wrt", "wrf", "er", "ett", "rftt"}
	//words := []string{"a", "b", "a"}
	//words := []string{"z", "z"}
	//words := []string{"wrtkj", "wrt"}
	fmt.Println(alienOrder(words))
}

func alienOrder(words []string) (result string) {
	defer func() {
		if r := recover(); r != nil {
			result = ""
		}
	}()

	graph := make(map[byte][]byte)
	order(words, graph)

	d := make(map[byte]int)

	for x, bs := range graph {
		if _, ok := d[x]; !ok {
			d[x] = 0
		}
		for _, b := range bs {
			d[b]++
		}
	}
	var res []byte
	for len(d) > 0 {
		var v byte = ' '
		for b, c := range d {
			if c == 0 {
				v = b
				delete(d, b)
				break
			}
		}
		if v == ' ' {
			return ""
		}
		res = append(res, v)
		for _, w := range graph[v] {
			d[w]--
		}
	}

	result = string(res)
	return
}

func order(words []string, g map[byte][]byte) {
	var left []string
	for i := 0; i < len(words); i++ {
		word := words[i]
		a := word[0]
		if len(g[a]) == 0 {
			g[a] = make([]byte, 0)
		}
		b := word[1:]
		if i > 0 && a == words[i-1][0] {
			if len(b) > 0 {
				left = append(left, b)
			} else if len(left) > 0 {
				panic("invalid input")
			}
			continue
		}
		if len(left) > 0 {
			order(left, g)
		}
		left = make([]string, 0)
		if len(b) > 0 {
			left = append(left, b)
		}
		if i == 0 {
			continue
		}

		x := words[i-1][0]
		g[x] = append(g[x], a)
	}

	if len(left) > 0 {
		order(left, g)
	}
}
