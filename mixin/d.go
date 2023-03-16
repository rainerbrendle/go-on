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

type DbSession struct {
	conf DbConfig
	db   *sql.DB
}

var session DbSession

func DbTests() {

	defer DBClose(&session)

	TestDB(&session)
	ActivateDB(&session)
	// ...
}

func TestDB(s *DbSession) {

	fmt.Printf("  database test %v\n", s.conf.active)
	if s.conf.active == false {
		return
	}

}

func ActivateDB(s *DbSession) {

	s.conf.active = false
	s.conf.dbInfo = DbInfo()

	DBConnect(s)

	fmt.Printf("  database info after connect \n...: %v tested %v\n", s.conf.dbInfo, s.conf.tested)

}

func DBConnect(s *DbSession) {

	fmt.Printf("database open connection\n ...:%v\n", s.conf.dbInfo)

	db, err := sql.Open("postgres", s.conf.dbInfo)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	s.db = db

	fmt.Println(" ping ....")
	err = s.db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")

	s.conf.active = true
	s.conf.tested = true
}

func DBClose(s *DbSession) {

	if s == nil {
		return
	}

	fmt.Println(" ... closing")

	if s.db != nil {
		s.db.Close()
	}

	s.conf.tested = false
	s.conf.active = false

}
