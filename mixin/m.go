package mixin

import (
	"encoding/json"
	"fmt"
)

type Admin2 struct {
	ID         int    `json:"id"`
	ShardID    string `json:"shardID"`
	VersionID  string `json:"versionID"`
	ValidFrom  string `json:"validFrom"`
	OrgName    string `json:"orgName"`
	DbTypeName string `json:"dbTypeName"`
	Replica    int    `json:"replica"`
}

type Request struct {
	ReceiverID string `json:"receiverID"`
	ActorID    string `json:"actorID"`
	SequenceID int    `json:"sequenceID"`
	Epoch      int    `json:"epoche"`
	ActionID   string `json:"actionID"`
	SenderID   string `json:"senderID"`
}

type AdminRequest struct {
	Request Request `json:"request"`
	Admin   Admin2  `json:"admin"`
}

type Record struct {
	ReceiverID string `json:"receiverID"`
	ActorID    string `json:"actorID"`
	SequenceID int    `json:"sequenceID"`
	Epoch      int    `json:"epoch"`
	ActionID   string `json:"actionID"`
	SenderID   string `json:"senderID"`
        Reason     string `json:"reason"`
        DateTime   string `json:"dateTime"`
        Signature  string `json:"signature"`
}

type AdminRecord struct {
	Record Record `json:"record"`
	Admin  Admin2 `json:"admin"`
}

type A1 struct {
	Astr1 string `json:"actor"`
	Astr2 string `json:"sequence"`
}

type B1 struct {
	ID string `json"node"`
}

type Comp struct {
	Astr string `json:"action"`
	A    A1     `json:"message"`
	B    B1     `json:"content"`
}

var a1dat []AdminRequest
var a2dat []AdminRecord

func TestInit() {

	a1dat = make([]AdminRequest, 2)
	m := Request{
		ReceiverID: "0",
		ActorID:    "YellowPages",
		SequenceID: 1,
		Epoch:      0,
		ActionID:   "create",
		SenderID:   "client1"}
	a1 := Admin2{
		ID:         1,
		ShardID:    "Control",
		VersionID:  "1.1.0.1",
		ValidFrom:  "2023-03-01",
		OrgName:    "IT",
		DbTypeName: "postgres",
		Replica:    3}

	var a1buffer AdminRequest

	a1buffer.Request = m
	a1buffer.Admin = a1
	a1dat[0] = a1buffer

	m.SequenceID = 2
	m.ActionID = "modify"
	a1.ID = 2
	a1buffer.Admin = a1
	a1buffer.Request = m
	a1dat[1] = a1buffer

	fmt.Println("Initialize Mixin capable test data")

}

func TestM01() {
	var err error
	var jsonBlob []byte

	// testing : dynamic creation of composite slice

	// dat1 := []Comp{ "IT","Sales"}}
	var dat1 []Comp

	dat1 = make([]Comp, 2)

	dat1[0].Astr = "create"
	dat1[0].A.Astr1 = "Control"
	dat1[1].Astr = "modify"
	dat1[1].A.Astr1 = "Control"
	dat1[0].B.ID = "YP"
	dat1[1].B.ID = "YP"

	jsonBlob, err = json.Marshal(dat1)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf(" testM01: length %v\n %v\n", len(dat1), dat1)

	fmt.Printf(" testM01:\n more%v\n", string(jsonBlob))

	jsonBlob, err = json.Marshal(a1dat)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf(" testM01: length %v\n %v\n", len(a1dat), a1dat)

	fmt.Printf(" testM01:\n %v\n", string(jsonBlob))

}

// we can combine structures into bigger ones
// We can define
//
//   Message
//   Request
//   Record
//   View
//
//   base classes and add payload with equal or similiar data
//
//   It comes with a price. The base classes are visible.
