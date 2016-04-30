// poolar2cartesian
package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"runtime"
)

type polar struct {
	radius float64
	ø      float64
}

type cartesian struct {
	x float64
	y float64
}

var prompt = "Enter a radius and angle (in degrees), e.g. 12.5 90, or %s to quit."

func init() {
	if runtime.GOOS == "windows" {
		prompt = fmt.Sprintf(prompt, "Ctrl + Z, Enter")
	} else {
		prompt = fmt.Sprintf(prompt, "Ctrl + D")
	}
}

func main() {
	questions := make(chan polar)
	defer close(questions)

	answers := createSolver(questions)
	defer close(answers)

	interact(questions, answers)
}

func createSolver(questions chan polar) chan cartesian {
	answers := make(chan cartesian)

	go func() {
		for {
			polarCoord := <-questions
			ø := polarCoord.ø * math.Pi / 180.0
			x := polarCoord.radius * math.Cos(ø)
			y := polarCoord.radius * math.Sin(ø)
			answers <- cartesian{x, y}
		}
	}()

	return answers
}

const result = "Polar radius =%.02f ø=%.02fº -> Cartesian x = %.02f y = %.02f\n"

func interact(questions chan polar, answers chan cartesian) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(prompt)
	for {
		fmt.Println("Radius and angle: ")
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		var radius, ø float64
		if _, err := fmt.Sscanf(line, "%f %f", &radius, &ø); err != nil {
			fmt.Println(os.Stderr, "invalid input")
			continue
		}
		questions <- polar{radius, ø}
		coord := <-answers
		fmt.Printf(result, radius, ø, coord.x, coord.y)
	}
	fmt.Println()
}
