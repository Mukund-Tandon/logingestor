package buffer

import (
	"fmt"
	"logingrestor/pkg/models"
	"time"
)

func LogBuffer( logBatchOutputChannel chan models.Logbatch) chan models.Log {
	logChannel := make(chan models.Log)
	buffer := make([]models.Log, 0, 100)
	ticker := time.NewTicker(15 * time.Second)

	go func() {
		for {
			select {
			case log := <-logChannel:
				buffer = append(buffer, log)
				if len(buffer) >= 100 {
					logBatchOutputChannel <- models.Logbatch{Logbatch: buffer}
					buffer = buffer[:0] // Clear the buffer
				}
			case <-ticker.C:
				if len(buffer) > 0 {
					// output(buffer)
					fmt.Println("Buffer tome out")
					logBatchOutputChannel <- models.Logbatch{Logbatch: buffer}
					buffer = buffer[:0] // Clear the buffer
				}
				ticker.Reset(5 * time.Second)
			}
		}
	}()

	return logChannel
}