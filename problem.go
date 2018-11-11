package main

// Problem represents a problem
// Note that we store the two phases of a variable separately.
type Problem struct {
	// Mapping from a clause to a list of it's variables,
	// if the variable is positive/negative it expects true/false
	Clauses map[int]map[int]bool

	// Mapping from a variable phase to the clauses containing it
	Variables map[int]map[int]bool

	// Mapping of a variable to it's current assignment
	Assignment map[int]bool

	// A flag showing wheither the current assignment is UNSAT
	Unsatisfiable bool
}

// NewProblem creates a new problem
func NewProblem() Problem {
	return Problem{
		Clauses:       map[int]map[int]bool{},
		Variables:     map[int]map[int]bool{},
		Assignment:    map[int]bool{},
		Unsatisfiable: false,
	}
}

// AddClause adds a clause to the problem, updating the Variables and Clauses maps
// this does not update Unsatisfiable, would be fun if it would be that easy though...
func (p *Problem) AddClause(clause []int) {
	// Build a set of the variables in the clause
	// checking for tautologies.
	cl := map[int]bool{}
	for _, vr := range clause {
		if cl[-vr] {
			// Other phase of variable already in clause
			// we have a tautology
			return
		}

		// Else add the variable to the clause
		cl[vr] = true
	}

	// No tautologies found
	// Add the contraint to the Clauses
	id := len(p.Clauses) + 1
	p.Clauses[id] = cl

	// Add the constraint to the variables
	for vr := range cl {
		// New variable phase?
		_, exists := p.Variables[vr]
		if !exists {
			p.Variables[vr] = map[int]bool{}
		}

		p.Variables[vr][id] = true
	}
}

// Assign variable false if negative, true if positive, removing the variable
// from all lists and adding the assignment to the Assignment map.
// If the assignment results in a conflict the Unsatisfiable flag is set.
func (p Problem) Assign(variable int) Problem {
	np := p.Copy()

	// Save the variables assignment
	np.Assignment[abs(variable)] = variable > 0

	// Remove all instances of this phase
	// If the clause contains the current assignment, we remove the clause.
	for clause := range np.Variables[variable] {
		if len(np.Clauses[clause]) == 1 {
			// No need to update other variables, there are none.
			// So just remove the clause
			delete(np.Clauses, clause)
			continue
		}

		// We make the clause true so:
		// Delete all references to this clause...
		for vr := range np.Clauses[clause] {
			delete(np.Variables[vr], clause)
		}

		// ...and delete the clause itself
		delete(np.Clauses, clause)
	}

	// Remove all instances of the other phase
	for clause := range np.Variables[-variable] {
		// We have unsatisfiability if the opposite phase is required
		if len(np.Clauses[clause]) == 1 {
			np.Unsatisfiable = true

			// This was the only variable in the clause,
			// remove the clause and continue
			delete(np.Clauses, clause)
			continue
		}

		// There are other variables in the clause
		// only remove this phase
		delete(np.Clauses[clause], -variable)
	}

	// Remove the variable instances
	delete(np.Variables, variable)
	delete(np.Variables, -variable)

	return np
}

// Copy creates a copy of the Problem
func (p Problem) Copy() Problem {
	np := NewProblem()

	for clause, vars := range p.Clauses {
		np.Clauses[clause] = map[int]bool{}

		for vr := range vars {
			np.Clauses[clause][vr] = true
		}
	}

	for vr, clauses := range p.Variables {
		np.Variables[vr] = map[int]bool{}

		for clause := range clauses {
			np.Variables[vr][clause] = true
		}
	}

	for vr, vl := range p.Assignment {
		np.Assignment[vr] = vl
	}

	np.Unsatisfiable = p.Unsatisfiable

	return np
}
