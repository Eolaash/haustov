package haustov

// SolutionRotate - Array shifter
func SolutionRotate(A []int, k int) []int {

	// Preventer
	if k <= 0 || len(A) == 0 {
		return A
	}

	// Autocorrection to short fullround shiftings
	tShift := k % len(A)

	// Preventer
	if tShift == 0 {
		return A
	}

	// Shift direction correction (LEN-SHIFT = RIGHT SHIFT)
	tShift = len(A) - tShift

	//2 methods - slice and array rebuilding in 1 cycle
	//using slices (MAKE to save the CAP\LEN of original array; coz append can extend and overCAP original)
	tSliceOut := make([]int, len(A), cap(A))
	tSliceOut = append(A[tShift:], A[:tShift]...)

	// success return
	return tSliceOut
}
