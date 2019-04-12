package structs

import (
	"github.com/jinzhu/gorm"
)

// Record is a struct that represent phonebook data
type Record struct {
	gorm.Model
	Name  string `json:"Name,omitempty"`
	Phone string `json:"Phone,omitempty"`
}
