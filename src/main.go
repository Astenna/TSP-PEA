package main

import (
	"exact"
	"fmt"
	"local"
	"log"
	"os"
	"strconv"
)

func main() {

	path := "C:\\Users\\KM\\Downloads\\PEA\\TSP\\TSP\\data17.txt"
	adjacencyMatrix, _ := local.LoadAdjacencyMatrixFromFile(path)
	lbAnnealing := local.ListBasedSimulatedAnnealing{AdjacencyMatrix: adjacencyMatrix}
	lbAnnealing.NeighboursGenerator = local.MultipleMove{adjacencyMatrix}
	solution, _ := lbAnnealing.Resolve(10000)
	fmt.Println(local.CalculateCost(solution, adjacencyMatrix))

	/*annealing := local.SimulatedAnnealing{}
	annealing.AnnealingSchedule = local.GeometricAnnealing{(0.999999)}
	annealing.InitialTemperature = 1000
	annealing.MaxCalculationTime = time.Second * 120 * 3
	annealing.NeighboursGenerator = local.Swap{}

	path := "C:\\Users\\KM\\Downloads\\PEA\\TSP\\TSP\\data17.txt"
	annealing.LoadDataFromFile(path)
	//time := annealing.Resolve()
	time := annealing.ResolveListBased(100000)
	fmt.Println(annealing.GetSolution())
	fmt.Println(annealing. GetSolutionCost())
	fmt.Println(time)*/
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
