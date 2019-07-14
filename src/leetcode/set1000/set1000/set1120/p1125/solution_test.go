package p1125

import (
	"reflect"
	"sort"
	"testing"
)

func runSample(t *testing.T, skills []string, people [][]string, expect []int) {
	res := smallestSufficientTeam(skills, people)
	sort.Ints(expect)
	sort.Ints(res)
	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %v %v, expect %v, but got %v", skills, people, expect, res)
	}
}

func TestSample1(t *testing.T) {
	skills := []string{"java", "nodejs", "reactjs"}
	people := [][]string{
		{"java"}, {"nodejs"}, {"nodejs", "reactjs"},
	}
	expect := []int{0, 2}
	runSample(t, skills, people, expect)
}

func TestSample2(t *testing.T) {
	skills := []string{"algorithms", "math", "java", "reactjs", "csharp", "aws"}
	people := [][]string{
		{"algorithms", "math", "java"},
		{"algorithms", "math", "reactjs"},
		{"java", "csharp", "aws"},
		{"reactjs", "csharp"},
		{"csharp", "math"},
		{"aws", "java"},
	}
	expect := []int{1, 2}
	runSample(t, skills, people, expect)
}
