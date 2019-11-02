package main

import (
	"math"
)

// BruteForce defines solution for TSP
type BruteForce struct {
}

// Resolve defines a main method of
func (b BruteForce) Resolve(adjacencyMatrix [][]int) []int {
	var notVisitedNodes []int
	var solution []int

	for i := 0; i < len(adjacencyMatrix); i++ {
		notVisitedNodes = append(notVisitedNodes, i)
	}

	allCycles := b.FindAllCycles(make([]int, 0), notVisitedNodes)
	min := math.MaxInt64

	for _, cycle := range allCycles {
		currentCost := b.TargetFunction(cycle, adjacencyMatrix)
		if currentCost < min {
			min = currentCost
			solution = cycle
		}
	}

	return solution
}

// FindAllCycles is a recursive function that finds all cycles in graph using search tree
func (b *BruteForce) FindAllCycles(path []int, notVisitedNodes []int) [][]int {

	var cycles [][]int

	if len(notVisitedNodes) > 0 {
		for index, node := range notVisitedNodes {

			/*notVisitedNodesModified := make([]int, len(notVisitedNodes))
			copy(notVisitedNodesModified, notVisitedNodes)

			notVisitedNodesModified[index] = notVisitedNodesModified[len(notVisitedNodes)-1]
			notVisitedNodesModified = notVisitedNodesModified[:len(notVisitedNodes)-1]*/
			b.Swap(notVisitedNodes, index)
			cycles = append(cycles, b.FindAllCycles(append(path, node), notVisitedNodes[:len(notVisitedNodes)-1])...)
			b.Swap(notVisitedNodes, index)
		}
	} else {
		return append(cycles, path)
	}
	return cycles
}

func (b *BruteForce) Swap(path []int, index int) {
	if len(path) > 0 {
		replaced := path[len(path)-1]
		path[len(path)-1] = path[index]
		path[index] = replaced
	}
}

// TargetFunction returns total cost of given path in given adjacencyMatrix
func (b *BruteForce) TargetFunction(nodes []int, adjacencyMatrix [][]int) int {
	var result int
	last := nodes[0]
	for _, node := range nodes[1:] {
		result = result + adjacencyMatrix[last][node]
		last = node
	}
	result = result + adjacencyMatrix[nodes[len(nodes)-1]][nodes[0]]
	return result
}
