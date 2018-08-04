package repository

import (
	"fmt"
	"log"
	"testing"

	"github.com/neocxf/go-exercises/bookcase/models"
)

var (
	instance      SQL
	closeDbHandle func()
)

func init() {

	instance = SQL{}

	closeDbHandle = instance.Open("sqlite3", `sqlite.db`)

	instance.InitSchema()

}

func TestCreateUser(t *testing.T) {
	// t.SkipNow()

	defer func() {
		fmt.Println("althrough the test case failed, this line should still appear...")

		if r := recover(); r != nil {
			fmt.Println("recovered in f", r)
		}

		//defer instance.Close()

	}()

	log.SetPrefix("bookcase ==> ")

	u := &models.User{Name: "fei"}

	err := instance.CreateUser(u)

	if err != nil {
		log.Fatal(err)
	}

}

func TestCreateUsers(t *testing.T) {
	//defer instance.Close()

	log.SetPrefix("bookcase ==> ")

	users := make([]*models.User, 5)

	for i := range users {
		name := fmt.Sprintf("%s%d", "fei", i)

		users = append(users, &models.User{Name: name})
	}

	err := instance.CreateUsers(users)

	if err != nil {
		log.Fatal(err)
		return
	}

	log.Printf("successfully finish the job")
}

func TestSelectUser(t *testing.T) {
	//defer closeDbHandle()

	rows, _ := instance.DB.Query(`select name, id from user`)

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
}

func TestThrowErrors(t *testing.T) {
	t.SkipNow()

	defer func() {

		if r := recover(); r != nil {
			fmt.Println("recovered in f", r)
		}

	}()

	panic("deliberately throw a panic for fun")

	t.Error(`there is an error`)

	//logrus.Fatal(errors.New("exit directly"))

}

func BenchmarkHello(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Println(fmt.Sprintf("hello"))
	}
}
