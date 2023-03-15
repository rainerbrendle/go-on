package mixin

import (
	"fmt"
)


type DbConfig struct {
   active bool }


// var conf DbConfig( active: false )
var conf DbConfig

func TestDB() {
  
   fmt.Printf( "  database test %v\n", conf.active )
   if conf.active == false {
          return
   }

}
