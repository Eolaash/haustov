package channels

import (
	"fmt"
	"sync"
	"time"
)

// FinalChannel -
//type FinalChannel chan struct{}

// PrinterPipe -pipe for print
func PrinterPipe(input ResultsChannel) {
	//ch := make(FinalChannel)
	var wg sync.WaitGroup

	wg.Add(1)

	go func() {
		for x := range input {
			fmt.Println("=== print result ", x)
			time.Sleep(time.Second)
		}
		wg.Done()
		//ch <- struct{}{}
	}()

	wg.Wait()
	//return ch
}
