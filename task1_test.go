package haustov

import (
	"testing"
)

func TestSolutionSquareGeneratorSingle(t *testing.T) {
	tAnswer := SolutionSquareGenerator(2, 3)
	tAnswerAwait := []int{4, 9, 16}
	if len(tAnswerAwait) != len(tAnswer) {
		t.Errorf("len of SolutionSquareGenerator(2, 3) = %d; want %d", len(tAnswer), len(tAnswerAwait))
		return
	}

	for tIndex := range tAnswerAwait {
		if tAnswer[tIndex] != tAnswerAwait[tIndex] {
			t.Errorf("SolutionSquareGenerator(2, 3)[%d] = %d; want %d", tIndex, tAnswer[tIndex], tAnswerAwait[tIndex])
			return
		}
	}
}
