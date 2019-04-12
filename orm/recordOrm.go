package orm

import (
	"github.com/jinzhu/gorm"
	"github.com/mahdifr17/phonebook-v2/structs"
)

// RecordsORM is a struct act as ORM for phonebook model
type RecordsORM struct {
	DB *gorm.DB
}

// GetAllRecords handle query
func (orm *RecordsORM) GetAllRecords() []structs.Record {
	var records []structs.Record

	orm.DB.Find(&records)

	return records
}
