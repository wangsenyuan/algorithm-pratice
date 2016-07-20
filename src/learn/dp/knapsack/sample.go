package main

import (
	"fmt"
	"sort"
)

func main() {
	ws := []int{5, 10, 20}
	vs := []int{50, 60, 140}
	v := dp(ws, vs, 30)
	fmt.Printf("%d\n", v)
}

func dp(ws []int, vs []int, weight int) int {
	n := len(ws)

	fx := make([]int, weight+1)
	//f[i][w] = max(f[i][w], f[i - 1][w - wi] + vi| j from 1 until i)

	for i := 1; i <= n; i++ {
		for w := weight; w >= ws[i-1]; w-- {
			fx[w] = max(fx[w], fx[w-ws[i-1]]+vs[i-1])
		}
	}

	return fx[weight]
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func backTrack(ws []int, vs []int, W int) int {
	ws, vs = sortByUnitValue(ws, vs)

	var knapsack func(i int, profit int, weight int)

	maxProfit := 0

	promising := func(i int, profit int, weight int) bool {
		if weight >= W {
			return false
		}

		total := weight
		bound := float64(profit)
		j := i + 1

		for ; j < len(ws) && total+ws[j] <= W; j++ {
			total += ws[j]
			bound += float64(vs[j])
		}

		if j < len(ws) {
			bound += float64(W-total) * float64(vs[j]) / float64(ws[j])
		}
		return bound > float64(maxProfit)
	}

	knapsack = func(i int, profit int, weight int) {
		if weight <= W && profit > maxProfit {
			maxProfit = profit
		}

		if promising(i, profit, weight) {
			knapsack(i+1, profit+vs[i+1], weight+ws[i+1])
			knapsack(i+1, profit, weight)
		}
	}

	knapsack(-1, 0, 0)

	return maxProfit
}

func sortByUnitValue(ws []int, vs []int) ([]int, []int) {
	n := len(ws)
	rs := make([]unitValue, n)
	for i := 0; i < n; i++ {
		rs[i] = unitValue{ws[i], vs[i]}
	}
	sort.Sort(byUnitValue(rs))

	as := make([]int, n)
	bs := make([]int, n)

	for i, r := range rs {
		as[i] = r.w
		bs[i] = r.v
	}

	return as, bs
}

type unitValue struct {
	w int
	v int
}

type byUnitValue []unitValue

func (a byUnitValue) Len() int {
	return len(a)
}

func (a byUnitValue) Less(i, j int) bool {
	x := float64(a[i].v) / float64(a[i].w)
	y := float64(a[j].v) / float64(a[j].w)
	return x < y
}

func (a byUnitValue) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
