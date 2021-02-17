package haustov

import "fmt"

// TestTask for export test
func TestTask() {
	fmt.Println("HELLO!")
}

// SolutionSquareGenerator - generate int array of SQ of naturalnumbers
func SolutionSquareGenerator(start int, n int) []int {

	// logic preventers (overload preveneter? maxint*maxint)
	if n < 1 || start < 1 {
		tOutArray := make([]int, 0, 0)

		// fail return
		return tOutArray
	}

	// blank array
	tOutArray := make([]int, n, n)
	for tIndex := range tOutArray {
		tOutArray[tIndex] = (start + tIndex) * (start + tIndex)
	}

	// success return
	return tOutArray
}
