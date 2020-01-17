package genetic

import (
	"math/rand"
	"sort"
)

type GeneticAlgorithm struct {
	adjacencyMatrix [][]int
	size            int
	bestPath        Individual
	CrossoverProbability	float64
	MutationProbability	float64
	MaxNumberOfGenerations int
	GenerationSize int
}

func (g GeneticAlgorithm) Resolve(adjacencyMatrix [][]int) []int {
	g.adjacencyMatrix = adjacencyMatrix
	g.size = len(adjacencyMatrix[0])

	initialPopulation := g.generateInitialPopulation()
	g.LoopGenerations(initialPopulation)
	return g.bestPath.path
}

func (g GeneticAlgorithm) generateInitialPopulation() []Individual {
	initialPopulation :=  make([]Individual, g.GenerationSize)

	for i:=0; i < g.GenerationSize; i++ {
		individual := Individual{}
		initialPopulation[i] = individual.GenerateIndividual(g.size)
	}
	return initialPopulation
}

func (g *GeneticAlgorithm) LoopGenerations(initialPopulation []Individual) {
	var parentIndex int
	var children []Individual
	currentPopulation := initialPopulation

	for generationNumber := 0; generationNumber < g.MaxNumberOfGenerations; generationNumber++ {
		// CROSSOVER
		for index:=0; index < g.GenerationSize; index++ {
			shouldCrossOver := rand.Float64()

			if shouldCrossOver > g.CrossoverProbability {
				parentIndex = rand.Intn(g.GenerationSize-1)
				for index == parentIndex {
					parentIndex =rand.Intn(g.GenerationSize-1)
				}
				child1, child2 := currentPopulation[index].Crossover(currentPopulation[parentIndex])
				if len(child1.path) > 0 {
					children = append(children, child1)
					children = append(children, child2)
				}
			}
		}
		// MUTATE
		for index:=0; index < g.GenerationSize; index++ {
			shouldMutate := rand.Float64()

			if shouldMutate < g.MutationProbability {
				currentPopulation[index] = currentPopulation[index].Mutate()
			}
		}

		if generationNumber == int(float64(g.MaxNumberOfGenerations)*0.8) {
			g.MutationProbability = g.MutationProbability*2
		}
		currentPopulation = g.nextPopulation(currentPopulation, children)
		children = []Individual{}
	}
}

func (g *GeneticAlgorithm) nextPopulation(currentPopulation []Individual, children []Individual) []Individual{
	var nextPopulation []Individual

	nextPopulation = append(nextPopulation, children...)
	g.calculateWeights(currentPopulation)
	sort.Slice(currentPopulation, func(i, j int) bool {
		return currentPopulation[i].cost < currentPopulation[j].cost
	})

	individualsThatSurviveCount := g.GenerationSize - len(children)
	if len(g.bestPath.path) == 0 {
		g.bestPath= currentPopulation[0]
	} else {
		if g.bestPath.CalculateCost(g.adjacencyMatrix) > currentPopulation[0].cost {
			g.bestPath = currentPopulation[0]
		}
	}


	nextPopulation = append(nextPopulation, currentPopulation[0:individualsThatSurviveCount]...)
	return nextPopulation
}

func (g *GeneticAlgorithm) calculateWeights(population []Individual) {
	for index := 0; index < len(population); index++ {
		population[index].CalculateCost(g.adjacencyMatrix)
	}
}

