package status

import (
	"saver/model"

	"gorm.io/gorm"
)

type Status interface {
	Insert(status model.Status) error
}

type SQLStatus struct {
	DB *gorm.DB
}

func NewSQLStatus(d *gorm.DB) SQLStatus {
	return SQLStatus{DB: d}
}

func (m SQLStatus) Insert(status model.Status) error {
	result := m.DB.Create(&status)

	return result.Error
}
