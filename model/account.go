package model

import (
	"account-book/lib/pgdb/schema"
	"errors"

	"gorm.io/gorm"
)

type IModelAccount interface {
	Get(id string) ([]schema.Account, error)
	Add(data *schema.Account) error
	Update(data *schema.Account) error
	Delete(id string) error
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
func (m *AccountModel) Get(uid string) (data []schema.Account, err error) {
	if m.GetDB() == nil {
		return nil, errors.New("[ AccountBook ] model.AccountModel::Get - Invalid DB")
	}

	err = m.GetDB().Where("user_id = ?", uid).Find(&data).Error
	return
}

func (m *AccountModel) Add(data *schema.Account) error {
	if m.GetDB() == nil {
		return errors.New("[ AccountBook ] model.AccountModel::Add - Invalid DB")
	}

	err := m.GetDB().Create(data).Error
	return err
}

func (m *AccountModel) Update(data *schema.Account) error {
	if m.GetDB() == nil {
		return errors.New("[ AccountBook ] model.AccountModel::Update - Invalid DB")
	}

	err := m.GetDB().Where("id = ?", data.Id).Updates(data).Error
	return err
}

func (m *AccountModel) Delete(id string) error {
	if m.GetDB() == nil {
		return errors.New("[ AccountBook ] model.AccountModel::Delete - Invalid DB")
	}

	err := m.GetDB().Where("id = ?", id).Delete(&schema.Account{}).Error
	return err
}
