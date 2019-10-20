package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
	"math/rand"
)

type Graph struct {
	Name string
	Size int
	Values [][]int
}

func (g Graph) String() string {
	return fmt.Sprintf("Name: %vSize: %v\nValues: %v", g.Name, g.Size, g.Values)
}

func (g *Graph) CreateFromFile(fileName string) Graph {
	
	file, err := os.Open(fileName)
	if err!= nil {
		fmt.Print(err)
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
	return *g
}


func (g *Graph) CreateRandom(size int, name string) Graph {
	g.Size = size
	g.Values = make([][]int, g.Size)
	
	for index:=0; index<g.Size; index++ {
		row  := make([]int, g.Size)
		for index2:=0; index2<g.Size; index2++ {
			row[index2] = rand.Intn(490) + 10
		}
		g.Values[index] = row
	}
	
	return *g
}
