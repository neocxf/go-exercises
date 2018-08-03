package models

import (
	"encoding/xml"
)

type User struct {
	XMLName xml.Name `json:"-"  xml:"person"`
	Name    string   `json:"name" xml:"name"`
	Id      int64    `json:"id" xml:"id"`
}
