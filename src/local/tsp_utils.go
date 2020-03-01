package local

import (
	"bufio"
	"errors"
	"os"
	"reflect"
	"strconv"
	"strings"
)

var floatType = reflect.TypeOf(float64(0))
var stringType = reflect.TypeOf("")

func LoadAdjacencyMatrixFromFile(fileName string) ([][]int, error) {

	file, err := os.Open(fileName)
	if err != nil {
		return [][]int{}, errors.New("Could not open file")
}

	reader := bufio.NewReader(file)
	reader.ReadString('\n')
	sizeString, err := reader.ReadString('\n')
	size, err := strconv.Atoi(strings.TrimSpace(sizeString))
	adjacencyMatrix := make([][]int, size)

	for index := 0; index < size; index++ {
		line, _ := reader.ReadString('\n')
		lineOfValues := strings.Fields(line)

		row := make([]int, size)
		for index2, element := range lineOfValues {
			value, _ := strconv.Atoi(strings.TrimSpace(element))
			row[index2] = value
		}
		adjacencyMatrix[index] = row
	}

	defer file.Close()
	return adjacencyMatrix, nil
}