package main

import "fmt"

func main() {
	var tsp TravellingSalesmanProblem
	tsp.LoadDataFromFile("C:\\Users\\KM\\Downloads\\PEA\\SMALL\\data10.txt")
	tsp.Algorithm = BruteForce{}
	tsp.Resolve()
	fmt.Println(tsp.Solution)
	fmt.Println(tsp.MinimumCost)
	fmt.Println(tsp.CalculationTime)
	//var fileName string
	//var testChoice string
	//var solution []int

	/*fmt.Println()
	fmt.Println("  PEA - Kinga Marek - 21.10.2019  ")
	fmt.Println("============ Stage 0. =============")
	fmt.Println()
	fmt.Println("Enter path to the file with data: ")
	fmt.Scan(&fileName)
	fmt.Println()
	_, err := g.CreateFromFile(fileName)
	for err != nil {
		fmt.Println(err)
		fmt.Println("Enter path to the file with data: ")
		fmt.Scan(&fileName)
		_, err = g.CreateFromFile(fileName)
	}
	fmt.Println("Loaded graph:")
	fmt.Println(g)
	fmt.Println()
	fmt.Println("Choose a way to test: (enter 'a' or 'b')")
	fmt.Println("a - enter nodes manually")
	fmt.Println("b - generate nodes randomly")
	fmt.Scan(&testChoice)
	for (testChoice != "a" && testChoice != "b") {
		fmt.Scan(&testChoice)
	}
	if testChoice == "a" {
		solution = g.GetCycleFromUser()
	} else {
		solution = g.CreateRandomSolution()
	}
	fmt.Println("Entered cycle:")
	fmt.Println(solution)
	fmt.Println("Cost of cycle:")
	fmt.Println(g.TargetFunction(solution))
	g.CreateFromFile("C:\\Users\\KM\\Downloads\\PEA\\SMALL\\data12.txt")
	foundCycle := g.SolveUsingBruteForce()
	fmt.Println(foundCycle)
	cost := g.TargetFunction(foundCycle)
	fmt.Println("COST: ", cost)*/
}
