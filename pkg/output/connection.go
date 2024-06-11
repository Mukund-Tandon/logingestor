package output

import (
	"fmt"

	"github.com/ClickHouse/clickhouse-go/v2"
)

func getConnection() (clickhouse.Conn,error){
	// Connect to the database
	fmt.Println("Connecting to Database..")
	conn, err := clickhouse.Open(&clickhouse.Options{
		Addr: []string{"127.0.0.1:9000"}, // Corrected to a slice
		Auth: clickhouse.Auth{
			Database: "testdb",
			Username: "default",
			Password: "", // Make sure this matches your setup
		},
	})

	if err!= nil {
		fmt.Println("Error connecting to ClickHouse:", err)
		return nil,err
	}
	return conn,nil
}