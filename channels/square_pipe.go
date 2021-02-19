package channels

import (
	"fmt"
	"time"

	"github.com/haustov/simple"
)

// ResultsChannel - results
type ResultsChannel chan int

// ComputeSquare - compute
func ComputeSquare(input GenerationChannel) ResultsChannel {
	ch := make(ResultsChannel, 1)

	go func() {
		for x := range input {
			sq := simple.Square(x)
			fmt.Println("*** square compted ", sq)
			ch <- sq
			time.Sleep(time.Second / 2)
		}
		close(ch)
	}()

	return ch
}
