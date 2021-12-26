package p2120

import (
	"reflect"
	"sort"
	"testing"
)

func runSample(t *testing.T, recipes []string, ingredients [][]string, supplies []string, expect []string) {
	res := findAllRecipes(recipes, ingredients, supplies)

	sort.Strings(expect)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	recipes := []string{"bread"}
	ingredients := [][]string{{"yeast", "flour"}}
	supplies := []string{"yeast", "flour", "corn"}
	expect := []string{"bread"}
	runSample(t, recipes, ingredients, supplies, expect)
}

func TestSample2(t *testing.T) {
	recipes := []string{"bread", "sandwich"}
	ingredients := [][]string{{"yeast", "flour"}, {"bread", "meat"}}
	supplies := []string{"yeast", "flour", "meat"}
	expect := []string{"bread", "sandwich"}
	runSample(t, recipes, ingredients, supplies, expect)
}

func TestSample3(t *testing.T) {
	recipes := []string{"bread", "sandwich", "burger"}
	ingredients := [][]string{{"yeast", "flour"}, {"bread", "meat"}, {"sandwich", "meat", "bread"}}
	supplies := []string{"yeast", "flour", "meat"}
	expect := []string{"bread", "sandwich", "burger"}
	runSample(t, recipes, ingredients, supplies, expect)
}

func TestSample4(t *testing.T) {
	recipes := []string{"bread"}
	ingredients := [][]string{{"yeast", "flour"}}
	supplies := []string{"yeast"}
	
	runSample(t, recipes, ingredients, supplies, nil)
}
