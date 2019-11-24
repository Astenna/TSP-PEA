package main

import (
	"math"
)

// BranchAndBound defines data structure crucial
// for resolving TSP with Branch and Bound algorithm
type BranchAndBound struct {
	adjacencyMatrix [][]int
	size            int
	startingNode    int
	bestPath        []int
	upperBound      int
}

// Resolve defines a main method that is the entry point
// for the Branch and Bound algorithm evaluation
func (b BranchAndBound) Resolve(adjacencyMatrix [][]int) []int {
	b.adjacencyMatrix = adjacencyMatrix
	b.size = len(adjacencyMatrix[0])
	b.startingNode = STARTING_NODE
	b.upperBound = math.MaxInt64
	isVisited := make(map[int]bool)

	for i := 0; i < b.size; i++ {
		isVisited[i] = false
	}

	isVisited[STARTING_NODE] = true
	firstLower := b.calculateFirstLowerBound()
	b.evaluateRecursively(firstLower, []int{STARTING_NODE}, isVisited, 0)

	return b.bestPath
}

func (b *BranchAndBound) evaluateRecursively(lowerBound int, path []int, isVisited map[int]bool, cost int) {

	if len(path) == b.size {
		cost += b.adjacencyMatrix[path[len(path)-1]][0]
		if cost < b.upperBound {
			b.bestPath = make([]int, len(path))
			copy(b.bestPath, path)
			b.upperBound = cost
		}

		return
	}

	for i := 0; i < b.size; i++ {

		if !isVisited[i] {

			newBound := lowerBound
			newCost := cost

			if len(path) == 1 {
				newBound -= (b.findFirstMin(0) + b.findFirstMin(i)) / 2
			} else {
				newBound -= (b.findSecondMin(path[len(path)-1]) + b.findFirstMin(i)) / 2
			}

			isVisited[i] = true
			newCost += b.adjacencyMatrix[path[len(path)-1]][i]

			if newBound+newCost <= b.upperBound {
				newPath := append(path, i)
				b.evaluateRecursively(newBound, newPath, isVisited, newCost)
			}

			isVisited[i] = false
		}
	}
}

func (b *BranchAndBound) calculateFirstLowerBound() int {
	lowerBound := 0
	for row := 0; row < b.size; row++ {
		lowerBound += b.findFirstMin(row)
		lowerBound += b.findSecondMin(row)
	}

	return lowerBound / 3
}

func (b BranchAndBound) findFirstMin(row int) int {
	min := math.MaxInt64

	for column := 0; column < b.size; column++ {
		if b.adjacencyMatrix[row][column] < min && b.adjacencyMatrix[row][column] != -1 {
			min = b.adjacencyMatrix[row][column]
		}
	}

	return min
}

func (b BranchAndBound) findSecondMin(row int) int {
	min1 := b.findFirstMin(row)
	isMin1AlreadyFound := false
	min2 := math.MaxInt64

	for column := 0; column < b.size; column++ {
		currentValue := b.adjacencyMatrix[row][column]

		if currentValue <= min2 && currentValue != -1 {
			if currentValue == min1 {
				if !isMin1AlreadyFound {
					isMin1AlreadyFound = true
				} else {
					return min1
				}
			} else {
				min2 = currentValue
			}
		}
	}

	return min2
}
