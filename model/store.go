package model

import (
	"account-book/lib/pgdb/schema"
	"errors"

	"gorm.io/gorm"
)

type IModelStore interface {
	Get(id string) ([]schema.Store, error)
	Add(data *schema.Store) error
	Update(data *schema.Store) error
	Delete(id string) error
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
func (m *StoreModel) Get(uid string) (data []schema.Store, err error) {
	if m.GetDB() == nil {
		return nil, errors.New("[ AccountBook ] model.StoreModel::Get - Invalid DB")
	}

	err = m.GetDB().Where("user_id = ?", uid).Find(&data).Error
	return
}

func (m *StoreModel) Add(data *schema.Store) error {
	if m.GetDB() == nil {
		return errors.New("[ AccountBook ] model.StoreModel::Add - Invalid DB")
	}

	err := m.GetDB().Create(data).Error
	return err
}

func (m *StoreModel) Update(data *schema.Store) error {
	if m.GetDB() == nil {
		return errors.New("[ AccountBook ] model.StoreModel::Update - Invalid DB")
	}

	err := m.GetDB().Where("id = ?", data.Id).Updates(data).Error
	return err
}

func (m *StoreModel) Delete(id string) error {
	if m.GetDB() == nil {
		return errors.New("[ AccountBook ] model.StoreModel::Delete - Invalid DB")
	}

	err := m.GetDB().Where("id = ?", id).Delete(&schema.Store{}).Error
	return err
}
