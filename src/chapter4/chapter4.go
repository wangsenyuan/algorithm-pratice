// chapter4
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World!")
	//var buffer [20]byte
	var grid1 [3][3]int
	grid1[1][0], grid1[1][1], grid1[1][2] = 8, 6, 2
	//grid2 := [3][3]int{{4, 3}, {8, 6, 2}}
	//cites := [...]string{"Shanghai", "Mumbai", "Beijing"}

	massForPlanet := make(map[string]float64)

	massForPlanet["Mercury"] = 0.06
	massForPlanet["Venus"] = 0.82
	massForPlanet["Earth"] = 1.00
	massForPlanet["Mars"] = 0.11

	fmt.Println(massForPlanet)
}
