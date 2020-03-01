package exact

import (
	"bufio"
	"errors"
	"os"
	"sliceExtensions"
	"strconv"
	"strings"
	"time"
)

// Algorithm defines a way to solve travelling salesman problem
type Algorithm interface {
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
	Algorithm       Algorithm
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
	t.MinimumCost = sliceExtensions.CalculateCost(t.AdjacencyMatrix, t.Solution)
}
