package exact

import (
	"math"
	"slice"
)

// BruteForce defines solution for TSP
type BruteForce struct {
	minimumCost     int
	bestPath        []int
	adjacencyMatrix [][]int
}

// Resolve defines a main method of
func (b BruteForce) Resolve(adjacencyMatrix [][]int) []int {
	var notVisitedNodes []int
	b.minimumCost = math.MaxInt64
	b.adjacencyMatrix = adjacencyMatrix

	for i := 0; i < len(adjacencyMatrix); i++ {
		notVisitedNodes = append(notVisitedNodes, i)
	}

	b.FindBestPathRecursively(make([]int, 0), notVisitedNodes)
	return b.bestPath
}

// FindBestPathRecursively is a recursive function that finds all cycles in graph using search tree
func (b *BruteForce) FindBestPathRecursively(path []int, notVisitedNodes []int) {

	if len(notVisitedNodes) > 0 {
		for index, node := range notVisitedNodes {
			slice.SwapLastAndIndex(notVisitedNodes, index)
			b.FindBestPathRecursively(append(path, node), notVisitedNodes[:len(notVisitedNodes)-1])
			slice.SwapLastAndIndex(notVisitedNodes, index)
		}
	} else {
		cost := CalculateCost(path, b.adjacencyMatrix)
		if cost < b.minimumCost {
			b.bestPath = make([]int, len(path))
			copy(b.bestPath, path)
			b.minimumCost = cost
		}
	}
}
