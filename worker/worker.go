package worker

import (
	"fmt"
	"time"
)

// Long running task running for x seconds
const timeSeconds = 10

// Worker - Worker to process NATS message
func Worker(s string) {
	fmt.Printf("Received a message: %s\n", s)

	time.AfterFunc(time.Second*timeSeconds, func() {
		fmt.Printf("Congratulations! Your %d second timer finished.", timeSeconds)
	})
	// defer timer.Stop()
}
