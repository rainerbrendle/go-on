package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func main() {

	fmt.Println("Hello")

	test()
	test01()

}

func test() {
	// test 1
	// create basid marshaling

	fmt.Println("Marshal some basid data types")

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
	strB, _ := json.Marshal("test")
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

	jsonBlob1 = []byte(`[    
     {  
        "ID": 1, 
        "Shard": "0",
        "Version": "1.0.0.1",
        "ValidFrom": "2023-03-01",
        "Organization" : "IT",
        "Database" : "postgres",
        "Replica": 3
     },
     {
        "ID": 2,
        "Shard": "100",
        "Version": "1.0.0.1",
        "ValidFrom": "2023-03-01",
        "Organization" : "Sales",
        "Database" : "postgres",
        "Replica": 3
     }
]  `)

	type Admin struct {
		ID           int
		Shard        int
		Version      string
		ValidFrom    string
		Organization string
		Database     string
		Replica      int
	}

	dat := []Admin{
		{
			1,
			0,
			"1.0.0.1",
			"2023-03-01",
			"IT",
			"postgres",
			3},
		{
			2,
			100,
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

	fmt.Printf("input %v\n", dat)

	s := string(jsonBlob2)

	fmt.Printf("output %v\n", s)

	// remove white space from JSON
	// we need a bytes.Buffer as intermediate buffer
	buffer := new(bytes.Buffer)
	if err := json.Compact(buffer, jsonBlob1); err != nil {
		fmt.Println(err)
	}
	fmt.Printf("check %v\n", buffer)

	// map to a generic slice (array)
	// copy  line by line
	var anything []interface{}
	for _, val := range dat {
		anything = append(anything, val)
	}

	fmt.Printf("anything %v\n", anything)

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
