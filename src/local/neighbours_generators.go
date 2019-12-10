package local

import (
	"math/rand"
	"slice"
)

type NeighboursGenerator interface {
	GetSolutionFromNeighbourhood(solution []int) []int
}

type Swap struct {
}

func (s Swap) GetSolutionFromNeighbourhood(solution []int) []int {
	index1 := rand.Intn(len(solution))
	index2 := rand.Intn(len(solution))
	slice.SwapOnIndexes(solution, index1, index2)
	return solution
}

type Reverse struct {
}

func (r Reverse) GetSolutionFromNeighbourhood(solution []int) []int {
	index1 := rand.Intn(len(solution) - 1)
	index2 := rand.Intn(len(solution) - 1)

	comparison := int(index1) > int(index2)

	if comparison {
		replaced := index1
		index1 = index2
		index2 = replaced
	}

	for ; index1 <= index2; index1, index2 = index1+1, index2-1 {
		slice.SwapOnIndexes(solution, index1, index2)
	}

	return solution
}

type Insert struct {
}

func (i Insert) GetSolutionFromNeighbourhood(solution []int) []int {
	var newSolution []int
	index1 := rand.Intn(len(solution) - 1)
	index2 := rand.Intn(len(solution) - 1)

	if index1 > index2 {
		replaced := index1
		index1 = index2
		index2 = replaced
	}

	copy(newSolution, solution[0:index1])
	newSolution[index1+1] = solution[index2]
	copy(newSolution[index1+2:index2], solution[index1+1:index2-1])
	copy(newSolution[index2+1:], solution[index2+1:])

	return newSolution
}
