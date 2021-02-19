package haustov

import (
	"fmt"

	"github.com/haustov/channels"
)

// VectorLength - length
const VectorLength = 5

// DeadLockTask TASK6 func
func DeadLockTask() {
	inputChannel := channels.GenerateInts(VectorLength)
	resultsChannel := channels.ComputeSquare(inputChannel)
	final := channels.PrinterPipe(resultsChannel)
	fmt.Println("Over")
	<-final
}
