package main

import (
	"math"
)

// STARTING_NODE defines a node to start
// the evaluation on HeldKarp algorithm
const (
	STARTING_NODE = 0
)

// PartialSolution defines value stored in partialSolutions
// map of HeldKarp algorithm
type PartialSolution struct {
	cost     int
	lastNode int
}

// Key defines identifier for PartialSolution
// in partialSolutions map of HeldKarp algorithm
type Key struct {
	lastNode     int
	visitedNodes int
}

// HeldKarp defines data structure crucial
// for resolving TSP with Held-Karp algorithm
type HeldKarp struct {
	adjacencyMatrix  [][]int
	partialSolutions map[Key]PartialSolution
	startingNode     int
}

// Resolve defines a main method that is the entry point
// for the Held Karp algorithm evaluation
func (h HeldKarp) Resolve(adjacencyMatrix [][]int) []int {
	h.adjacencyMatrix = adjacencyMatrix
	h.startingNode = STARTING_NODE
	h.partialSolutions = make(map[Key]PartialSolution)

	var nodes []int
	for i := 1; i < len(h.adjacencyMatrix[0]); i++ {
		nodes = append(nodes, i)
	}

	solution := h.calculatePaths(nodes, Key{h.startingNode, 1})
	return h.backtrackOptimalPath(solution.lastNode)
}

//
func (h *HeldKarp) calculatePaths(nodesToVisit []int, lastSolutionKey Key) PartialSolution {

	if len(nodesToVisit) <= 0 {
		var newPartialSolution PartialSolution
		newPartialSolution.cost = h.adjacencyMatrix[lastSolutionKey.lastNode][h.startingNode]
		newPartialSolution.lastNode = lastSolutionKey.lastNode

		return newPartialSolution
	}

	currentBestSolution := PartialSolution{math.MaxInt64, lastSolutionKey.lastNode}
	iterationSolution := PartialSolution{}

	for index, node := range nodesToVisit {
		keyWithNewNode := Key{node, setBit(lastSolutionKey.visitedNodes, node)}
		partialSolution, keyExists := h.partialSolutions[keyWithNewNode]

		if !keyExists {
			SwapLastAndIndex(nodesToVisit, index)
			partialSolution = h.calculatePaths(nodesToVisit[:len(nodesToVisit)-1], keyWithNewNode)
			SwapLastAndIndex(nodesToVisit, index)
			if len(nodesToVisit) > 1 {
				h.partialSolutions[keyWithNewNode] = partialSolution
			}
		}
		iterationSolution.cost = partialSolution.cost + h.adjacencyMatrix[lastSolutionKey.lastNode][node]
		iterationSolution.lastNode = node

		if iterationSolution.cost < currentBestSolution.cost {
			currentBestSolution = iterationSolution
		}
	}

	return currentBestSolution
}

func (h *HeldKarp) backtrackOptimalPath(lastNode int) []int {
	var optimalPath []int
	optimalPath = append(optimalPath, h.startingNode)
	optimalPath = append(optimalPath, lastNode)
	lastVisitedNodesKey := setBit(1, lastNode)
	lastKey := Key{lastNode, lastVisitedNodesKey}
	partialSolution := h.partialSolutions[lastKey]

	for len(optimalPath) < len(h.adjacencyMatrix[0]) {
		optimalPath = append(optimalPath, partialSolution.lastNode)
		lastVisitedNodesKey = setBit(lastVisitedNodesKey, partialSolution.lastNode)
		lastKey := Key{partialSolution.lastNode, lastVisitedNodesKey}
		partialSolution = h.partialSolutions[lastKey]
	}

	return optimalPath
}

func clearBit(n int, pos int) int {
	n &^= (1 << pos)
	return n
}

func setBit(n int, pos int) int {
	n |= (1 << pos)
	return n
}
