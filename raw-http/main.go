package main

import (
	"database/sql"
	"encoding/xml"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"github.com/neocxf/go-exercises/bookcase/models"
)

type HelloHandler struct {
	db *sql.DB
}

type NameHolderArray struct {
	XMLName xml.Name      `json:"-" xml:"persons"`
	Person  []*NameHolder `json:"persons" xml:"person"`
}

type NameHolder struct {
	XMLName xml.Name `json:"-"  xml:"person"`
	Name    string   `json:"name" xml:"name"`
	Id      int64    `json:"id" xml:"id"`
}

func (h *HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	result := make([]*NameHolder, 0)

	rows, err := h.db.Query(`select name, id from user`)

	CheckErr(err)

	for rows.Next() {
		t := new(NameHolder)
		err := rows.Scan(&t.Name, &t.Id)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		result = append(result, t)
	}

	for _, elem := range result {
		fmt.Println(elem)
	}

	w.Header().Set("Content-Type", "application/xml; charset=utf-8")

	xml.NewEncoder(w).Encode(&NameHolderArray{Person: result})

}

func main() {
	db, err := sql.Open("mysql", `root:derbysoft@/dplatform_rtd?charset=utf8`)

	rows, _ := db.Query(`select name, id from user`)

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

	CheckErr(err)

	//http.Handle("/hello", &HelloHandler{db: db})
	//
	//http.ListenAndServe(":3000", nil)

}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
