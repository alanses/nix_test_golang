package signatures

import (
	"flag"
	"strconv"
	"strings"
)

type ConsoleReader struct {
	CountIteration int
	Width          int
	Height         int
	Matrix         [][]bool
}

func NewConsoleReader() *ConsoleReader {
	countOfIterations := flag.Int("count_iterations", 999, "count iterations into program")
	width := flag.Int("width", 999, "count width of matrix into program")
	height := flag.Int("height", 999, "count height of matrix into program")
	matrix := flag.String("matrix", "", "matrix by using")

	flag.Parse()

	return &ConsoleReader{
		CountIteration: *countOfIterations,
		Width:          *width,
		Height:         *height,
		Matrix:         ReadInputMatrix(*matrix, *width, *height),
	}
}

func (consoleReader *ConsoleReader) GetMatrix() [][]bool {
	return consoleReader.Matrix
}

func ReadInputMatrix(matrixString string, width int, height int) [][]bool {
	mapGenerator := NewMapGenerator()

	if matrixString == "" || matrixString == "default" {
		return mapGenerator.GenerateTestMap()
	}

	if matrixString == "rand" {
		return mapGenerator.GenerateRandomMap(width, height)
	}

	inputs := strings.FieldsFunc(matrixString, func(r rune) bool {
		return r == '[' || r == ']'
	})

	mapGeneration := make([][]bool, width)

	for i := 0; i < len(inputs); i++ {
		col := strings.Split(inputs[i], ",")
		mapGeneration[i] = []bool{}
		for j := 0; j < len(col); j++ {
			item, err := strconv.Atoi(col[j])

			if err != nil {
				panic("parse err")
			}

			if item == 0 {
				mapGeneration[i] = append(mapGeneration[i], false)
			} else {
				mapGeneration[i] = append(mapGeneration[i], true)
			}
		}
	}

	return mapGeneration
}
