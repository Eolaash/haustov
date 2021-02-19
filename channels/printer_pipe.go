package channels

import (
	"fmt"
	"time"
)

// FinalChannel -
type FinalChannel chan struct{}

// PrinterPipe -pipe for print
func PrinterPipe(input ResultsChannel) FinalChannel {
	ch := make(FinalChannel)

	go func() {
		for x := range input {
			fmt.Println("=== print result ", x)
			time.Sleep(time.Second)
		}

		ch <- struct{}{}
	}()

	return ch
}
