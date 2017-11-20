/* POTW #2 - Fall 2017
Eye of the Storm
Patrick Malolepszy */

package main

import "fmt"

func stormEye(pressures []float32) int {

	n := len(pressures)

	// single measurement defaults to eye of the storm.
	if n == 1 {
		return 0
	}

	// case for first number (left is default to at least 15% above)
	if pressures[0] < pressures[1]*0.85 {
		return 0
	}

	// look at each pressure which two adjacent pressures:
	for i := 1; i < n-1; i++ {
		avg := (pressures[i-1] + pressures[i+1]) / 2
		if pressures[i] <= avg*0.85 {
			return i
		}
	}

	// case for last number (right is default to at least 15% above)
	if pressures[n-1] < pressures[n-2]*0.85 {
		return n - 1
	}

	// no eye found
	return -1
}

func main() {
	var n int
	fmt.Scan(&n)
	pressures := make([]float32, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&pressures[i])
	}
	fmt.Println(stormEye(pressures))
}
