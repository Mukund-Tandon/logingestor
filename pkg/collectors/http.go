package collectors

import (
	"fmt"
	"logingrestor/pkg/models"
	"logingrestor/pkg/transformers"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HttpCollector struct{
	logbufferChannel chan models.Log
}


func (c *HttpCollector) Start() error {
	fmt.Println("Starting HTTP Collector")
    
	router := gin.Default()
	router.POST("/log", c.handleLogs)

	// Start the HTTP server and block until it exits
	err := router.Run("localhost:8080")
	if err != nil {
		fmt.Println("Error starting HTTP server:", err)
		return err
	}

	return nil
}
func (c *HttpCollector) Stop() {
	fmt.Println("Stopped HTTP Collector")
} 
func (c *HttpCollector) handleLogs(ctx *gin.Context) {
	log, err := transformer.HttpToLog(ctx)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// Send the log to the log buffer channel for further processing
	c.logbufferChannel <- log

	ctx.Status(http.StatusOK)
}

// func handleLogs(ctx *HttpCollector) (c *gin.Context) {
//     log ,err := transformer.HttpToLog(c)

// 	if err != nil {
// 		c.AbortWithError(http.StatusBadRequest, err)
//         return
// 	}



// 	fmt.Println(log)
//     // Here, you can parse the log data from the request body
//     // and send it to the log channel for further processing

//     c.Status(http.StatusOK)
// }