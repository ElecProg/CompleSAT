package main

// Abs returns the absolute value of an integer
func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}
