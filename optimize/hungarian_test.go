package optimize

import (
	"testing"
)

func TestHungarian(t *testing.T) {
	weightMatrix := [][]int{
		{4, 10, 10, 10, 2, 9, 3},
		{6, 8, 5, 12, 9, 7, 2},
		{11, 9, 6, 7, 9, 5, 15},
		{3, 9, 6, 7, 5, 6, 3},
		{2, 6, 5, 3, 2, 4, 2},
		{10, 8, 11, 4, 11, 2, 11},
		{3, 4, 5, 4, 3, 6, 8},
	}
	matching := Hungarian(weightMatrix, true)

	nrow := len(matching)
	ncol := len(matching[0])
	score := 0

	for row := 0; row < nrow; row++ {
		for col := 0; col < ncol; col++ {
			if matching[row][col] {
				score += weightMatrix[row][col]
			}
		}
	}

	if score != 65 {
		t.Errorf("Score %d is not 65", score)
	}
}
