package main

import (
	"bufio"
	"os"
	"reflect"
	"testing"
)

func runSample(t *testing.T, A []int, B [][]int, C [][]int, expect []int) {
	res := solve(A, B, C)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{50, 40, 30, 20, 10}
	B := [][]int{
		{1, 2, 3, 100},
		{4},
		{6, 7},
	}
	C := [][]int{
		{1, 3, 300},
		{2, 1, 400},
		{2, 1, 5},
	}
	expect := []int{1, 0, 1}
	runSample(t, A, B, C, expect)
}

func TestSampleData(t *testing.T) {

	dir, _ := os.Getwd()

	dat, _ := os.Open(dir + "/sample/data.in")

	defer dat.Close()

	reader := bufio.NewReader(dat)

	res := process(reader)

	ans, _ := os.Open(dir + "/sample/data.ans")
	defer ans.Close()

	ansReader := bufio.NewReader(ans)

	for i := 0; i < len(res); i++ {
		expect := readNum(ansReader)
		if res[i] != expect {
			t.Fatalf("Sample expect %d, but got %d[@%d]", expect, res[i], i)
		}
	}

}
