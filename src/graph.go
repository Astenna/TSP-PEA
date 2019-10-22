package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
	"math/rand"
	"errors"
)

type Graph struct {
	Name string
	Size int
	Values [][]int
}

func (g Graph) String() string {
	return fmt.Sprintf("Name: %vSize: %v\nValues: %v", g.Name, g.Size, g.Values)
}

func (g *Graph) CreateFromFile(fileName string) (*Graph, error) {
	
	file, err := os.Open(fileName)
	if err!= nil {
		return nil,  errors.New("Could not open file!")
	}

	reader := bufio.NewReader(file)
	g.Name, err =  reader.ReadString('\n')
	size, err := reader.ReadString('\n')
	g.Size, err = strconv.Atoi(strings.TrimSpace(size))
	g.Values = make([][]int, g.Size)

	for index:=0; index<g.Size; index++ {
		line, _ := reader.ReadString('\n')
		lineOfValues := strings.Fields(line);
		
		row  := make([]int, g.Size)
		for index2, element := range lineOfValues {
			value, _ := strconv.Atoi(strings.TrimSpace(element))
			row[index2] = value
		}
		g.Values[index] = row
    }

	defer file.Close()
	return g, nil
}

func (g *Graph) CreateRandomSolution() []int {
	var solution []int
	var notVisitedNodes []int

	for i:=1; i<g.Size; i++ {
		notVisitedNodes = append(notVisitedNodes, i)
	}

	solution = append(solution, 0)     

	for len(notVisitedNodes)!=0 {
		nodeToAddIndex := rand.Intn(len(notVisitedNodes))
		solution = append(solution, notVisitedNodes[nodeToAddIndex])
		notVisitedNodes[nodeToAddIndex] = notVisitedNodes[len(notVisitedNodes)-1]
		notVisitedNodes = notVisitedNodes[:len(notVisitedNodes)-1]
	}

	return solution
}

func (g *Graph) GetCycleFromUser() []int {
	var nodeToAdd int
	var solution []int
	var nodes = map[int]bool{}

	for i:=0; i<g.Size; i++ {
		nodes[i] = false
	}

	for len(solution) < g.Size {
		fmt.Println("Enter node to add to cycle:")
		fmt.Println("Choose one of not visited nodes (false status)")
		fmt.Println(nodes)
		fmt.Scan(&nodeToAdd)
		if !nodes[nodeToAdd] {
			solution = append(solution, nodeToAdd)
			nodes[nodeToAdd] = true
		} else {
			fmt.Println("This node has been already added to the cycle!")
		}
	}
	return solution
}

func (g *Graph) TargetFunction(nodes []int) int {
	var result int
	last := nodes[0]
	for _, node := range(nodes[1:]) {
		result = result + g.Values[last][node]
		last = node;
	}
	result = result + g.Values[nodes[len(nodes)-1]][nodes[0]]
	return result;
}

