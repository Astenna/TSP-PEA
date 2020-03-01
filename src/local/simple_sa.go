package local

import (
	"errors"
	"math"
	"math/rand"
	"time"
	"sliceExtensions"
)

type SimulatedAnnealing struct {
	AdjacencyMatrix [][]int
	NeighboursGenerator      NeighboursGenerator
	minimumCost		int
	bestPath		[]int
	size int
	temperatures []float64
}

func (l SimulatedAnnealing) GetMinimumCost() int {
	return l.minimumCost
}

func (l SimulatedAnnealing) GetBestPath() []int {
	return l.bestPath
}

func (l *SimulatedAnnealing) Resolve(steps int, temperatureStep float64, initialTemperature float64) ([]int, error){
	if l.AdjacencyMatrix == nil {
		return []int{}, errors.New("Adjacency Matrix not found! Initialize struct first!")
	}
	rand.Seed(time.Now().UTC().UnixNano())
	l.size = len(l.AdjacencyMatrix[0])
	solution := l.createInitialSolution()
	currentBestSolution := solution
	currentBestCost := sliceExtensions.CalculateCost(l.AdjacencyMatrix, solution)
	currentTemperature := initialTemperature

	var random float64
	var probability float64
	var newSolution []int
	var newCost int

	for k := 0; k<steps && currentTemperature > 0; k++ {

		index1 := rand.Intn(l.size)
		index2 := rand.Intn(l.size)
		newSolution = l.NeighboursGenerator.GetSolutionFromNeighbourhood(currentBestSolution, index1, index2)
		newCost = sliceExtensions.CalculateCost(l.AdjacencyMatrix, newSolution)

			if newCost < currentBestCost {
				currentBestCost = newCost
				currentBestSolution = newSolution
			} else {
				random = rand.Float64()
				probability = math.Exp(-float64(newCost-currentBestCost)/currentTemperature)
				if random < probability {
					currentBestCost = newCost
					currentBestSolution = newSolution
				}
			}

		currentTemperature = temperatureStep * currentTemperature
	}

	return currentBestSolution, nil
}

func (l SimulatedAnnealing) createInitialSolution() []int {
	var solution []int

	for i:=0; i<l.size; i++ {
		solution = append(solution, i)
	}

	return solution
}