package local

import (
	"math/rand"
	"slice"
)

type NeighboursGenerator interface {
	GetSolutionFromNeighbourhood(solution []int, index1 int, index2 int) []int
}

type Swap struct {
}

func (s Swap) GetSolutionFromNeighbourhood(solution []int, index1 int, index2 int) []int {
	slice.SwapOnIndexes(solution, index1, index2)
	return solution
}

type Reverse struct {
}

func (r Reverse) GetSolutionFromNeighbourhood(solution []int, index1 int, index2 int) []int {

	if int(index1) > int(index2) {
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

	if CalculateCost(insertSolution, m.AdjacencyMatrix) > CalculateCost(reverseSolution, m.AdjacencyMatrix) {
		bestSolution = reverseSolution
	}
	if CalculateCost(bestSolution, m.AdjacencyMatrix) > CalculateCost(swapSolution, m.AdjacencyMatrix) {
		bestSolution = swapSolution
	}

	return bestSolution
}
