package model

import (
	"account-book/lib/pgdb/schema"
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type IModelAccount interface {
	Get(uid string, tx *gorm.DB) ([]schema.Account, error)
	Add(data *schema.Account, tx *gorm.DB) (*schema.User, error)
	Update(data *schema.Account, tx *gorm.DB) error
	Delete(id string, tx *gorm.DB) error
}

// Derived from Base Class: Model
type AccountModel struct {
	Model
}

func (m *AccountModel) InitModel(db *gorm.DB) {
	m.Model.InitModel(db)
}

func (m *AccountModel) GetDB() *gorm.DB {
	return m.Model.GetDB()
}

func NewAccountModel() *AccountModel {
	m := new(AccountModel)
	m.InitModel(db)
	return m
}

// Implement Interface
func (m *AccountModel) Get(uid string, tx *gorm.DB) (data []schema.Account, err error) {
	if m.GetDB() == nil {
		return nil, errors.New("[ AccountBook ] model.AccountModel::Get - Invalid DB")
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

func (m *AccountModel) Add(data *schema.Account, tx *gorm.DB) (*schema.Account, error) {
	if m.GetDB() == nil {
		return nil, errors.New("[ AccountBook ] model.AccountModel::Add - Invalid DB")
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

func (m *AccountModel) Update(data *schema.Account, tx *gorm.DB) error {
	if m.GetDB() == nil {
		return errors.New("[ AccountBook ] model.AccountModel::Update - Invalid DB")
	}

	if tx == nil {
		return m.GetDB().Transaction(func(tx *gorm.DB) error {
			return tx.Where("id = ?", data.Id).Updates(data).Error
		})
	}

	return tx.Where("id = ?", data.Id).Updates(data).Error
}

func (m *AccountModel) Delete(id string, tx *gorm.DB) error {
	if m.GetDB() == nil {
		return errors.New("[ AccountBook ] model.AccountModel::Delete - Invalid DB")
	}

	if tx == nil {
		return m.GetDB().Transaction(func(tx *gorm.DB) error {
			return tx.Where("id = ?", id).Delete(&schema.Account{}).Error
		})
	}

	return tx.Where("id = ?", id).Delete(&schema.Account{}).Error
}
