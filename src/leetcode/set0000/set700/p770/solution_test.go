package p770

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, expr string, evalvars []string, evalints []int, expect []string) {
	res := basicCalculatorIV(expr, evalvars, evalints)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("sample %s %v %v, expect %v, but got %v", expr, evalvars, evalints, expect, res)
	}
}

func TestSample1(t *testing.T) {
	expression := "e + 8 - a + 5"
	evalvars := []string{"e"}
	evalints := []int{1}
	expect := []string{"-1*a", "14"}
	runSample(t, expression, evalvars, evalints, expect)
}

func TestSample2(t *testing.T) {
	expression := "e - 8 + temperature - pressure"
	evalvars := []string{"e", "temperature"}
	evalints := []int{1, 12}
	expect := []string{"-1*pressure", "5"}
	runSample(t, expression, evalvars, evalints, expect)
}

func TestSample3(t *testing.T) {
	expression := "(e + 8) * (e - 8)"
	evalvars := []string{}
	evalints := []int{}
	expect := []string{"1*e*e", "-64"}
	runSample(t, expression, evalvars, evalints, expect)
}

func TestSample4(t *testing.T) {
	expression := "7 - 7"
	evalvars := []string{}
	evalints := []int{}
	expect := []string{}
	runSample(t, expression, evalvars, evalints, expect)
}

func TestSample5(t *testing.T) {
	expression := "a * b * c + b * a * c * 4"
	evalvars := []string{}
	evalints := []int{}
	expect := []string{"5*a*b*c"}
	runSample(t, expression, evalvars, evalints, expect)
}

func TestSample6(t *testing.T) {
	expression := "((a - b) * (b - c) + (c - a)) * ((a - b) + (b - c) * (c - a))"
	evalvars := []string{}
	evalints := []int{}
	expect := []string{"-1*a*a*b*b", "2*a*a*b*c", "-1*a*a*c*c", "1*a*b*b*b", "-1*a*b*b*c", "-1*a*b*c*c", "1*a*c*c*c", "-1*b*b*b*c", "2*b*b*c*c", "-1*b*c*c*c", "2*a*a*b", "-2*a*a*c", "-2*a*b*b", "2*a*c*c", "1*b*b*b", "-1*b*b*c", "1*b*c*c", "-1*c*c*c", "-1*a*a", "1*a*b", "1*a*c", "-1*b*c"}
	runSample(t, expression, evalvars, evalints, expect)
}

func TestCompareSlice(t *testing.T) {
	a := []string{"a", "b", "b"}
	b := []string{"a", "a", "c"}
	if compareSlice(a, b) {
		t.Errorf("%v should less than %v, but got not", a, b)
	}
}
