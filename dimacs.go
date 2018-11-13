package main

import (
	"bufio"
	"errors"
	"os"
	"strconv"
	"strings"
)

func readDimacsCnfFile(path string) (*Problem, error) {
	file, err := os.Open(path)

	defer file.Close()

	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)

	numVars, numClauses := 0, 0
	problem := NewProblem()

	// Read header
	for scanner.Scan() {
		if scanner.Err() != nil {
			return nil, err
		}

		line := scanner.Text()
		line = strings.TrimSpace(line)

		if strings.HasPrefix(line, "p ") && strings.Fields(line)[1] == "cnf" {
			// Get the ammount of variables
			numVars, err = strconv.Atoi(strings.Fields(line)[2])

			if err != nil {
				return nil, errors.New("unknown line type '" + line + "'")
			}

			// And constraints
			numClauses, err = strconv.Atoi(strings.Fields(line)[3])

			if err != nil {
				return nil, errors.New("unknown line type '" + line + "'")
			}

			// Done reading header
			break

		} else if strings.HasPrefix(line, "c ") || line == "c" || line == "" {
			// Skip comments and empty lines

		} else {
			return nil, errors.New("unknown line type '" + line + "'")
		}
	}

	// Read clauses
	readClauses, line := 0, ""
	for scanner.Scan() {
		if scanner.Err() != nil {
			return nil, err
		}

		segment := scanner.Text()

		// Skip comments
		if strings.HasPrefix(segment, "c") {
			continue
		}

		line += " " + segment
		line = strings.TrimSpace(line)

		if !strings.HasSuffix(line, " 0") {
			continue
		}

		line = strings.TrimSpace(strings.TrimSuffix(line, " 0"))

		if line == "%" {
			// End of file marker, seen in some files
			break
		}

		readClauses++

		if readClauses > numClauses {
			return nil, errors.New("too many clauses in file")
		}

		clause := map[int]bool{}

		for _, v := range strings.Fields(line) {
			vr, err := strconv.Atoi(v)

			if err != nil || vr > numVars || vr == 0 {
				return nil, errors.New("invalid constraint " + line)
			}

			clause[vr] = true
		}

		// Add clause and reset line
		problem.AddClause(clause)
		line = ""
	}

	// Check if the ammount of clauses is correct
	if readClauses != numClauses {
		return nil, errors.New("too few clauses in file")
	}

	return problem, nil
}
