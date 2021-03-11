package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	s, _ := reader.ReadString('\n')
	n, _ := strconv.Atoi(s[:len(s)-1])
	res := solve(n, reader)
	fmt.Print(res)
}

const productDimSize = 3

const locationDimSize = 4

const sexDimSize = 1

const ageDimSize = 90

var productLevelSize = []int{1, 10, 3}
var locationLevelSize = []int{1, 10, 20, 5}

var sexLevelSz = []int{2}
var purchaseData [][][][][][][]int

var entitySize [][]int

func init() {
	psz := 1
	purchaseData = make([][][][][][][]int, productDimSize)
	for i := 0; i < productDimSize; i++ {
		lsz := 1
		psz *= productLevelSize[i]
		purchaseData[i] = make([][][][][][]int, locationDimSize)
		for j := 0; j < locationDimSize; j++ {
			ssz := 1
			lsz *= locationLevelSize[j]
			purchaseData[i][j] = make([][][][][]int, sexDimSize)
			for k := 0; k < sexDimSize; k++ {
				ssz *= sexLevelSz[k]
				purchaseData[i][j][k] = make([][][][]int, psz)
				for a := 0; a < psz; a++ {
					purchaseData[i][j][k][a] = make([][][]int, lsz)
					for b := 0; b < lsz; b++ {
						purchaseData[i][j][k][a][b] = make([][]int, ssz)
						for c := 0; c < ssz; c++ {
							purchaseData[i][j][k][a][b][c] = make([]int, ageDimSize+1)
						}
					}
				}
			}
		}
	}
	entitySize = make([][]int, 2)
	entitySize[0] = productLevelSize
	entitySize[1] = locationLevelSize
}

func solve(n int, reader *bufio.Reader) string {
	var buf bytes.Buffer
	for n > 0 {
		n--
		s, _ := reader.ReadString('\n')
		s = s[:len(s)-1]
		if s[0] == 'I' {
			input(s[2:])
		} else {
			tmp := query(s[2:])
			buf.WriteString(fmt.Sprintf("%d\n", tmp))
		}
	}
	return buf.String()
}

func input(s string) {
	ss := strings.Split(s, " ")
	products := getEntity(ss[0], 0)
	locations := getEntity(ss[1], 1)
	var sl int
	var sv int
	if ss[2] == "M" {
		sv = 1
	}
	age, _ := strconv.Atoi(ss[3])
	unitSold, _ := strconv.Atoi(ss[4])

	for i := 0; i < len(products); i++ {
		pv := products[i].first
		pl := products[i].second
		for j := 0; j < len(locations); j++ {
			lv := locations[j].first
			ll := locations[j].second
			bitUpdate(purchaseData[pl][ll][sl][pv][lv][sv], ageDimSize, age, unitSold)
		}
	}

}

func query(s string) int {
	ss := strings.Split(s, " ")
	products := getEntity(ss[0], 0)
	pv := last(products).first
	pl := last(products).second

	locations := getEntity(ss[1], 1)
	lv := last(locations).first
	ll := last(locations).second

	var sl int
	var sv int
	if ss[2] == "M" {
		sv = 1
	}

	var ageStart, ageEnd int
	if strings.Contains(ss[3], "-") {
		xx := strings.Split(ss[3], "-")
		ageStart, _ = strconv.Atoi(xx[0])
		ageEnd, _ = strconv.Atoi(xx[1])
	} else {
		ageStart, _ = strconv.Atoi(ss[3])
	}

	res := bitQuery(purchaseData[pl][ll][sl][pv][lv][sv], ageDimSize, ageEnd)
	res -= bitQuery(purchaseData[pl][ll][sl][pv][lv][sv], ageDimSize, ageStart-1)
	return res
}

func bitUpdate(arr []int, sz int, idx int, val int) {
	for idx <= sz {
		arr[idx] += val
		idx += idx & -idx
	}
}

func bitQuery(arr []int, sz int, idx int) int {
	var res int
	for idx > 0 {
		res += arr[idx]
		idx -= idx & -idx
	}
	return res
}

func getEntity(s string, entityId int) []Pair {
	var level, val int
	sz := entitySize[entityId][level]
	var res []Pair
	res = append(res, Pair{val, level})
	level++

	data := strings.Split(s, ".")
	for i := 0; i < len(data); i++ {
		id, _ := strconv.Atoi(data[i])
		if id < 0 {
			break
		}
		val = val + id*sz
		res = append(res, Pair{val, level})

		sz *= entitySize[entityId][level]
		level++
	}
	return res
}

type Pair struct {
	first, second int
}

func last(arr []Pair) Pair {
	return arr[len(arr)-1]
}
