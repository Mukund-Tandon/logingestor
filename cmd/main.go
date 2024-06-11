package main

import (
	// "crypto/tls"
	"fmt"
	"sync"

	// "logingrestor/pkg/collectors"
	"logingrestor/pkg/buffer"
	"logingrestor/pkg/collectors"
	"logingrestor/pkg/output"
)

func main() {
	fmt.Println("Starting application...")
	
	logBatchOutputChannel,err := output.Output()
	if err != nil{
		fmt.Println(err)
	} 
    logChannel := buffer.LogBuffer(logBatchOutputChannel)
	fmt.Println(logChannel)
    
	var wg sync.WaitGroup
	
	httpcollector := collectors.NewHTTPCollector(logChannel)
	
	wg.Add(1) // Increment the counter by 1

	go func() {
		defer wg.Done() // Decrement the counter when the goroutine completes
		err := httpcollector.Start()
		if err != nil {
			fmt.Println("Error starting HTTP collector:", err)
		}
	}()

	wg.Wait() 
}

