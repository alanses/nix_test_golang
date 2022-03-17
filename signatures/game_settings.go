package signatures

type GameSettings struct {
	Width      int
	Height     int
	Iterations int
}

func (gameSettings *GameSettings) GetWidth() int {
	return gameSettings.Width
}

func (gameSettings *GameSettings) GetHeight() int {
	return gameSettings.Height
}

func (gameSettings *GameSettings) GetCountIterations() int {
	return gameSettings.Iterations
}
