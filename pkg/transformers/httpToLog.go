package transformer

import (
	"io"
	"logingrestor/pkg/models"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
)

func HttpToLog(httprequest *gin.Context) (models.Log,error) {
	// clientIP := httprequest.ClientIP()
	var log models.Log
    requestBodyBytes, err := io.ReadAll(httprequest.Request.Body)
	if err != nil {
		return log, err
	}
    
    log.Timestamp = gjson.GetBytes(requestBodyBytes, "timestamp").String()
	log.Level = gjson.GetBytes(requestBodyBytes, "level").String()
	log.Message = gjson.GetBytes(requestBodyBytes, "message").String()
	log.ResourceID = gjson.GetBytes(requestBodyBytes, "resourceID").String()

	return log, nil

}