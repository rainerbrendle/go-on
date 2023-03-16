package mixin

import (
	"fmt"
)

type DbConfig struct {
	dbInfo string
	tested bool
	active bool
}

// var conf DbConfig( active: false )
var conf DbConfig

func TestDB() {

	fmt.Printf("  database test %v\n", conf.active)
	if conf.active == false {
		return
	}

}

func ActivateDB() {

	conf.active = false
	conf.dbInfo = DbInfo()

	fmt.Printf("  database info %v\n", conf.dbInfo)

}
