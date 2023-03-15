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
	ReceiverID string `json:"receiverID"`
	ActorID    string `json:"actorID"`
	SequenceID int    `json:"sequenceID"`
	Epoche     int    `json:"epoche"`
	ActionID   string `json:"actionID"`
	SenderID   string `json:"senderID"`
}

type AdminRequest struct {
	Request Request `json:"request"`
	Admin   Admin   `json:"admin"`
}

type Record struct {
	ReceiverID string `json:"receiverID"`
	ActorID    string `json:"actorID"`
	SequenceID int    `json:"sequenceID"`
	Epoche     int    `json:"epoche"`
	ActionID   string `json:"actionID"`
	SenderID   string `json:"senderID"`
}

type AdminRecord struct {
	Record Record `json:"record"`
	Admin  Admin  `json:"admin"`
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

func TestM01() {
	var err error
	var jsonBlob []byte

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

	fmt.Printf(" testM01: length %v %v\n", len(dat1), dat1)

	fmt.Printf(" testM01: %v\n", string(jsonBlob))

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
