package main

// Problem represents a problem
// Note that we store the two phases of a variable separately.
type Problem struct {
	// Mapping from a clause to a list of it's variables,
	// if the variable is positive/negative it expects true/false
	Clauses map[uint]map[int]bool

	// Mapping from a variable phase to the clauses containing it
	Variables map[int]map[uint]bool

	// Set of assignments
	Assigned map[int]bool

	// A flag showing wheither the current assignment is UNSAT
	Unsatisfiable bool
}

// NewProblem creates a new problem
func NewProblem() *Problem {
	return &Problem{
		Clauses:       map[uint]map[int]bool{},
		Variables:     map[int]map[uint]bool{},
		Assigned:      map[int]bool{},
		Unsatisfiable: false,
	}
}

// AddClause adds a clause to the problem, updating the Variables and Clauses maps
// this does not update Unsatisfiable, would be fun if it would be that easy though...
func (p *Problem) AddClause(clause map[int]bool) {
	// Check for tautologies.
	for vr := range clause {
		if clause[-vr] {
			// Other phase of variable already in clause
			// we have a tautology
			return
		}
	}

	// No tautologies found
	// Add the contraint to the Clauses
	id := uint(len(p.Clauses) + 1)
	p.Clauses[id] = clause

	// Add the constraint to the variables
	for vr := range clause {
		// New variable phase?
		_, exists := p.Variables[vr]
		if !exists {
			p.Variables[vr] = map[uint]bool{}
		}

		p.Variables[vr][id] = true
	}
}

// Assign variable false if negative, true if positive, removing the variable
// from all lists and adding the assignment to the Assignment map.
// If the assignment results in a conflict the Unsatisfiable flag is set.
func (p *Problem) Assign(variable int) {
	// Save the variables assignment
	p.Assigned[variable] = true

	// Remove all instances of this phase
	// If the clause contains the current assignment, we remove the clause.
	for clause := range p.Variables[variable] {
		if len(p.Clauses[clause]) == 1 {
			// No need to update other variables, there are none.
			// So just remove the clause
			delete(p.Clauses, clause)
			continue
		}

		// We make the clause true so:
		// Delete all references to this clause...
		for vr := range p.Clauses[clause] {
			// If the variable phase only appeared here,
			// remove the variable phase
			if len(p.Variables[vr]) == 1 {
				delete(p.Variables, vr)

			} else {
				delete(p.Variables[vr], clause)
			}
		}

		// ...and delete the clause itself
		delete(p.Clauses, clause)
	}

	// Remove all instances of the other phase
	for clause := range p.Variables[-variable] {
		// We have unsatisfiability if the opposite phase is required
		if len(p.Clauses[clause]) == 1 {
			p.Unsatisfiable = true

			// This was the only variable in the clause,
			// remove the clause and continue
			delete(p.Clauses, clause)
			continue
		}

		// There are other variables in the clause
		// only remove this phase
		delete(p.Clauses[clause], -variable)
	}

	// Remove the variable instances
	delete(p.Variables, variable)
	delete(p.Variables, -variable)
}

// Copy creates a copy of the Problem
func (p *Problem) Copy() *Problem {
	np := NewProblem()

	for clause, vars := range p.Clauses {
		np.Clauses[clause] = map[int]bool{}

		for vr := range vars {
			np.Clauses[clause][vr] = true
		}
	}

	for vr, clauses := range p.Variables {
		np.Variables[vr] = map[uint]bool{}

		for clause := range clauses {
			np.Variables[vr][clause] = true
		}
	}

	for vr := range p.Assigned {
		np.Assigned[vr] = true
	}

	np.Unsatisfiable = p.Unsatisfiable

	return np
}
