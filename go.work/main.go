package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "MR.ibrahim2001"
	dbname   = "db_1"
)

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {


	// DB connection var
	psqlconn := fmt.Sprintf("host= %s port= %d user= %s password= %s dbname= %s sslmode=disable", host, port, user, password, dbname)

	// DB open
	db, e := sql.Open("postgres", psqlconn)
	CheckError(e)

	// DB close
	defer db.Close()

	// Insert into table (hardcoded)
	insert := `insert into "Students"("Name, Id) values('ibrahim', 20)`
	_, er := db.Exec(insert) // which executes (db.Exec) what was inserted above !
	CheckError(er)

	// dynamic insert
	insertdyn := `insert into "Students"(Name, Id) values($1, $2)`
	_, errr := db.Exec(insertdyn, "ibrahim", 20)
	CheckError(errr)

	// Update table 
	// update is to add a new field to the table using the "where" keyword
	updateStm := `update "Students" set "Name"=$1, "Id"=$2 where "roll_number"=$3`
	_, errrr := db.Exec(updateStm, "abunaib", 2, 3)

	CheckError(errrr)

	// Delete from table
	// Using the where keyword as well

	deleteStm := `delete from "Students" where "roll_number"=$3`
	_, eror := db.Exec(deleteStm, 3)
	CheckError(eror)

	// Query rows from table
	rows, errorr := db.Query(`SELECT "Name", "Id" FROM "Students"`)
	CheckError(errorr)

	defer rows.Close()
	for rows.Next() {
		var Name string
		var Id int 

		errorr = rows.Scan(&Name, &Id)
		CheckError(errorr)

		fmt.Println(Name, Id)

	}
	CheckError(errorr)
}
