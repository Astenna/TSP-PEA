package main

import (
	"fmt"
	"math"
)

type Pair struct {
	cost           int
	lastSubpathKey int
}

type HeldKarp struct {
	adjacencyMatrix  [][]int
	partialSolutions map[int]Pair
	startingNode     int
}

// Resolve defines a main method of
func (h HeldKarp) Resolve(adjacencyMatrix [][]int) []int {
	h.adjacencyMatrix = adjacencyMatrix
	h.startingNode = 1
	var nodes []int
	for i := 2; i < len(h.adjacencyMatrix[0])-2; i++ {
		nodes = append(nodes, i)
	}
	solution := h.CalculatePaths(1, nodes, h.startingNode)
	fmt.Println(solution.cost)
	return h.BacktrackOptimalPath(solution)
}

func (h HeldKarp) CalculatePaths(destination int, notVisited []int, lastKey int) Pair {

	fmt.Println("Call destination: ", destination, "not visited: ", notVisited)

	if len(notVisited) <= 0 {
		var newPair Pair
		newPair.cost = h.adjacencyMatrix[h.startingNode][destination]
		newPair.lastSubpathKey = setBit(lastKey, destination)
		return newPair
	} else {
		currentBestSolution := Pair{math.MaxInt64, lastKey}

		for index, node := range notVisited {
			keyWithNewNode := setBit(lastKey, node)
			solution, keyExists := h.partialSolutions[keyWithNewNode]

			if !keyExists {
				SwapLastAndIndex(notVisited, index)
				solution = h.CalculatePaths(node, notVisited[:len(notVisited)-1], keyWithNewNode)
				SwapLastAndIndex(notVisited, index)
			}

			solution.cost = solution.cost + h.adjacencyMatrix[destination][node]
			solution.lastSubpathKey = keyWithNewNode
			fmt.Println("From notVisited ", notVisited, " chosen ", node, " cost ", solution.cost)
			if solution.cost < currentBestSolution.cost {
				currentBestSolution = solution
			}
		}
		return currentBestSolution
	}
}

func (h HeldKarp) BacktrackOptimalPath(solution Pair) []int {
	var optimumPath []int
	optimumPath = append(optimumPath, h.startingNode)
	lastKey := h.startingNode
	solutionExists := true

	for solutionExists {
		// XOR
		diff := solution.lastSubpathKey - lastKey
		node := int(math.Log2(float64(diff)))
		optimumPath = append(optimumPath, node)
		lastKey = solution.lastSubpathKey
		solution, solutionExists = h.partialSolutions[solution.lastSubpathKey]
	}

	return optimumPath
}

func clearBit(n int, pos int) int {
	n &^= (1 << pos)
	return n
}

func setBit(n int, pos int) int {
	n |= (1 << pos)
	return n
}
