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

func (ind *Individual) CalculateCost(adjacencyMatrix [][]int) int {
	ind.cost = sliceExtensions.CalculateCost(adjacencyMatrix, ind.path)
	return ind.cost
}

/// OX - Order Crossover Operator
func (ind Individual) Crossover(individual Individual) (child1, child2 Individual) {

	p1 := make([]int, len(ind.path))
	copy(p1, ind.path)
	p2 := make([]int, len(individual.path))
	copy(p2, individual.path)
	o1 := make([]int, len(p1))
	o2 := make([]int, len(p1))

	index1, index2 := ind.getTwoRandomIndexes(len(p1)-1)
	if index1 > index2 {
		replaced := index1
		index1 = index2
		index2 = replaced
	}

	map1 := make(map[int]bool, len(p1))
	for i:=0; i<len(p1); i++ {
		map1[i] = false
	}
	for i:=index1; i<index2; i++ {
		map1[p1[i]]=true
	}

	map2 := make(map[int]bool, len(p1))
	for i:=0; i<len(p1); i++ {
		map2[i] = false
	}
	for i:=index1; i<index2; i++ {
		map2[p2[i]]=true
	}

	copy(o1[index1:index2], p1[index1:index2])
	copy(o2[index1:index2], p2[index1:index2])

	// from index2 to the end
	otherParent := index2
	for i:= index2; i<len(p1); i++ {
		indexAssigned := false
		for !indexAssigned {
			if !map1[p2[otherParent]] {
				map1[p2[otherParent]] = true
				o1[i] = p2[otherParent]
				indexAssigned = true
			}
			otherParent++
			if otherParent == len(p1) {
				otherParent = 0
			}
		}
	}

	// from start to index1
	for i:= 0; index1>i; i++ {
		indexAssigned := false
		for !indexAssigned {
			if !map1[p2[otherParent]] {
				map1[p2[otherParent]] = true
				o1[i] = p2[otherParent]
				indexAssigned = true
			}
			otherParent++
			if otherParent == len(p1) {
				otherParent = 0
			}
		}
	}

	// from index2 to the end
	otherParent = index2
	for i:= index2; i<len(p1); i++ {
		indexAssigned := false
		for !indexAssigned {
			if !map2[p1[otherParent]] {
				map2[p1[otherParent]] = true
				o2[i] = p1[otherParent]
				indexAssigned = true
			}
			otherParent++
			if otherParent == len(p1) {
				otherParent = 0
			}
		}
	}

	// from start to index1
	for i:= 0; index1>i; i++ {
		indexAssigned := false
		for !indexAssigned {
			if !map2[p1[otherParent]] {
				map2[p1[otherParent]] = true
				o2[i] = p1[otherParent]
				indexAssigned = true
			}
			otherParent++
			if otherParent == len(p1) {
				otherParent = 0
			}
		}
	}

	child1 = Individual{
		path: o1,
		cost: 0,
	}
	child2 = Individual{
		path: o2,
		cost: 0,
	}
	return child1, child2
}

func (ind Individual) Mutate() Individual {

	index1, index2 := ind.getTwoRandomIndexes(len(ind.path)-1)
	sliceExtensions.SwapOnIndexes(ind.path, index1, index2)

	return ind
}

func (ind Individual) getTwoRandomIndexes(maxIndex int) (index1, index2 int) {
	index1 = rand.Intn(maxIndex)
	index2 = rand.Intn(maxIndex)

	for index1 == index2 {
		index2 = rand.Intn(maxIndex)
	}

	return index1, index2
}