package mixin

import (
	"crypto/sha256"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

type DbConfig struct {
	dbInfo string
	tested bool
	active bool
}

type DbSession struct {
	conf       DbConfig
	db         *sql.DB
	prefixCode string
}

type testRecord struct {
	record Record
	json   string
}

/*
type Database struct {
  Url string
}

func NewDatabase(url string) Database {
  return Database{url}
}

func main() {
  db := NewDatabase("localhost:5432")

  http.HandleFunc("/hello", hello(db))
  http.ListenAndServe(":3000", nil)
}

func hello(db Database) func(http.ResponseWriter, *http.Request) {
  return func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, db.Url)
  }
}
*/

func NewActor(url string) *Record {
	rec := new(Record)

	/*  RecordFromURL{url, &rec } */

	return rec
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

	fmt.Printf("  DropTestData 1: SELECT FUNCTION: statement %v\n", statement)

	result, err := s.db.Exec(statement)

	if err != nil {
		panic(err)
	}

	var dest string
	var seq int

	statement = `CALL PR2.Echo( $1, $2, $3, $4);`
	fmt.Printf("  DropTestData 2: CALL PROCEDURE: statement %v\n", statement)

	err = s.db.QueryRow(statement, "XXX", 100, nil, nil).Scan(&dest, &seq)
	fmt.Printf("  Echo: XXX/100 =? %v/%v \n", dest, seq)

	if err != nil {
		panic(err)
	}

	statement = `DELETE FROM PR2.SenderJournal WHERE SenderID = $1;`
	result, err = s.db.Exec(statement, "TEST")
	fmt.Printf("  DropTestData 3: statement %v result %T\n", statement, result)

	if err != nil {
		panic(err)
	}

	return

}

func recordSlice(record *testRecord) {

}

func createTestData(s *DbSession) {
	fmt.Printf("  Creating test data\n")

	var tr testRecord

	tm := time.Now()

	tmd := tm.Format(time.RFC3339)

	tr.record.ActorID = "admin"
	tr.record.ReceiverID = "IT"
	tr.record.Reason = "0"
	tr.record.DateTime = tmd
	tr.record.ActionID = "create"
	tr.json = `[{"id":101}]`

	fmt.Printf("%v\n", time.Now().Format(time.RFC3339))

	for i := 0; i < 2; i++ {
		fmt.Printf("  %v: %v %v\n", i, tr.record.DateTime, PrefixCode)

		sign(&tr)
		insertTestRecord(s, &tr)
	}

	fmt.Printf("%v\n", time.Now().Format(time.RFC3339))

}

func sign(tr *testRecord) {
	var str string

	str = tr.json + tr.record.DateTime + PrefixCode // seed

	hash := sha256.Sum256([]byte(str))
	fmt.Printf("SHA256 Hash code with Prefix for\n %v:\n %x\n", str, hash)

	/* sign record */
	tr.record.Signature = string(hash[:])
}

func insertTestRecord(s *DbSession, tr *testRecord) {
	statement := `CALL PR2.TestData( $1, $2, $3, $4, $5, $6, $7);`

	converted := fmt.Sprintf("%x", tr.record.Signature)
	fmt.Printf("  signature: %v\n", converted)

	_, err := s.db.Exec(statement,
		tr.record.ActorID, tr.record.ReceiverID, tr.record.Reason, tr.record.DateTime,
		tr.record.ActionID, tr.json, converted)

	if err != nil {
		panic(err)
	}

	fmt.Printf("   %v %v \n    %v %v\n     %v %v\n    %v\n", tr.record.ActorID, tr.record.ReceiverID,
		tr.record.Reason, tr.record.DateTime, tr.record.ActionID, tr.json, converted)

	/*	CREATE OR REPLACE PROCEDURE PR2.TestData(IN _actor text,
			IN _receiver text,
			IN _reason text,
			IN _dateTime text,
			IN _action text,
			IN _json text,
			IN _signature text
		   )
	*/

	/* (not supported by pg driver)
	   lastId, err  := result.LastInsertId()
	    if err != nil {
		   panic(err)
	    }
	*/
	fmt.Printf("  TestData: CALL PROCEDURE: statement %v\n", statement)

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
	s.prefixCode = "SE3D"
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
