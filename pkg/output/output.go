package output

import (
	"context"
	"fmt"
	"logingrestor/pkg/models"
	"time"

	"github.com/ClickHouse/clickhouse-go/v2"
)

func Output() (chan models.Logbatch ,error ) {
	conn,err := getConnection()
	if err != nil{
		fmt.Println("Databae connection failed")
		return nil,err
	}
	fmt.Println(conn)
	logbatchChannel := make(chan models.Logbatch)
    go func() {
		for {
			select {
				case logbatch := <-logbatchChannel:
					// output(logbatch)
					fmt.Println("Output")
					fmt.Println(logbatch)
					doBatchInsert(logbatch, conn)
			}
		}
	}()

	return logbatchChannel,nil

}

func doBatchInsert(logbatch models.Logbatch, conn clickhouse.Conn) {
    ctx := context.Background()
    batch, err := conn.PrepareBatch(ctx, "INSERT INTO logs")
    if err!= nil {
        fmt.Println("Error preparing batch:", err)
        return // Ensure we don't proceed with a failed batch preparation
    }
    
    logBatchSize := len(logbatch.Logbatch)
    for i := 0; i < logBatchSize; i++ {
        timestampStr := logbatch.Logbatch[i].Timestamp
        timestamp, err := time.Parse(time.RFC3339, timestampStr)
		fmt.Println("Timestamp")
		fmt.Println(timestamp)
        if err!= nil {
            fmt.Printf("Error parsing timestamp '%s': %v\n", timestampStr, err)
            continue // Skip this entry and move to the next one
        }

        clickHouseFormatTimeStamp := timestamp.Format("2006-01-02 15:04:05")
        message := logbatch.Logbatch[i].Message
        level := logbatch.Logbatch[i].Level
        resourceID := logbatch.Logbatch[i].ResourceID
        
        err = batch.Append(
            clickHouseFormatTimeStamp,
            message,
            level,
            resourceID,
        )
        if err!= nil {
            fmt.Println("Error executing query:", err)
            return // Stop execution on query error
        }
    }
    err = batch.Send() // Pass the context to Send
    if err!= nil {
        fmt.Println("Error sending batch:", err)
    }
}

	// rows, err := conn.Query(ctx, "select name from testdb.test_table")
	// if err!= nil {
	// 	fmt.Println("Error executing query:", err)
	// 	return
	// }
	// defer rows.Close()

	// for rows.Next() {
	// 	var name string
	// 	err := rows.Scan(&name)
	// 	if err!= nil {
	// 		fmt.Println("Error scanning row:", err)
	// 		continue // Continue to the next iteration instead of returning
	// 	}
	// 	fmt.Printf("row: col1=%s\n", name) // Corrected printf format specifier
	// }

	// if err := rows.Err(); err!= nil {
	// 	fmt.Println("Error iterating over rows:", err)
	// }