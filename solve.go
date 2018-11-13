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

// Solve is the solver itself
func Solve(problem Problem) Problem {
	// This is a wrapper around solve to provide memory safety to the end user
	return solve(problem.Copy())
}

// Solve is actually the real solver
func solve(problem Problem) Problem {
	// No need to create copy of problem, the caller should have done that.

	// Literal elimination
	for variable := range problem.Variables {
		if _, exists := problem.Variables[-variable]; !exists {
			// The other polarity does not exist, we can eliminate the literal
			problem.Assign(variable)
		}
	}

	// Unit propagation
	for _, vrs := range problem.Clauses {
		if len(vrs) == 1 {
			for vr := range vrs {
				problem.Assign(vr)
			}
		}
	}

	// We are done if there are no more variables to assign or
	// the current assignment is unsatisfiable.
	if len(problem.Variables) == 0 || problem.Unsatisfiable {
		return problem
	}

	vr := ChooseVariable(problem)

	// The solver changes the problem directly and we need the
	// current version if this try is unsatisfiable.
	firstTry := problem.Copy()

	firstTry.Assign(vr)
	res := Solve(firstTry)

	if !res.Unsatisfiable {
		return res
	}

	// No need to create a copy, we'll never use problem again
	problem.Assign(-vr)
	return Solve(problem)
}
