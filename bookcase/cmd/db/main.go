package main

import (
	"database/sql"
	"fmt"
	"github.com/neocxf/go-exercises/bookcase/models"
	_ "github.com/go-sql-driver/mysql"

)

func main() {

	db, err := sql.Open("mysql", `root:derbysoft@/dplatform_rtd?charset=utf8`)

	if err != nil {
		fmt.Printf(err.Error())
	}

	rows, err := db.Query(`select name, id from user`)

	result := make([]*models.User, 0)

	for rows.Next() {
		t := new(models.User)
		err := rows.Scan(&t.Name, &t.Id)
		if err != nil {
			return
		}

		result = append(result, t)
	}

	for _, elem := range result {
		fmt.Println(elem)
	}



	fmt.Println("it works")
}
