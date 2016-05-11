package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	line, _ := reader.ReadString('\n')
	t, _ := strconv.Atoi(strings.Trim(line, "\n "))

	for i := 1; i <= t; i++ {
		line, _ = reader.ReadString('\n')
		n, _ := strconv.Atoi(strings.Trim(line, "\n"))
		topics := make([]topic, 0, n)
		for j := 0; j < n; j++ {
			line, _ = reader.ReadString('\n')
			words := strings.Split(line, " ")
			topic := topic{words[0], words[1]}
			topics = append(topics, topic)
		}
		r := play(topics)
		fmt.Printf("Case #%d: %d\n", i, r)
	}
}

type topic struct {
	first, second string
}

const INF = math.MaxInt64

func play(topics []topic) int {
	g := &graph{make(map[string][]string), make(map[string][]string),
		make(map[string]string), make(map[string]string), make(map[string]int)}

	for _, t := range topics {
		g.Connect(t.first, t.second)
	}

	return len(topics) - g.HK()
}

type graph struct {
	U    map[string][]string
	V    map[string][]string
	pu   map[string]string
	pv   map[string]string
	dist map[string]int
}

func (g *graph) Connect(u, v string) {
	g.U[u] = append(g.U[u], v)
	g.V[v] = append(g.V[v], u)
}

func (g *graph) HK() int {
	for u, _ := range g.U {
		g.pu[u] = ""
	}

	for v, _ := range g.V {
		g.pv[v] = ""
	}

	matching := 0
	for g.bfs() {
		//fmt.Println("call dfs()")
		for u, _ := range g.U {
			if g.pu[u] == "" && g.dfs(u) {
				matching++
			}
		}
	}

	return matching + (len(g.U) - countPair(g.pu)) + (len(g.V) - countPair(g.pv))
}

func countPair(p map[string]string) int {
	cnt := 0
	for _, v := range p {
		if v != "" {
			cnt += 1
		}
	}
	return cnt
}

type queue []string

func enqueue(q queue, elem string) queue {
	return queue(append(q, elem))
}

func dequeue(q queue) (queue, string) {
	if len(q) == 0 {
		return q, ""
	}

	return q[1:], q[0]
}

func (g *graph) bfs() bool {
	q := queue(make([]string, 0, len(g.U)))

	for u, _ := range g.U {
		if g.pu[u] == "" {
			g.dist[u] = 0
			q = enqueue(q, u)
		} else {
			g.dist[u] = INF
		}
	}

	g.dist[""] = INF

	var u = ""
	for len(q) > 0 {
		//fmt.Printf("%s in queue\n", q)
		q, u = dequeue(q)
		for _, v := range g.U[u] {
			w := g.pv[v]
			if g.dist[w] == INF {
				g.dist[w] = g.dist[u] + 1
				q = enqueue(q, w)
			}
		}
	}
	//fmt.Printf("bfs() => %d %b\n", g.dist[""], g.dist[""] < INF)

	return g.dist[""] < INF
}

func (g *graph) dfs(u string) bool {
	if u == "" {
		return true
	}

	for _, v := range g.U[u] {
		w := g.pv[v]
		if g.dist[w] == g.dist[u]+1 && g.dfs(w) {
			g.pu[u] = v
			g.pv[v] = u
			//fmt.Printf("%s -> %s\n", u, v)
			return true
		}
	}
	g.dist[u] = INF
	return false
}
