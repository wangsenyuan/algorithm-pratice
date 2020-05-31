package p1460

import (
	"reflect"
	"sort"
)

func canBeEqual(target []int, arr []int) bool {
	sort.Ints(target)
	sort.Ints(arr)

	return reflect.DeepEqual(target, arr)
}
