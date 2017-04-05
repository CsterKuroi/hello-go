package main

import(
	"fmt"
	"log"

	r "gopkg.in/gorethink/gorethink.v3"
)

func main() {
	session, err := r.Connect(r.ConnectOpts{
	    Address: "localhost:28015",
	})
	if err != nil {
	    log.Fatalln(err)
	}

	res, err := r.Expr("Hello World").Run(session)
	if err != nil {
	    log.Fatalln(err)
	}

	var response string
	err = res.One(&response)
	if err != nil {
	    log.Fatalln(err)
	}

	fmt.Println(response)

	resp, err := r.DBCreate("examples").RunWrite(session)
	if err != nil {
	    fmt.Print(err)
	}

	fmt.Println("%d DB created", resp.DBsCreated)

	// Setup database
	r.DB("examples").TableDrop("table").Run(session)

	respo, err := r.DB("examples").TableCreate("table").RunWrite(session)
	if err != nil {
	    log.Fatalf("Error creating table: %s", err)
	}

	fmt.Println("%d table created", respo.TablesCreated)
}
