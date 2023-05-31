package model

import (
	"account-book/lib/pgdb/schema"
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type IModelType interface {
	Get(tx *gorm.DB) ([]schema.Type, error)
	GetByID(uid string, tx *gorm.DB) ([]schema.Account, error)
	Add(data *schema.Type, tx *gorm.DB) (*schema.Type, error)
	Update(data *schema.Type, tx *gorm.DB) error
	Delete(id string, tx *gorm.DB) error
}

// Derived from Base Class: Model
type TypeModel struct {
	Model
}

func (m *TypeModel) InitModel(db *gorm.DB) {
	m.Model.InitModel(db)
}

func (m *TypeModel) GetDB() *gorm.DB {
	return m.Model.GetDB()
}

func NewTypeModel() *TypeModel {
	m := new(TypeModel)
	m.InitModel(db)
	return m
}

// Implement Interface
func (m *TypeModel) Get(tx *gorm.DB) (data []schema.Type, err error) {
	if m.GetDB() == nil {
		return nil, errors.New("[ AccountBook ] model.TypeModel::Get - Invalid DB")
	}

	if tx == nil {
		err = m.GetDB().Transaction(func(tx *gorm.DB) error {
			return tx.Find(&data).Error
		})
		return
	}

	err = tx.Find(&data).Error
	return
}

func (m *TypeModel) GetByID(id string, tx *gorm.DB) (data []schema.Type, err error) {
	if m.GetDB() == nil {
		return nil, errors.New("[ AccountBook ] model.TypeModel::GetByID - Invalid DB")
	}

	if tx == nil {
		err = m.GetDB().Transaction(func(tx *gorm.DB) error {
			return tx.Clauses(clause.Locking{Strength: "NO KEY UPDATE"}).Where("id = ?", id).Find(&data).Error
		})
		return
	}

	err = tx.Clauses(clause.Locking{Strength: "NO KEY UPDATE"}).Where("id = ?", id).Find(&data).Error
	return
}

func (m *TypeModel) Add(data *schema.Type, tx *gorm.DB) (*schema.Type, error) {
	if m.GetDB() == nil {
		return nil, errors.New("[ AccountBook ] model.TypeModel::Add - Invalid DB")
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

func (m *TypeModel) Update(data *schema.Type, tx *gorm.DB) error {
	if m.GetDB() == nil {
		return errors.New("[ AccountBook ] model.TypeModel::Update - Invalid DB")
	}

	if tx == nil {
		return m.GetDB().Transaction(func(tx *gorm.DB) error {
			return tx.Where("id = ?", data.Id).Updates(data).Error
		})
	}

	return tx.Where("id = ?", data.Id).Updates(data).Error
}

func (m *TypeModel) Delete(id string, tx *gorm.DB) error {
	if m.GetDB() == nil {
		return errors.New("[ AccountBook ] model.TypeModel::Delete - Invalid DB")
	}

	if tx == nil {
		return m.GetDB().Transaction(func(tx *gorm.DB) error {
			return tx.Where("id = ?", id).Delete(&schema.Type{}).Error
		})
	}

	return tx.Where("id = ?", id).Delete(&schema.Type{}).Error
}
