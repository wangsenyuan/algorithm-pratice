package main

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func TestSample1(t *testing.T) {
	n := 5
	islands := [][]int{
		{0, 0},
		{1, 0},
		{2, 0},
		{0, -1},
		{0, -2},
	}

	res := solve(n, islands)

	fmt.Println("get solution: ")
	for _, ans := range res {
		fmt.Println(ans)
	}

	if len(res) != 1 {
		t.Errorf("this sample could have one island to light up, but get a solution with %d islands", len(res))
	} else {
		lightUp := make([]bool, n)
		lightUpIslands(n, islands, res[0], lightUp)
		for j := 0; j < n; j++ {
			if !lightUp[j] {
				t.Errorf("the solution should light up all islands, but %d is not", j)
				break
			}
		}
	}
}

func TestSample2(t *testing.T) {
	n := 4
	islands := [][]int{
		{5, 0},
		{-5, 0},
		{0, 5},
		{0, -5},
	}

	res := solve(n, islands)

	fmt.Println("get solution: ")
	for _, ans := range res {
		fmt.Println(ans)
	}

	if len(res) != 2 {
		t.Errorf("this sample could have two islands to light up, but get a solution with %d islands", len(res))
	} else {
		lightUp := make([]bool, n)
		lightUpIslands(n, islands, res[0], lightUp)
		lightUpIslands(n, islands, res[1], lightUp)
		for j := 0; j < n; j++ {
			if !lightUp[j] {
				t.Errorf("the solution should light up all islands, but %d is not", j)
				break
			}
		}
	}
}

func lightUpIslands(n int, islands [][]int, ans string, lightUp []bool) {
	ss := strings.Split(ans, " ")
	is, d := ss[0], ss[1]
	i, _ := strconv.Atoi(is)
	for j := 0; j < n; j++ {
		switch d {
		case "SW":
			if islands[j][0] <= islands[i][0] && islands[j][1] <= islands[i][1] {
				lightUp[j] = true
			}
		case "NW":
			if islands[j][0] <= islands[i][0] && islands[j][1] >= islands[i][1] {
				lightUp[j] = true
			}
		case "SE":
			if islands[j][0] >= islands[i][0] && islands[j][1] <= islands[i][1] {
				lightUp[j] = true
			}
		case "NE":
			if islands[j][0] >= islands[i][0] && islands[j][1] >= islands[i][1] {
				lightUp[j] = true
			}
		}
	}
}
