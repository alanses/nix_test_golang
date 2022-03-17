package signatures

type GameValidationPoint struct {
	Width  int
	Height int
}

func (gameValidationPoint *GameValidationPoint) GetCountAliveNeighbours(x int, y int, data [][]bool) int {
	aliveNeighbours := 0

	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if (gameValidationPoint.notIncludeSelf(i, j)) && (gameValidationPoint.getPoint(x+i, y+j, data)) {
				aliveNeighbours++
			}
		}
	}

	return aliveNeighbours
}

func (gameValidationPoint *GameValidationPoint) CheckIfPointAliveAndHasLessTwoAliveNeighbours(x int, y int, data [][]bool, countAliveNeighbours int) bool {
	return gameValidationPoint.isPointAlive(x, y, data) && countAliveNeighbours < 2
}

func (gameValidationPoint *GameValidationPoint) CheckIfPointIsAliveAndHasMoreThenTreeNeighbours(x int, y int, data [][]bool, countAliveNeighbours int) bool {
	return gameValidationPoint.isPointAlive(x, y, data) && countAliveNeighbours > 3
}

func (gameValidationPoint *GameValidationPoint) CheckIfPointIsAliveAndHasTwoOrTreeNeighbours(x int, y int, data [][]bool, countAliveNeighbours int) bool {
	return gameValidationPoint.isPointAlive(x, y, data) && (countAliveNeighbours == 2 || countAliveNeighbours == 3)
}

func (gameValidationPoint *GameValidationPoint) CheckIfPointIsDeadAndHasTreeNeighbors(x int, y int, data [][]bool, countAliveNeighbours int) bool {
	return !gameValidationPoint.isPointAlive(x, y, data) && countAliveNeighbours == 3
}

func (gameValidationPoint *GameValidationPoint) isPointAlive(x int, y int, data [][]bool) bool {
	return data[y][x]
}

func (gameValidationPoint *GameValidationPoint) notIncludeSelf(x int, y int) bool {
	return x != 0 || y != 0
}

func (gameValidationPoint *GameValidationPoint) getPoint(x int, y int, data [][]bool) bool {
	x = gameValidationPoint.getPointForX(x)
	y = gameValidationPoint.getVitalityForY(y)

	return data[x][y]
}

func (gameValidationPoint *GameValidationPoint) getPointForX(x int) int {
	x += gameValidationPoint.Height
	x %= gameValidationPoint.Height

	return x
}

func (gameValidationPoint *GameValidationPoint) getVitalityForY(y int) int {
	y += gameValidationPoint.Width
	y %= gameValidationPoint.Width

	return y
}
