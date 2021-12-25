package main

import "testing"

func runSample(t *testing.T, n int) {
	ans := solve(n)
	plates := ans.plates
	for i := 0; i < len(plates); i++ {
		if plates[i] > n {
			t.Fatalf("Sample result %v, has more plates than %d", plates, n)
		}
	}
	has := make([]map[int]int, len(plates))
	for i := 0; i < len(plates); i++ {
		has[i] = make(map[int]int)
	}
	for i := 1; i <= n; i++ {
		cur := ans.dayPuts[i-1]
		if len(cur) > 4 {
			t.Fatalf("Sample result %v, has more than 4 fruits at day %d", cur, i)
		}
		for j := 0; j < len(cur); j++ {
			p, f := cur[j].plate, cur[j].fruit
			if f > i {
				t.Fatalf("Sample result %v, fruit %d, not available at day %d yet", cur, f, i)
			}
			has[p-1][f]++
			// let find some plates that has [1..i] fruits
			var ok bool
			for k := 0; k < len(has) && !ok; k++ {
				if len(has[k]) < (i+1)/2 {
					continue
				}
				ii := 1
				for ii <= i && has[k][ii] > 0 {
					ii++
				}
				ok = ii > i
			}
			if !ok {
				t.Fatalf("not found a plate that has [1..%d] at day %d", i, i)
			}
		}
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 1)
}

func TestSample2(t *testing.T) {
	runSample(t, 3)
}

func TestSample3(t *testing.T) {
	runSample(t, 10)
}

func TestSample4(t *testing.T) {
	runSample(t, 1000)
}
