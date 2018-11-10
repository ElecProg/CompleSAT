package main

// ChooseVariable selects a variable which is currently most promissing
func ChooseVariable(problem Problem) int {
	// Find the variable phase affecting the most clauses
	bestVariable, bestClauses := 0, 0
	for variable, clauses := range problem.Variables {
		if len(clauses) > bestClauses {
			bestVariable, bestClauses = variable, len(clauses)
		}
	}

	return bestVariable
}

// Solve is the solver itself, takes solution spaces in pairs and combines them
// it finds the entire solution space
func Solve(problem Problem) Problem {
	// Literal elimination
	for variable := range problem.Variables {
		if _, exists := problem.Variables[-variable]; !exists {
			// The other polarity does not exist, we can eliminate the literal
			problem = problem.Assign(variable)
		}
	}

	// Unit propagation
	for _, vrs := range problem.Clauses {
		if len(vrs) == 1 {
			for vr := range vrs {
				problem = problem.Assign(vr)
			}
		}
	}

	// We are done if there are no more variables to assign or
	// the current assignment is unsatisfiable.
	if len(problem.Variables) == 0 || problem.Unsatisfiable {
		return problem
	}

	vr := ChooseVariable(problem)
	res := Solve(problem.Assign(vr))

	if !res.Unsatisfiable {
		return res
	}

	return Solve(problem.Assign(-vr))
}
