package signatures

import (
	"fmt"
	"time"
)

type IterationSignature struct {
	CountIteration int
	MapGameBord    *MapGameBord
}

func (iterationSignature *IterationSignature) Iterate() {
	iterations := 0

	for iterations < iterationSignature.CountIteration {
		game := iterationSignature.MapGameBord.NextGeneration()

		fmt.Println(game.String())

		time.Sleep(1 * time.Second)
		iterations++
	}
}
