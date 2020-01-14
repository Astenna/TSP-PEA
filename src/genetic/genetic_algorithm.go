package genetic

import (
	"math/rand"
	"sort"
)

type GeneticAlgorithm struct {
	adjacencyMatrix [][]int
	size            int
	bestPath        []int
	CrossoverProbability	float64
	MutationProbability	float64
	MaxNumberOfGenerations int
	GenerationSize int
}

func (g GeneticAlgorithm) Resolve(adjacencyMatrix [][]int) []int {
	g.adjacencyMatrix = adjacencyMatrix
	g.size = len(adjacencyMatrix[0])

	initialPopulation := g.generateInitialPopulation()

}

func (g GeneticAlgorithm) generateInitialPopulation() []Individual {
	initialPopulation :=  make([]Individual, g.GenerationSize)

	for i:=0; i < g.GenerationSize; i++ {
		individual := Individual{}
		initialPopulation = append(initialPopulation, individual.GenerateIndividual(g.size))
	}
	return initialPopulation
}

func (g GeneticAlgorithm) LoopGenerations(initialPopulation []Individual) {
	var parentIndex int
	var children []Individual
	currentPopulation := initialPopulation

	for generationNumber := 0; generationNumber < g.MaxNumberOfGenerations; generationNumber++ {
		// CROSSOVER
		for index:=0; index < g.GenerationSize; index++ {
			shouldCrossOver := rand.Float64()

			if shouldCrossOver > g.CrossoverProbability {
				parentIndex = rand.Intn(g.GenerationSize-1)
				for index != parentIndex {
					parentIndex =rand.Intn(g.GenerationSize-1)
				}
				newChild := currentPopulation[index].Crossover(currentPopulation[parentIndex])
				children = append(children, newChild)
			}
		}
		// MUTATE
		for index:=0; index < g.GenerationSize; index++ {
			shouldMutate := rand.Float64()

			if shouldMutate > g.MutationProbability {
				currentPopulation[index] = currentPopulation[index].Mutate()
			}
		}

		currentPopulation = g.nextPopulation(currentPopulation, children)
		children = []Individual{}
	}
}

func (g GeneticAlgorithm) nextPopulation(currentPopulation []Individual, children []Individual) []Individual{
	var nextPopulation []Individual

	nextPopulation = append(nextPopulation, children...)
	g.calculateWeights(currentPopulation)
	sort.Slice(currentPopulation, func(i, j int) bool {
		return currentPopulation[i].cost < currentPopulation[j].cost
	})

	individualsThatSurviveCount := g.GenerationSize - len(children)
	nextPopulation = append(nextPopulation, currentPopulation[0:individualsThatSurviveCount]...)
	return nextPopulation
}

func (g GeneticAlgorithm) calculateWeights(population []Individual) {
	for index := 0; index < len(population); index++ {
		population[index].CalculateCost(g.adjacencyMatrix)
	}
}

