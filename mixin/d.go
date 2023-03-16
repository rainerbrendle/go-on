package mixin

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type DbConfig struct {
	dbInfo string
	tested bool
	active bool
}

// var conf DbConfig( active: false )
var conf DbConfig
var db *sql.DB

func TestDB() {

	fmt.Printf("  database test %v\n", conf.active)
	if conf.active == false {
		return
	}

}

func ActivateDB() {

	conf.active = false
	conf.dbInfo = DbInfo()

	DBConnect()

	fmt.Printf("  database info %v\n", conf.dbInfo)

}

func DBConnect() {

	fmt.Printf("database open connection\n %v\n", conf.dbInfo)

	db, err := sql.Open("postgres", conf.dbInfo)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer DBClose()

	fmt.Println(" ping ....")
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")

	conf.active = true
	conf.tested = true
}

func DBClose() {

	if db != nil {
		db.Close()
	}

	conf.tested = false
	conf.active = false

}
