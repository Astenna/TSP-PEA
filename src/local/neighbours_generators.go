package local

import (
	"math/rand"
	"sliceExtensions"
)

type NeighboursGenerator interface {
	GetSolutionFromNeighbourhood(solution []int, index1 int, index2 int) []int
}

type Swap struct {
}

func (s Swap) GetSolutionFromNeighbourhood(solution []int, index1 int, index2 int) []int {
	sliceExtensions.SwapOnIndexes(solution, index1, index2)
	return solution
}

type Reverse struct {
}

func (r Reverse) GetSolutionFromNeighbourhood(solution []int, index1 int, index2 int) []int {

	if index1 > index2 {
		replaced := index1
		index1 = index2
		index2 = replaced
	}

	for ; index1 <= index2; index1, index2 = index1+1, index2-1 {
		sliceExtensions.SwapOnIndexes(solution, index1, index2)
	}

	return solution
}

type Insert struct {
}

func (i Insert) GetSolutionFromNeighbourhood(solution []int, index1 int, index2 int) []int {
	newSolution := make([]int, len(solution))

	if index1 == index2 {
		index2 = rand.Intn(len(solution) - 1)
	}

	if index1 > index2 {
		replaced := index1
		index1 = index2
		index2 = replaced
	}

	if index1 != 0 {
		copy(newSolution, solution[0:index1])
	}
	newSolution[index1] = solution[index2]
	copy(newSolution[index1+1:index2+1], solution[index1:index2])
	copy(newSolution[index2+1:len(solution)], solution[index2+1:])
	copy(solution, newSolution)
	return newSolution
}

type MultipleMove struct{
	AdjacencyMatrix [][]int
}

func (m MultipleMove) GetSolutionFromNeighbourhood(solution []int, index1 int, index2 int) []int{
	insertSolution := Insert{}.GetSolutionFromNeighbourhood(append([]int(nil),solution...), index1, index2)
	reverseSolution := Reverse{}.GetSolutionFromNeighbourhood(append([]int(nil),solution...), index1, index2)
	swapSolution := Swap{}.GetSolutionFromNeighbourhood(append([]int(nil),solution...), index1, index2)

	bestSolution := insertSolution

	if sliceExtensions.CalculateCost(m.AdjacencyMatrix, insertSolution) > sliceExtensions.CalculateCost(m.AdjacencyMatrix, reverseSolution) {
		bestSolution = reverseSolution
	}
	if sliceExtensions.CalculateCost(m.AdjacencyMatrix, bestSolution) > sliceExtensions.CalculateCost(m.AdjacencyMatrix, swapSolution) {
		bestSolution = swapSolution
	}

	return bestSolution
}
