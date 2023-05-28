package model

import (
	"gorm.io/gorm"
)

var db *gorm.DB

// Model Base Class
type IModel interface {
	InitModel(db *gorm.DB)
}

type Model struct {
	db *gorm.DB
}

func (m *Model) InitModel(db *gorm.DB) {
	m.db = db
}

func (m *Model) GetDB() *gorm.DB {
	return m.db
}

// Init DB Pointer
func Init(lpDB *gorm.DB) {
	db = lpDB
}

func GetDB() *gorm.DB {
	return db
}
