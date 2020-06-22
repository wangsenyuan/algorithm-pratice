package p1487

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, names []string, expect []string) {
	res := getFolderNames(names)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %v, expect %v, but got %v", names, expect, res)
	}
}

func TestSample1(t *testing.T) {
	names := []string{"pes", "fifa", "gta", "pes(2019)"}
	expect := []string{"pes", "fifa", "gta", "pes(2019)"}
	runSample(t, names, expect)
}

func TestSample2(t *testing.T) {
	names := []string{"gta", "gta(1)", "gta", "avalon"}
	expect := []string{"gta", "gta(1)", "gta(2)", "avalon"}
	runSample(t, names, expect)
}

func TestSample3(t *testing.T) {
	names := []string{"onepiece", "onepiece(1)", "onepiece(2)", "onepiece(3)", "onepiece"}
	expect := []string{"onepiece", "onepiece(1)", "onepiece(2)", "onepiece(3)", "onepiece(4)"}
	runSample(t, names, expect)
}

func TestSample4(t *testing.T) {
	names := []string{"kaido", "kaido(1)", "kaido", "kaido(1)"}
	expect := []string{"kaido", "kaido(1)", "kaido(2)", "kaido(1)(1)"}
	runSample(t, names, expect)
}
