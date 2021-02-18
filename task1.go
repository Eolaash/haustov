package haustov

// SolutionSquareGenerator - generate int array of SQ of naturalnumbers
func SolutionSquareGenerator(start int, n int) []int {

	// logic preventers (overload preveneter? maxint*maxint; should i use math import ? how compiler system define when to use int32 or int64 :: need to read about compiler logic more - overget by instructions?)
	// current limiter for int64
	if n < 1 || start < 1 || start+n > 3037000500 {
		return nil
	}

	// blank array
	tOutArray := make([]int, n, n)

	// main work (using sum var for START+INDEX.. will it speed up MULTI?)
	for tIndex := range tOutArray {
		tOutArray[tIndex] = (start + tIndex) * (start + tIndex)
	}

	// success return
	return tOutArray
}
