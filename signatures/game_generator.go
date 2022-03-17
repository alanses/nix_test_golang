package signatures

import "bytes"

type MapGameBord struct {
	MapData             [][]bool
	Width               int
	Height              int
	GameValidationPoint *GameValidationPoint
}

func (mapGameBord MapGameBord) NextGeneration() MapGameBord {
	for y, rows := range mapGameBord.MapData {
		for x, _ := range rows {
			aliveNeighbours := mapGameBord.GameValidationPoint.GetCountAliveNeighbours(x, y, mapGameBord.MapData)
			mapGameBord.MapByRules(x, y, aliveNeighbours)
		}
	}

	return mapGameBord
}

func (mapGameBord *MapGameBord) MapByRules(x int, y int, aliveNeighbours int) {
	if mapGameBord.GameValidationPoint.CheckIfPointAliveAndHasLessTwoAliveNeighbours(x, y, mapGameBord.MapData, aliveNeighbours) {
		mapGameBord.MapData[y][x] = true
	}

	if mapGameBord.GameValidationPoint.CheckIfPointIsAliveAndHasMoreThenTreeNeighbours(x, y, mapGameBord.MapData, aliveNeighbours) {
		mapGameBord.MapData[y][x] = false
	}

	if mapGameBord.GameValidationPoint.CheckIfPointIsAliveAndHasTwoOrTreeNeighbours(x, y, mapGameBord.MapData, aliveNeighbours) {
		mapGameBord.MapData[y][x] = true
	}

	if mapGameBord.GameValidationPoint.CheckIfPointIsDeadAndHasTreeNeighbors(x, y, mapGameBord.MapData, aliveNeighbours) {
		mapGameBord.MapData[y][x] = true
	}
}

func (mapGameBord *MapGameBord) String() string {
	var buf bytes.Buffer

	for _, col := range mapGameBord.MapData {
		for _, cell := range col {
			b := byte(' ')
			if cell {
				b = '1'
			} else {
				b = '0'
			}
			buf.WriteByte(b)
		}
		buf.WriteByte('\n')
	}

	return buf.String()
}
