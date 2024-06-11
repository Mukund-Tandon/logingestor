package collectors

import "logingrestor/pkg/models"


type Collector interface {
	Start()
	Stop()
} 

func NewHTTPCollector(logbufferChannel chan models.Log) *HttpCollector {
    return &HttpCollector{
		logbufferChannel: logbufferChannel,
	}
}