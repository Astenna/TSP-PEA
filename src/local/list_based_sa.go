package local

import (
	"errors"
	"fmt"
	queue "github.com/jupp0r/go-priority-queue"
	"math"
	"math/rand"
	"reflect"
)

type ListBasedSimulatedAnnealing struct {
	AdjacencyMatrix [][]int
	NeighboursGenerator      NeighboursGenerator
	minimumCost		int
	bestPath		[]int
	size int
	temperatures []float64
}

func (l ListBasedSimulatedAnnealing) GetMinimumCost() int {
	return l.minimumCost
}

func (l ListBasedSimulatedAnnealing) GetBestPath() []int {
	return l.bestPath
}

func (l *ListBasedSimulatedAnnealing) Resolve(steps int) ([]int, error){
	repeatTemperature := 200

	if l.AdjacencyMatrix == nil {
		return []int{}, errors.New("Adjacency Matrix not found! Initialize struct first!")
	}


	l.size = len(l.AdjacencyMatrix[0])
	solution := l.createInitialSolution()
	currentBestSolution := solution
	currentBestCost := CalculateCost(solution, l.AdjacencyMatrix)
	temperatures := l.setInitialTemperatureList(solution)
	var floatInterface interface{}
	var currentTemperature float64
	var random float64
	var probability float64
	var temperaturesSum float64
	var acceptedWorseSolutionCount int
	var newSolution []int
	var newCost int

	for k := 0; k<steps && temperatures.Len() > 0; k++ {
		temperaturesSum = 0.0
		acceptedWorseSolutionCount = 0
		floatInterface,_ = temperatures.Pop()
		reflected := reflect.ValueOf(floatInterface)
		if reflected.IsValid() {
			currentTemperature = float64(reflected.Float())
		}

		for m := 0; m<repeatTemperature; m++ {
			index1 := rand.Intn(l.size)
			index2 := rand.Intn(l.size)
			newSolution = l.NeighboursGenerator.GetSolutionFromNeighbourhood(currentBestSolution, index1, index2)
			newCost = CalculateCost(newSolution, l.AdjacencyMatrix)

			if newCost < currentBestCost {
				currentBestCost = newCost
				currentBestSolution = newSolution
			} else {
				random = rand.Float64()
				probability = math.Exp(-float64(newCost-currentBestCost)/currentTemperature)
				if random < probability {
					acceptedWorseSolutionCount = acceptedWorseSolutionCount + 1
					temperaturesSum = temperaturesSum + (-float64(newCost-currentBestCost)/math.Log(random))
					currentBestCost = newCost
					currentBestSolution = newSolution
				}
			}
		}

		fmt.Println(currentTemperature)
		fmt.Println(acceptedWorseSolutionCount)
		if acceptedWorseSolutionCount > 0 {
			temperatures.Pop()
			newTemperature := temperaturesSum/float64(acceptedWorseSolutionCount)
			temperatures.Insert(newTemperature, -newTemperature)
		}
	}
	return currentBestSolution, nil
}

func (l ListBasedSimulatedAnnealing) setInitialTemperatureList(initialSolution []int) queue.PriorityQueue {
    var newSolution []int
	var newCost int
    var newTemperature float64
    temperaturesQueue := queue.New()
    currentBestSolution := initialSolution
    currentBestCost := CalculateCost(currentBestSolution, l.AdjacencyMatrix)
	listLength := 500
	initialProbability := 0.9

	for i:=0; i<listLength; i++ {
		index1 := rand.Intn(l.size)
		index2 := rand.Intn(l.size)
		newSolution = l.NeighboursGenerator.GetSolutionFromNeighbourhood(currentBestSolution, index1, index2)
		newCost = CalculateCost(newSolution, l.AdjacencyMatrix)

		if newCost < currentBestCost {
			currentBestCost = newCost
			currentBestSolution = newSolution
		} else {
			newTemperature = -(float64(newCost-currentBestCost))/math.Log(initialProbability)
			temperaturesQueue.Insert(newTemperature, -newTemperature)
		}
	}

	return temperaturesQueue
}

func (l ListBasedSimulatedAnnealing) createInitialSolution() []int {
	var solution []int

	for i:=0; i<l.size; i++ {
		solution = append(solution, i)
	}

	return solution
}