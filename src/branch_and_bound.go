package main

import (
	"fmt"
	"math"
)

// BranchAndBound defines data structure crucial
// for resolving TSP with Branch and Bound algorithm
type BranchAndBound struct {
	adjacencyMatrix [][]int
	startingNode    int
	bestPath        []int
	minimumCost     int
}

// Resolve defines a main method that is the entry point
// for the Branch and Bound algorithm evaluation
func (b BranchAndBound) Resolve(adjacencyMatrix [][]int) []int {
	b.adjacencyMatrix = adjacencyMatrix
	b.startingNode = STARTING_NODE
	b.minimumCost = math.MaxInt64
	var isVisited map[int]bool
	isVisited = make(map[int]bool)

	for i := 0; i < len(b.adjacencyMatrix[0]); i++ {
		isVisited[i] = false
	}

	isVisited[0] = true
	firstLower := b.calculateFirstLowerBound()
	fmt.Println("FIRST LOWER BOUND: ", firstLower)
	b.evaluateRecursively(firstLower, []int{STARTING_NODE}, isVisited, 0)
	fmt.Println("BB:", b.bestPath)
	fmt.Println("BB:", b.minimumCost)
	return b.bestPath
}

func (b *BranchAndBound) evaluateRecursively(lowerBound int, path []int, isVisited map[int]bool, cost int) {

	if len(path) == len(b.adjacencyMatrix[0]) {
		fmt.Println("Add to ", cost, " + ", b.adjacencyMatrix[path[len(path)-1]][0])
		cost += b.adjacencyMatrix[path[len(path)-1]][0]
		if cost < b.minimumCost {
			b.minimumCost = cost
			b.bestPath = path
			fmt.Println("NEW BEST path", b.minimumCost, ":", b.bestPath)
		}

		return
	}

	for i := 0; i < len(b.adjacencyMatrix[0]); i++ {

		if !isVisited[i] {

			newBound := lowerBound
			newCost := cost
			if len(path) == 1 {
				newBound -= (b.findFirstMin(0) + b.findFirstMin(i)) / 2
			} else {
				//fmt.Println(path)
				//fmt.Println("FOR ", path[len(path)-1], "MINUS ", b.findSecondMin(path[len(path)-1]))
				//fmt.Println("FOR ", i, "MINUS ", b.findFirstMin(i))
				newBound -= (b.findSecondMin(path[len(path)-1]) + b.findFirstMin(i)) / 2
			}

			isVisited[i] = true
			newBound += b.adjacencyMatrix[path[len(path)-1]][i]
			newCost += b.adjacencyMatrix[path[len(path)-1]][i]

			//fmt.Println("NEW LOWER BOUND: ", newBound, " PATH ", path, " newNode ", i)
			if newBound < b.minimumCost {
				//fmt.Println("will call recursively, replaced: ", b.minimumCost, "with cost ", cost)
				newPath := append(path, i)
				b.evaluateRecursively(newBound, newPath, isVisited, newCost)
			}

			isVisited[i] = false
		}
	}
}

func (b *BranchAndBound) calculateFirstLowerBound() int {
	lowerBound := 0
	for row := 0; row < len(b.adjacencyMatrix[0]); row++ {
		lowerBound += b.findFirstMin(row)
		lowerBound += b.findSecondMin(row)
	}

	return lowerBound / 2
}

func (b BranchAndBound) findFirstMin(row int) int {
	min := math.MaxInt64

	for column := 0; column < len(b.adjacencyMatrix[0]); column++ {
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

	for column := 0; column < len(b.adjacencyMatrix[0]); column++ {
		currentValue := b.adjacencyMatrix[row][column]

		if currentValue <= min2 && currentValue != -1 {
			if currentValue == min1 {
				if !isMin1AlreadyFound {
					isMin1AlreadyFound = true
				} else {
					min2 = min1
				}
			} else {
				min2 = currentValue
			}
		}
	}

	return min2
}
