package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	var tsp TravellingSalesmanProblem
	dataSizes := []string{"4", "10", "11", "12", "13", "14", "15", "16", "17", "18", "21"}
	algorithms := []TspAlgorithm{HeldKarp{}, BranchAndBound{}, BruteForce{}}
	algorithmNames := []string{"HeldKarp", "BranchAndBound", "BruteForce"}
	csvExtension := ".csv"
	pathToDirectory := "C:\\Users\\KM\\Downloads\\PEA\\SMALL\\"
	baseName := "data"
	extenstion := ".txt"

	for index, algorithm := range algorithms {
		tsp.Algorithm = algorithm
		file, err := os.Create(algorithmNames[index] + csvExtension)
		if err != nil {
			log.Fatalf("failed creating file: %s", err)
		}
		fmt.Println("============================= " + algorithmNames[index] + " =============================")
		for _, size := range dataSizes {
			fullPath := pathToDirectory + baseName + size + extenstion
			tsp.LoadDataFromFile(fullPath)
			tsp.Resolve()
			file.WriteString(size + ";" + tsp.CalculationTime.String() + "\n")
			fmt.Println(fullPath)
			fmt.Println(tsp.Solution)
			fmt.Println(tsp.MinimumCost)
			fmt.Println(tsp.CalculationTime)
		}

		file.Close()
	}
}
