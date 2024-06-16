package output

import (
	"fmt"
	"os"
	"github.com/ClickHouse/clickhouse-go/v2"
)

func getConnection() (clickhouse.Conn,error){
	// Connect to the database
	fmt.Println("Connecting to Database..")
	host := os.Getenv("CLICKHOUSE_HOST")
	port := os.Getenv("CLICKHOUSE_PORT")
	user := os.Getenv("CLICKHOUSE_USER")
	password := os.Getenv("CLICKHOUSE_PASSWORD")
	fmt.Println("ENV Vars")
	fmt.Println(host)
	fmt.Println(port)
	fmt.Println(user)
	fmt.Println(password)
	conn, err := clickhouse.Open(&clickhouse.Options{
		Addr: []string{fmt.Sprintf("%s:%s", host, port)}, // Corrected to a slice
		Auth: clickhouse.Auth{
			Database: "testdb",
			Username: user,
			Password: password, // Make sure this matches your setup
		},
	})

	if err!= nil {
		fmt.Println("Error connecting to ClickHouse:", err)
		return nil,err
	}
	return conn,nil
}