package haustov

import "fmt"

// SolutionBinaryGap - searching longes zero-stray in binary view of N
func SolutionBinaryGap(N int) int {

	// Prepare work vars
	tMaxStray := 0
	tCurrentStray := 0
	tBinaryString := fmt.Sprintf("%b", N)

	// Main scanner
	for tIndex := range tBinaryString {
		if tBinaryString[tIndex] == 48 {
			tCurrentStray++
		} else {
			tMaxStray = tCurrentStray
			tCurrentStray = 0
		}
	}

	// Finalyzer
	if tCurrentStray > tMaxStray {
		tMaxStray = tCurrentStray
	}

	// succes return
	return tMaxStray
}
