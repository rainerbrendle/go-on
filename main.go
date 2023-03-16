package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/json"
	"fmt"

	"go-on.com/mixin"
)

// global conversion buffer for testing
var stringBuffer bytes.Buffer

// Hash code sha256 for transport validation
// of stringBuffer
// It uses a standard seed, this needs to be configurable
// in a production environmant
// (we can use the same for encryption, would need a certificate for this, however)
var testHash [32]byte

const prefixCode = "#S33D"

type Admin struct {
	ID         int    `json:"id"`
	ShardID    string `json:"shardID"`
	VersionID  string `json:"versionID"`
	ValidFrom  string `json:"validFrom"`
	OrgName    string `json:"orgName"`
	DbTypeName string `json:"dbTypeName"`
	Replica    int    `json:"replica"`
}

// basic main function to run tests one after the other
func main() {

	fmt.Println("Hello to JSON mashaling testing")

	// primary mashaling test of simple data types
	test()
	// Marshaling test for business documents
	// (We had called them "Newdocs")

	// Mashaling
	test01()
	// Unmarshaling
	testUM01()

	// Initialize Mixin
	mixin.TestInit()

	// new mpdule
	mixin.TestM01()

	// database connections
	mixin.ActivateDB()
	mixin.TestDB()

}

func test() {
	// test 1
	// create basic marshaling

	fmt.Println("Marshal some basic data types")

	// Testing some types
	// Boolean
	bolB, _ := json.Marshal(true)
	fmt.Printf("boolen %v\n", string(bolB))
	// Integer
	intB, _ := json.Marshal(100)
	fmt.Printf("integer %v\n", string(intB))

	// Float
	fltB, _ := json.Marshal(2.34)
	fmt.Println(string(fltB))

	// String
	strB, _ := json.Marshal("test string")
	fmt.Printf("string %v \n", string(strB))

	// List
	slcD := []string{"aaa", "bbb", "ccc"}
	slcB, _ := json.Marshal(slcD)
	fmt.Printf("list %v \n", string(slcB))

}

func test01() {
	// marshal a table of structs
	var jsonBlob1 []byte
	var jsonBlob2 []byte

	// a byte array of a JSON array
	// we use a basic cluster configuration as an example
	// (this needs to be the edge points of a swim lane)
	jsonBlob1 = []byte(`[    
     {  
        "id": 1, 
        "shardID": "0",
        "versionID": "1.0.0.1",
        "validFrom": "2023-03-01",
        "orgName" : "IT",
        "dbTypeName" : "postgres",
        "replica": 3
     },
     {
        "id": 2,
        "shardID": "100",
        "versionID": "1.0.0.1",
        "validFrom": "2023-03-01",
        "orgName" : "Sales",
        "dbTypeName" : "postgres",
        "replica": 3
     }
]  `)

	// the same as a Go struct

	// initialize the Admin table in GO
	dat := []Admin{
		{
			1,
			"0",
			"1.0.0.1",
			"2023-03-01",
			"IT",
			"postgres",
			3},
		{
			2,
			"100",
			"1.0.0.1",
			"2023-03-01",
			"Sales",
			"postgres",
			3}}

	fmt.Println("marshal more ...")

	jsonBlob2, err := json.Marshal(dat)

	if err != nil {
		fmt.Println("error", err)
		panic("marshal")
	}

	fmt.Printf("input: \n GO table %v\n", dat)

	s := string(jsonBlob2)

	fmt.Printf("output: \n JSON byte array %v\n", s)

	// remove white space from JSON
	// we need a bytes.Buffer as intermediate buffer

	if err := json.Compact(&stringBuffer, jsonBlob1); err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Without white space:\n check %v\n", stringBuffer.String())

	// hash code
	var str string
	str = stringBuffer.String() + prefixCode // seed

	hash := sha256.Sum256([]byte(str))
	fmt.Printf("Hash for %v:\n %x\n", s, hash)
	testHash = hash

	// map to a generic slice (array)
	// copy  line by line
	var anything []interface{}
	for _, val := range dat {
		anything = append(anything, val)
	}

	fmt.Printf("Copied to generic slice\n: anything\n %v\n", anything)

}

func testUM01() {

	fmt.Println("unmarshal string buffer")

	var dat2 []Admin

	fmt.Printf(" %v\n", stringBuffer.String())

	// JSON and Go types must be compabile !!!
	// ints and strings ...
	err := json.Unmarshal(stringBuffer.Bytes(), &dat2)
	if err != nil {
		fmt.Println(err)
	}

}

func testUnMarshal() {

	var jsonBlob []byte

	var dat map[string]interface{}

	fmt.Println("unmarshal")

	err := json.Unmarshal(jsonBlob, &dat)

	jsonBlob = []byte(`{"YellowPages"
	{"Shard":"0", "Department":"IT", "Replica":"3"}
    }
	`)

	fmt.Println("new ..")

	jsonBlob = []byte(`{"YellowPages" : 
	                    "Shard": {"Department":"IT", "Replica": 3}
	                   }`)

	err = json.Unmarshal(jsonBlob, &dat)

	if err != nil {
		fmt.Println("error", err)
	}

	fmt.Println(dat)

}

type Message struct {
	ShardID    int
	SequenceID int
	PayLoad    []string
}

type View2 struct {
	ShardID int      `json:"shardID"`
	Record  []string `json:"record"`
}

func test2() {

	// create Payload

	// Testing some types
	// Boolean
	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))
	// Integer
	intB, _ := json.Marshal(100)
	fmt.Println(string(intB))

	// Float
	fltB, _ := json.Marshal(2.34)
	fmt.Println(string(fltB))

	// String
	strB, _ := json.Marshal("test")
	fmt.Println(string(strB))

	// List
	slcD := []string{"aaa", "bbb", "ccc"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB))

	// Map
	mapD := map[string]int{"hugo": 5, "berta": 7}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB))

	mes1 := new(Message)
	mes1.ShardID = 0
	mes1.SequenceID = 0
	mes1.PayLoad = []string{"Admin", "YelllowPages", "BluePages"}

	mes1B, _ := json.Marshal(mes1)
	fmt.Println(string(mes1B))

	res2D := &View2{
		ShardID: 0,
		Record:  []string{"0", "101", "102"}}
	res2B, _ := json.Marshal(res2D)
	fmt.Println(string(res2B))
	/*
		byt := []byte(`{"num":6.13,"strs":["a","b"]}`)

		var dat map[string]interface{}

		if err := json.Unmarshal(byt, &dat); err != nil {
		    panic(err)
		}
		fmt.Println(dat)

		num := dat["num"].(float64)
		fmt.Println(num)

		strs := dat["strs"].([]interface{})
		str1 := strs[0].(string)
		fmt.Println(str1)

		str := `{"page": 1, "fruits": ["apple", "peach"]}`
		res := response2{}
		json.Unmarshal([]byte(str), &res)
		fmt.Println(res)
		fmt.Println(res.Fruits[0])

		enc := json.NewEncoder(os.Stdout)
		d := map[string]int{"apple": 5, "lettuce": 7}
		enc.Encode(d)
	*/
}

/*
  "Shard" : 0, "Actors": ["Admin", "YellowPages", "BluePages" ]
*/
