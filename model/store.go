package model

import (
	"account-book/lib/pgdb/schema"
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type IModelStore interface {
	Get(uid string, tx *gorm.DB) ([]schema.Store, error)
	Add(data *schema.Store, tx *gorm.DB) (*schema.Store, error)
	Update(data *schema.Store, tx *gorm.DB) error
	Delete(id string, tx *gorm.DB) error
}

// Derived from Base Class: Model
type StoreModel struct {
	Model
}

func (m *StoreModel) InitModel(db *gorm.DB) {
	m.Model.InitModel(db)
}

func (m *StoreModel) GetDB() *gorm.DB {
	return m.Model.GetDB()
}

func NewStoreModel() *StoreModel {
	m := new(StoreModel)
	m.InitModel(db)
	return m
}

// Implement Interface
func (m *StoreModel) Get(uid string, tx *gorm.DB) (data []schema.Store, err error) {
	if m.GetDB() == nil {
		return nil, errors.New("[ AccountBook ] model.StoreModel::Get - Invalid DB")
	}

	if tx == nil {
		err = m.GetDB().Transaction(func(tx *gorm.DB) error {
			return tx.Clauses(clause.Locking{Strength: "NO KEY UPDATE"}).Where("user_id = ?", uid).Find(&data).Error
		})
		return
	}

	err = tx.Clauses(clause.Locking{Strength: "NO KEY UPDATE"}).Where("user_id = ?", uid).Find(&data).Error
	return
}

func (m *StoreModel) Add(data *schema.Store, tx *gorm.DB) (*schema.Store, error) {
	if m.GetDB() == nil {
		return nil, errors.New("[ AccountBook ] model.StoreModel::Add - Invalid DB")
	}

	//
	if len(data.Id) == 0 {
		if id, err := uuid.NewRandom(); err != nil {
			return nil, err
		} else {
			data.Id = id.String()
		}
	}

	//
	var err error

	if tx == nil {
		err = m.GetDB().Transaction(func(tx *gorm.DB) error {
			return tx.Create(data).Error
		})
		return data, err
	}

	err = tx.Create(data).Error
	return data, err
}

func (m *StoreModel) Update(data *schema.Store, tx *gorm.DB) error {
	if m.GetDB() == nil {
		return errors.New("[ AccountBook ] model.StoreModel::Update - Invalid DB")
	}

	if tx == nil {
		return m.GetDB().Transaction(func(tx *gorm.DB) error {
			return tx.Where("id = ?", data.Id).Updates(data).Error
		})
	}

	return tx.Where("id = ?", data.Id).Updates(data).Error
}

func (m *StoreModel) Delete(id string, tx *gorm.DB) error {
	if m.GetDB() == nil {
		return errors.New("[ AccountBook ] model.StoreModel::Delete - Invalid DB")
	}

	if tx == nil {
		return m.GetDB().Transaction(func(tx *gorm.DB) error {
			return tx.Where("id = ?", id).Delete(&schema.Store{}).Error
		})
	}

	return tx.Where("id = ?", id).Delete(&schema.Store{}).Error
}
