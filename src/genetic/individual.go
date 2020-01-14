package genetic

import (
	"math/rand"
	"sliceExtensions"
)

type Individual struct {
	path []int
	cost int
}

func (ind Individual) GenerateIndividual(size int) Individual {

	for i:=0; i<size; i++ {
		ind.path = append(ind.path, i)
	}
	for i:=0; i<size; i++ {
		index1 := rand.Intn(size)
		index2 := rand.Intn(size)
		sliceExtensions.SwapOnIndexes(ind.path, index1, index2)
	}

	return ind
}

func (ind Individual) CalculateCost(adjacencyMatrix [][]int) {
	var result int
	last := ind.path[0]
	for _, node := range ind.path[1:] {
		result = result + adjacencyMatrix[last][node]
		last = node
	}
	result = result + adjacencyMatrix[ind.path[len(ind.path)-1]][ind.path[0]]
	ind.cost = result
}

func (ind Individual) Crossover(individual Individual) Individual {

}

func (ind Individual) Mutate() Individual {

}