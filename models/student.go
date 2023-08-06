package models

import (
	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	NoAbsen    int64 `gorm:"unique"`
	FullName   string
	Address    string `gorm:"type:text"`
	BirthDate  string
	Class      string
	Batch      int64
	SchoolName string
}
