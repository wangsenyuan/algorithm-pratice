package main

import (
	"testing"
)

func runSample(t *testing.T, n int, E [][]int, first, second int) {
	G := getGraph(n, E)
	dist1 := getDistanceFrom(n, G, first-1)
	dist2 := getDistanceFrom(n, G, second-1)

	ask := func(nodes []int) (int, int) {
		var node = -1
		var best int

		for _, x := range nodes {
			x--
			tmp := dist1[x] + dist2[x]
			if node < 0 || tmp < best {
				node = x + 1
				best = tmp
			}
		}
		return node, best
	}

	a, b := solve(n, E, ask)

	if first != a || second != b {
		t.Errorf("Sample expect %d %d, but got %d %d", first, second, a, b)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	E := [][]int{
		{1, 2},
		{1, 3},
	}
	first := 1
	second := 3
	runSample(t, n, E, first, second)
}
//
//func TestSample3(t *testing.T) {
//	content, error := ioutil.ReadFile("test3.in")
//	if error != nil {
//		t.FailNow()
//	}
//	reader := bufio.NewReader(strings.NewReader(string(content)))
//	tc := readNum(reader)
//
//	for tc > 0 {
//		tc--
//
//		n := readNum(reader)
//		E := make([][]int, n-1)
//		for i := 0; i < n-1; i++ {
//			E[i] = readNNums(reader, 2)
//		}
//
//		first := rand.Intn(n) + 1
//		second := first
//		for second == first {
//			second = rand.Intn(n) + 1
//		}
//		runSample(t, n, E, first, second)
//	}
//}
