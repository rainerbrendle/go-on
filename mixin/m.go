package mixin

import (
	"encoding/json"
	"fmt"
)

type Admin struct {
	ID         int    `json:"id"`
	ShardID    string `json:"shardID"`
	VersionID  string `json:"versionID"`
	ValidFrom  string `json:"validFrom"`
	OrgName    string `json:"orgName"`
	DbTypeNAme string `json:"dbTypeName"`
	Replica    int    `json:"replica"`
}

type Request struct {
	ReceiverID string
	ActorID    string
	SequenceID int
	Epoche     int
	ActionID   string
	SenderID   string
}

type AdminRequest struct {
	Request Request
	Admin   Admin
}

type Record struct {
	ReceiverID string
	ActorID    string
	SequenceID int
	Epoche     int
	ActionID   string
	SenderID   string
}

type AdminRecord struct {
	Record Record
	Admin  Admin
}

type A1 struct {
	Astr1 string `json:"receiver"`
	Astr2 string
}

type B1 struct {
	ID int
}

type Comp struct {
	Astr string `json:"department"`
	A    A1     `json:"node"`
}

func TestM01() {
	var err error
	var jsonBlob []byte

	// dat1 := []Comp{ "IT","Sales"}}
	var dat1 []Comp

	dat1 = make([]Comp, 2)
	dat1[0].Astr = "IT"
	dat1[0].A.Astr1 = "101"
	dat1[1].Astr = "Sales"

	jsonBlob, err = json.Marshal(dat1)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf(" testM01: length %v %v\n", len(dat1), dat1)

	fmt.Printf(" testM01: %v\n", string(jsonBlob))

}
