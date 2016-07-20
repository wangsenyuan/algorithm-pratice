package main

import "fmt"

func main() {
	t := 0
	fmt.Scanf("%d", &t)

	for i := 0; i < t; i++ {
		n := 0
		fmt.Scanf("\n%d", &n)
		nodes := make([][]int, n)

		for j := 0; j < n; j++ {
			nodes[j] = make([]int, 3)
			a := 0
			fmt.Scanf("%d", &a)
			nodes[j][0] = a
			if a > 0 {
				fmt.Scanf("%d %d", &nodes[j][1], &nodes[j][2])
				nodes[j][1]--
				nodes[j][2]--
			}
		}

		r := play(nodes)

		fmt.Printf("%.5f\n", r)
	}
}

func play(nodes [][]int) float64 {
	//fmt.Printf("%v\n", nodes)

	n := len(nodes)

	ps := make([]float64, n)
	left, right := 0.0, 1.0
	res := 0.0
	loop := 0
	for ps[n-1] != 0.5 {
		res = (left + right) / 2.0
		for i := 0; i < n; i++ {
			node := nodes[i]
			if node[0] == 0 {
				ps[i] = res
				continue
			}

			if node[0] == 1 {
				ps[i] = ps[node[1]] + (1-ps[node[1]])*ps[node[2]]
				continue
			}

			ps[i] = ps[node[1]] * ps[node[2]]
		}

		if ps[n-1] > 0.5 {
			right = res
		} else {
			left = res
		}
		loop++
		if loop > 30 {
			break
		}
	}

	return res
}
