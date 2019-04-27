package orm

import (
	"github.com/jinzhu/gorm"
	"github.com/mahdifr17/phonebook-v2/exception"
	"github.com/mahdifr17/phonebook-v2/structs"
)

// RecordsORM is a struct act as ORM for phonebook model
type RecordsORM struct {
	DB *gorm.DB
}

// GetAllRecords handle query to get all records
func (orm *RecordsORM) GetAllRecords() []structs.Record {
	var records []structs.Record

	orm.DB.Find(&records)

	return records
}

// GetRecord handle query to get specific record
func (orm *RecordsORM) GetRecord(id int) (structs.Record, error) {
	var record structs.Record

	if orm.DB.First(&record, "id = ?", id).RecordNotFound() {
		// Record not found
		return structs.Record{}, exception.NotFound(id)
	}
	return record, nil
}

// AddRecord handle query to insert record
func (orm *RecordsORM) AddRecord(newRecord structs.Record) (structs.Record, error) {
	if err := orm.DB.Create(&newRecord).Error; err != nil {
		// Error insert
		return structs.Record{}, err
	}
	return newRecord, nil
}

// UpdateRecord handle query to update changed field record
func (orm *RecordsORM) UpdateRecord(oldRecord structs.Record, newRecord structs.Record) (structs.Record, error) {
	if err := orm.DB.Model(&oldRecord).Updates(newRecord).Error; err != nil {
		// Error update
		return structs.Record{}, err
	}
	return oldRecord, nil
}

func (orm *RecordsORM) DeleteRecord(record structs.Record) (structs.Record, error) {
	if err := orm.DB.Delete(&record).Error; err != nil {
		// Error update
		return structs.Record{}, err
	}
	return record, nil
}
