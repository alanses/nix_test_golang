package signatures

import (
	"math/rand"
	"time"
)

type MapDataGenerator struct{}

func NewMapGenerator() *MapDataGenerator {
	return &MapDataGenerator{}
}

func (mapDataGenerator *MapDataGenerator) GenerateRandomMap(width int, height int) [][]bool {
	mapGeneration := make([][]bool, width)

	for x := 0; x < height; x++ {
		mapGeneration[x] = []bool{}
		for y := 0; y < width; y++ {
			rand.Seed(time.Now().UnixNano())

			randBoolValue := rand.Intn(2) == 0

			mapGeneration[x] = append(mapGeneration[x], randBoolValue)
		}
	}

	return mapGeneration
}

func (mapDataGenerator *MapDataGenerator) GenerateTestMap() [][]bool {
	return [][]bool{
		{true, false, true, false, false, false, false, false, false, false},
		{true, true, true, true, true, false, false, false, false, false},
		{false, false, false, false, true, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
	}
}
