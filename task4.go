package haustov

// SolutionUniq - Count UNIQ items in []int array
// IN:
// A []int
// OUT:
// int - Count of UNIQ intems
func SolutionUniq(A []int) int {

	// Preventer
	if len(A) < 2 {
		return len(A)
	}

	// HashTable for uniq counting
	tUniqHashTable := make(map[int]int)

	// Count items (++ is overtask; zeroing better?)
	for tIndex := range A {
		tUniqHashTable[A[tIndex]]++
	}

	// success return (len of HASH = uniq items)
	return len(tUniqHashTable)
}
