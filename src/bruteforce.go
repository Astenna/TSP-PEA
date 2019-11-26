package main

import "math"

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

	b.FindAllCycles(make([]int, 0), notVisitedNodes)
	return b.bestPath
}

// FindAllCycles is a recursive function that finds all cycles in graph using search tree
func (b *BruteForce) FindAllCycles(path []int, notVisitedNodes []int) {

	if len(notVisitedNodes) > 0 {
		for index, node := range notVisitedNodes {
			SwapLastAndIndex(notVisitedNodes, index)
			b.FindAllCycles(append(path, node), notVisitedNodes[:len(notVisitedNodes)-1])
			SwapLastAndIndex(notVisitedNodes, index)
		}
	} else {
		cost := b.TargetFunction(path)
		if cost < b.minimumCost {
			b.bestPath = make([]int, len(path))
			copy(b.bestPath, path)
			b.minimumCost = cost
		}
	}
}

// TargetFunction returns total cost of given path in given adjacencyMatrix
func (b *BruteForce) TargetFunction(nodes []int) int {

	var result int
	last := nodes[0]
	for _, node := range nodes[1:] {
		result = result + b.adjacencyMatrix[last][node]
		last = node
	}

	result = result + b.adjacencyMatrix[nodes[len(nodes)-1]][nodes[0]]
	return result
}
