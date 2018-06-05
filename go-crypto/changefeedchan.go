package main

import (
	"fmt"

	r "unicontract/src/core/db/rethinkdb"
)

func main() {
	var value interface{}
	res :=r.Changefeed("bigchain","backlog")
	for res.Next(&value) {
		fmt.Println(value)
	}
}
