package impl

import (
	"fmt"
	. "github.com/neocxf/go-exercises/test/iface"
)

type Repository struct {
}

func (r *Repository) Fetch(args Args) (Data, error) {
	if len(args) == 0 {
		return Data{}, fmt.Errorf("No arguments provided")
	}

	data := Data{
		"user": "root",
		"pass": "pass",
	}

	fmt.Printf("Repository fetch data successfully: %v\n", data)

	return data, nil
}
