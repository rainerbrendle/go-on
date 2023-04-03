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

	fmt.Printf("Database related testing ...\n")

	// close database connection when function returns
	defer DBClose(&session)

	TestDB(&session)

	ActivateDB(&session)

	CreateTestData(&session)

	NotifyDB(&session)
	AgentLoopDB(&session)
	ReadFromDB(&session)

}

func CreateTestData(s *DbSession) {
	fmt.Printf("  CreateTestData: create test data .. running %v\n", s.conf.tested)

	// DROP TEST DATA
	dropTestData(s)
	createTestData(s)

	// INSERTED TEST DATA
}

func dropTestData(s *DbSession) {
	//  statement := `DROP TABLE PR2.SenderJournal;`
	statement := `SELECT PR2.DoNothing();`

	fmt.Printf("  DropTestData 1: statement %v\n", statement)

	result, err := s.db.Exec(statement)

	if err != nil {
		panic(err)
	}

	var dest string
	var seq int

	statement = `CALL PR2.Echo( $1, $2, $3, $4);`
	fmt.Printf("  DropTestData 2: statement %v\n", statement)

	err = s.db.QueryRow(statement, "XXX", 100, nil, nil).Scan(&dest, &seq)
	fmt.Printf("  Echo: %v : %v \n", dest, seq)

	if err != nil {
		panic(err)
	}

	statement = `DELETE FROM PR2.SenderJournal WHERE SenderID = $1;`
	result, err = s.db.Exec(statement, "TEST")
	fmt.Printf("  DropTestData 3: statement %v result %v\n", statement, result)

	if err != nil {
		panic(err)
	}

	return

}

func createTestData(s *DbSession) {
	fmt.Printf("  Creating test data\n")

	for i := 0; i < 2; i++ {
		fmt.Printf("  %v: \n", i)

	}
}

func NotifyDB(s *DbSession) {
	fmt.Printf("  NotifyDB: try notify %v\n", s.conf.active)
	if s.conf.active == false {
		return
	}

}

func AgentLoopDB(s *DbSession) {
	fmt.Printf("  AgentLoopDB: agent .. %v\n", s.conf.active)
	if s.conf.active == false {
		return
	}

}

func ReadFromDB(s *DbSession) {
	fmt.Printf("  ReadFromDB: reading .. %v\n", s.conf.active)
	if s.conf.active == false {
		return
	}

}

func TestDB(s *DbSession) {

	fmt.Printf("  TestDB: database test %v\n", s.conf.active)
	if s.conf.active == false {
		return
	}

}

func ActivateDB(s *DbSession) {

	s.conf.active = false
	s.conf.dbInfo = DbInfo()

	DBConnect(s)

	fmt.Printf("  AcivateDB: database: info after connect \n...: %v ping tested %v\n", s.conf.dbInfo, s.conf.tested)

}

func DBConnect(s *DbSession) {

	fmt.Printf("database: open connection\n ...: %v\n", s.conf.dbInfo)

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
