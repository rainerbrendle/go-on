package mixin

import (
	"fmt"
)

// database connection constants
// (isolated from the rest of the test)

const host = "localhost"
const port = 5432
const user = "test"
const password = "test"
const dbname = "test"

func DbInfo() string {
	psqlInfo := fmt.Sprintf(
		"host=%s port=%d user=%s "+
			"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	return psqlInfo
}
