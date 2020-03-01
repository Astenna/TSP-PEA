package main

import (
	"exact"
	"fmt"
	"genetic"
	"local"
	"log"
	"math/rand"
	"os"
	"sliceExtensions"
	"strconv"
	"time"
)

func main() {
	testGenetic()
}

func testGenetic() {
	rand.Seed(time.Now().UnixNano())
	path := "C:\\Users\\KM\\Downloads\\PEA\\ATSP\\ATSP\\data"
	dataSizes := []string{"48", "53", "70", "100", "323", "403", "443"}
	opt := []int{14422, 6905, 38673, 36230, 1326, 2465, 2720 }
	//maxGenerations := []int{1000, 1500, 2000, 2500, 3000}
	//crossProbability := []float64{0.6, 0.7, 0.8, 0.85, 0.9, 0.95}
	mutationProbability := []float64{0.009, 0.03, 0.1, 0.15, 0.2, 0.3}
	//populationSize := []int{20, 50, 100, 150, 200, 300}
	extension := ".txt"

	file, err := os.Create("results.csv")
	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}

	for index, size := range dataSizes {
		adjacencyMatrix, _ := local.LoadAdjacencyMatrixFromFile(path + size + extension)
		file.WriteString("\n")
		for _, max := range mutationProbability {
			fullCost := 0
			for i := 0; i < 100; i++ {
				genetic := genetic.GeneticAlgorithm{}
				genetic.CrossoverProbability = 0.8
				genetic.MutationProbability = max
				genetic.GenerationSize = 100
				genetic.MaxNumberOfGenerations = 2500
				result := genetic.Resolve(adjacencyMatrix)
				cost := sliceExtensions.CalculateCost(adjacencyMatrix, result)
				fullCost += cost
			}
			//fmt.Println("Rozmiar ", size, "	", cost)
			difference := (float64(fullCost)/float64(100)) - float64(opt[index])
			file.WriteString(fmt.Sprint(float64(fullCost)/float64(100)) + ";" + fmt.Sprint(difference/float64(opt[index])) + ";")
		}
	}
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
				sum += sliceExtensions.CalculateCost(adjacencyMatrix, solution2)
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
	extension := ".txt"

	for index, algorithm := range algorithmsSlice {
		instance.Algorithm = algorithm
		file, err := os.Create(algorithmNames[index] + csvExtension)
		if err != nil {
			log.Fatalf("failed creating file: %s", err)
		}
		fmt.Println("============================= " + algorithmNames[index] + " =============================")
		for _, size := range dataSizes {
			fullPath := pathToDirectory + baseName + size + extension
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
