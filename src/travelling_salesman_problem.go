package main

import (
	"bufio"
	"errors"
	"os"
	"strconv"
	"strings"
	"time"
)

// TspAlgorithm defines a way to solve travelling salesman problem
type TspAlgorithm interface {
	Resolve(adjacencyMatrix [][]int) []int
}

// TravellingSalesmanProblem defines a complete data
// and an algorithm used to solve the problem.
type TravellingSalesmanProblem struct {
	Name            string
	Size            int
	AdjacencyMatrix [][]int
	Solution        []int
	MinimumCost     int
	CalculationTime time.Duration
	Algorithm       TspAlgorithm
}

// LoadDataFromFile reads data from given file and saves it to an AdjacencyMatrix
// The file should have a structure:
// name and size - each in new line
// and an adjacency matrix
func (t *TravellingSalesmanProblem) LoadDataFromFile(fileName string) ([][]int, error) {

	file, err := os.Open(fileName)
	if err != nil {
		return [][]int{}, errors.New("Could not open file")
	}

	reader := bufio.NewReader(file)
	t.Name, err = reader.ReadString('\n')
	size, err := reader.ReadString('\n')
	t.Size, err = strconv.Atoi(strings.TrimSpace(size))
	t.AdjacencyMatrix = make([][]int, t.Size)

	for index := 0; index < t.Size; index++ {
		line, _ := reader.ReadString('\n')
		lineOfValues := strings.Fields(line)

		row := make([]int, t.Size)
		for index2, element := range lineOfValues {
			value, _ := strconv.Atoi(strings.TrimSpace(element))
			row[index2] = value
		}
		t.AdjacencyMatrix[index] = row
	}

	defer file.Close()
	return t.AdjacencyMatrix, nil
}

// Resolve is defined to solve the TSP using the given algorithm
func (t *TravellingSalesmanProblem) Resolve() {
	startTime := time.Now()
	t.Solution = t.Algorithm.Resolve(t.AdjacencyMatrix)
	endTime := time.Now()
	t.CalculationTime = endTime.Sub(startTime)
	t.MinimumCost = t.CalculateCost(t.Solution)
}

// CalculateCost returns total cost of given path
func (t *TravellingSalesmanProblem) CalculateCost(nodes []int) int {
	var result int
	last := nodes[0]
	for _, node := range nodes[1:] {
		result = result + t.AdjacencyMatrix[last][node]
		last = node
	}
	result = result + t.AdjacencyMatrix[nodes[len(nodes)-1]][nodes[0]]
	return result
}
