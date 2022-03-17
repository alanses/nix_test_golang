package cmd

import (
	"nix_game_of_life/signatures"
)

func StartGame() {
	consoleReader := signatures.NewConsoleReader()

	gameSettings := &signatures.GameSettings{Width: consoleReader.Width, Height: consoleReader.Height, Iterations: consoleReader.CountIteration}

	gameMap := &signatures.MapGameBord{
		MapData:             consoleReader.GetMatrix(),
		Width:               gameSettings.GetWidth(),
		Height:              gameSettings.GetHeight(),
		GameValidationPoint: &signatures.GameValidationPoint{Width: gameSettings.GetWidth(), Height: gameSettings.GetHeight()},
	}

	iterator := &signatures.IterationSignature{
		CountIteration: gameSettings.GetCountIterations(),
		MapGameBord:    gameMap,
	}

	iterator.Iterate()
}
