package main

import (
	"os"
	"fmt"
	"github.com/ziutek/mymysql/mysql"
	_ "github.com/ziutek/mymysql/native"
)

// MySQL address
var address = "127.0.0.1:3306"

// Main Function
func main() {
	err := checkArgs()
	if err != 0 {
		return
	}
	user := os.Args[1]
	pass := os.Args[2]
	dbname := os.Args[3]
	table := os.Args[4]
	db, err := connect(user, pass, dbname)
	if err != 0 {
		return
	}
	rows, err := executeQuery(db, table)
	if err != 0 {
		return
	}
	output(rows)
	return
}

// Check Args count
func checkArgs() (int) {
	if len(os.Args) != 5 {
		fmt.Println("Argument Error:")
		fmt.Println("usage:")
		fmt.Println(" Need Three Arguments.")
		fmt.Println("  1st:Database User Name")
		fmt.Println("  2nd:Database User Password")
		fmt.Println("  3rd:Use Database Name")
		fmt.Println("  4th:Table Name")
		return -1
	}
	return 0
}

// Connection MySQL
func connect(user string, pass string, dbname string) (mysql.Conn, int) {
	db := mysql.New("tcp", "", address, user, pass, dbname)
	err := db.Connect()
	if err != nil {
		fmt.Println("MySQL Connection Error:")
		fmt.Println("=== Detail ===")
		panic(err)
		return nil, -1
	}
	return db, 0
}

// Execute Query
func executeQuery(db mysql.Conn, table string) ([]mysql.Row, int) {
	query := "select * from " + table
	rows, _, err := db.Query(query)
	if err != nil {
		fmt.Println("MySQL Query Error:")
		fmt.Println("=== Detail ===")
		panic(err)
		return nil, -1
	}
	return rows, 0
}

// Output Query Result
func output(rows []mysql.Row) {
	for _, row := range rows {
		str := ""
		for _, col := range row {
			if col != nil {
				str += string(col.([]byte))
			}
			str += "|\t"
		}
		fmt.Println(str)
	}
	return
}
