package structs

import (
	"github.com/jinzhu/gorm"
)

// Record is a struct that represent phonebook data
type Record struct {
	gorm.Model `json:"-"` // Ignore ID, CreatedAt, UpdatedAt, DeletedAt
	Name       string     `json:"name,omitempty" gorm:"not null"`
	Phone      string     `json:"phone,omitempty" gorm:"not null;unique"`
}
