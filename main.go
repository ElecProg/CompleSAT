package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	if len(os.Args)-1 != 1 {
		os.Stderr.WriteString(fmt.Sprintln("expecting 1 argument", len(os.Args)-1, "given"))
		os.Exit(1)
	}

	problem, err := readDimacsCnfFile(os.Args[1])

	if err != nil {
		os.Stderr.WriteString(fmt.Sprintln(err.Error()))
		os.Exit(1)
	}

	t0 := time.Now()
	solution := Solve(problem)
	os.Stderr.WriteString(fmt.Sprintln("Answer found after", time.Now().Sub(t0)))

	if solution.Unsatisfiable {
		fmt.Println("UNSAT")

	} else {
		fmt.Println("SAT")

		for i := 1; i <= len(solution.Assigned); i++ {
			if solution.Assigned[i] {
				fmt.Print(i, " ")

			} else {
				fmt.Print(-i, " ")
			}
		}

		fmt.Println("0")
	}
}
