package channels

import (
	"fmt"
	"time"

	"github.com/haustov/generator"
)

// GenerationChannel - channel
type GenerationChannel chan int

// GenerateInts - generate ints
func GenerateInts(k int) GenerationChannel {
	ch := make(chan int, 1)
	go func() {
		for i := 0; i < k; i++ {
			x := generator.GenerateInt()
			fmt.Println("+++ generated ", x)
			ch <- x
			time.Sleep(time.Second / 3)
		}
		close(ch)
	}()

	return ch
}
