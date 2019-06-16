package main

import "fmt"

func main() {
	equations := [][]string{[]string{"a", "b"}, []string{"b", "c"}, []string{"c", "d"}}
	values := []float64{2.0, 3.0, 4.0}
	//queries := [][]string{[]string{"a", "c"}, []string{"b", "a"}, []string{"a", "e"}, []string{"a", "a"}, []string{"x", "x"}}
	queries := [][]string{[]string{"a", "e"}}

	result := calcEquation(equations, values, queries)

	for i := range result {
		fmt.Println(result[i])
	}
}

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	graph := make(map[string]map[string]float64)

	updateGraph := func(a string, b string, value float64) {
		as := graph[a]
		if as == nil {
			as = make(map[string]float64)
		}
		as[b] = value
		graph[a] = as
	}
	for i := range equations {
		a, b := equations[i][0], equations[i][1]
		value := values[i]
		updateGraph(a, b, value)
		updateGraph(b, a, 1.0/value)
	}

	var query func(a, b string, path map[string]bool) float64

	query = func(a, b string, path map[string]bool) float64 {
		res := -1.0
		path[a] = true
		as := graph[a]
		if as != nil {
			if a == b {
				res = 1.0
			} else {
				if v, has := as[b]; has {
					res = v
				} else {
					for c, v := range as {
						if path[c] {
							continue
						}
						w := query(c, b, path)
						if w != -1.0 {
							res = v * w
							break
						}
					}
				}
			}
		}

		if res != -1.0 {
			updateGraph(a, b, res)
			updateGraph(b, a, 1.0/res)
		}
		path[a] = false

		return res
	}
	result := make([]float64, len(queries))
	for i := range queries {
		result[i] = query(queries[i][0], queries[i][1], make(map[string]bool))
	}
	return result
}
