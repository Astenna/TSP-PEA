package local

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

type InitialSolutionGenerator interface {
	GetInitialSolution(adjacencyMatrix [][]int) []int
}

type SimulatedAnnealing struct {
	size               int
	name               string
	adjacencyMatrix    [][]int
	solution           []int
	minimumCost        int
	calculationTime    time.Duration
	MaxCalculationTime time.Duration

	NeighboursGenerator      NeighboursGenerator
	InitialSolutionGenerator InitialSolutionGenerator
	AnnealingSchedule        AnnealingScheduler

	Steps              int
	InitialTemperature float64
	FinalTemperature   float64
}

func (s *SimulatedAnnealing) Resolve() time.Duration {
	s.size = len(s.adjacencyMatrix[0])
	s.FinalTemperature = 0
	startTime := time.Now()
	s.anneal()
	endTime := time.Now()
	s.calculationTime = endTime.Sub(startTime)
	s.minimumCost = s.GetCost(s.solution)

	return s.calculationTime
}

func (s SimulatedAnnealing) GetSolutionCost() int {
	return s.GetCost(s.solution)
}

func (s SimulatedAnnealing) GetSolution() []int {
	return s.solution
}

/*func (s SimulatedAnnealing) UseNaturalInitialSolution() {
	var notVisitedNodes []int
	b.minimumCost = math.MaxInt64
	b.adjacencyMatrix = adjacencyMatrix

	for i := 0; i < len(adjacencyMatrix); i++ {
		s.InitialSolution = append(notVisitedNodes, i)
	}
}*/

func (s *SimulatedAnnealing) anneal() {
	startTime := time.Now()
	for i := 0; i < s.size; i++ {
		s.solution = append(s.solution, i)
	}

	//s.bestPath = s.InitialSolutionGenerator.GetInitialSolution(s.adjacencyMatrix)
	currentTemperature := s.InitialTemperature
	bestPathCost := CalculateCost(s.solution, s.adjacencyMatrix)
	currentSolution := make([]int, s.size)
	var currentSolutionCost int
	var random float64
	var costChange int

	//temp
	copy(currentSolution, s.solution)
	for step := 0; s.MaxCalculationTime > time.Now().Sub(startTime) && currentTemperature > s.FinalTemperature; step++ {
		currentSolution = s.NeighboursGenerator.GetSolutionFromNeighbourhood(currentSolution)
		currentSolutionCost = CalculateCost(currentSolution, s.adjacencyMatrix)
		fmt.Println("New solution: ", currentSolution)
		fmt.Println("CurrentCost ", currentSolutionCost, " current Best cost ", bestPathCost)

		if bestPathCost > currentSolutionCost {
			bestPathCost = currentSolutionCost
			s.solution = currentSolution
		} else {
			costChange = currentSolutionCost - bestPathCost
			random = rand.Float64()
			evaluated := math.Exp(-float64(costChange) / currentTemperature)

			fmt.Println("costChange ", costChange, "random ", random, "? evaluated", evaluated, "temperature ", currentTemperature)

			if random < evaluated {
				bestPathCost = currentSolutionCost
				s.solution = currentSolution
			}
		}

		currentTemperature = s.AnnealingSchedule.NextTemperature(currentTemperature, step)
	}
}

func (s *SimulatedAnnealing) GetCost(path []int) int {
	var result int
	last := path[0]
	for _, node := range path[1:] {
		result = result + s.adjacencyMatrix[last][node]
		last = node
	}
	result = result + s.adjacencyMatrix[path[len(path)-1]][path[0]]
	return result
}

func (s *SimulatedAnnealing) LoadDataFromFile(fileName string) ([][]int, error) {

	file, err := os.Open(fileName)
	if err != nil {
		return [][]int{}, errors.New("Could not open file")
	}

	reader := bufio.NewReader(file)
	s.name, err = reader.ReadString('\n')
	size, err := reader.ReadString('\n')
	s.size, err = strconv.Atoi(strings.TrimSpace(size))
	s.adjacencyMatrix = make([][]int, s.size)

	for index := 0; index < s.size; index++ {
		line, _ := reader.ReadString('\n')
		lineOfValues := strings.Fields(line)

		row := make([]int, s.size)
		for index2, element := range lineOfValues {
			value, _ := strconv.Atoi(strings.TrimSpace(element))
			row[index2] = value
		}
		s.adjacencyMatrix[index] = row
	}

	defer file.Close()
	return s.adjacencyMatrix, nil
}
