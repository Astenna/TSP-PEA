package main

import (
	"exact"
	"fmt"
	"genetic"
	"local"
	"log"
	"os"
	"strconv"
)

func main() {
	path := "C:\\Users\\KM\\Downloads\\PEA\\TSP\\TSP\\data120.txt"
	//path := "C:\\Users\\KM\\Downloads\\PEA\\SMALL\\data10.txt"
	adjacencyMatrix, _ := local.LoadAdjacencyMatrixFromFile(path)
	genetic := genetic.GeneticAlgorithm{}
	genetic.CrossoverProbability = 0.7
	genetic.MutationProbability = 0.15
	genetic.GenerationSize = 200
	genetic.MaxNumberOfGenerations = 2000
	result := genetic.Resolve(adjacencyMatrix)
	cost := local.CalculateCost(result, adjacencyMatrix)
	fmt.Println(cost)
}

func TestLocalSearchAlgorithms() {
	path := "C:\\Users\\KM\\Downloads\\PEA\\ATSP\\ATSP\\data"
	dataSizes := []string{"17", "34", "36", "39", "43", "45", "48", "53", "56", "65", "70", "71", "100", "171", "323", "358", "403", "443"}
	file, err := os.Create("results.csv")
	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}

	for _, size := range dataSizes {
		file.WriteString(size + "\n")
		adjacencyMatrix, _ := local.LoadAdjacencyMatrixFromFile(path + size + ".txt")

		lbAnnealing := local.ListBasedSimulatedAnnealing{AdjacencyMatrix: adjacencyMatrix}
		lbAnnealing.NeighboursGenerator = local.MultipleMove{adjacencyMatrix}

		simpleAnnealing := local.SimulatedAnnealing{AdjacencyMatrix: adjacencyMatrix}
		simpleAnnealing.NeighboursGenerator = local.MultipleMove{adjacencyMatrix}

		/* initial temperature tests */
		temperatures := []float64{100.0, 400.0, 700.0, 1000.0, 1300.0, 1600.0, 2000.0}
		for _, temperature := range temperatures {
			sum := 0
			for i := 0; i < 100; i++ {
				solution2, _ := simpleAnnealing.Resolve(10000, 0.999, temperature)
				sum += local.CalculateCost(solution2, adjacencyMatrix)
			}
			file.WriteString(size + ";0.999;" + fmt.Sprintf("%f", temperature) + ";" + strconv.Itoa(sum/100) + "\n")
		}
	}
}

func TestExactAlgorithms() {
	var instance exact.TravellingSalesmanProblem
	dataSizes := []string{"4", "10", "11", "12", "13", "14", "15", "16", "17", "18", "21"}
	algorithmsSlice := []exact.Algorithm{exact.HeldKarp{}, exact.BranchAndBound{}, exact.BruteForce{}}
	algorithmNames := []string{"HeldKarp", "BranchAndBound", "BruteForce"}
	csvExtension := ".csv"
	pathToDirectory := "C:\\Users\\KM\\Downloads\\PEA\\SMALL\\"
	baseName := "data"
	extenstion := ".txt"

	for index, algorithm := range algorithmsSlice {
		instance.Algorithm = algorithm
		file, err := os.Create(algorithmNames[index] + csvExtension)
		if err != nil {
			log.Fatalf("failed creating file: %s", err)
		}
		fmt.Println("============================= " + algorithmNames[index] + " =============================")
		for _, size := range dataSizes {
			fullPath := pathToDirectory + baseName + size + extenstion
			instance.LoadDataFromFile(fullPath)
			instance.Resolve()
			file.WriteString(size + ";" + strconv.FormatInt(instance.CalculationTime.Microseconds(), 10) + "\n")
			fmt.Println(fullPath)
			fmt.Println(instance.Solution)
			fmt.Println(instance.MinimumCost)
			fmt.Println(instance.CalculationTime)
		}

		file.Close()
	}
}
