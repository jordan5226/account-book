package model

import (
	"account-book/lib/pgdb/schema"
	"errors"

	"gorm.io/gorm"
)

type IModelType interface {
	Get() ([]schema.Type, error)
	GetByID(uid string) ([]schema.Account, error)
	Add(data *schema.Type) (*schema.Type, error)
	Update(data *schema.Type) error
	Delete(id string) error
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
func (m *TypeModel) Get() (data []schema.Type, err error) {
	if m.GetDB() == nil {
		return nil, errors.New("[ AccountBook ] model.TypeModel::Get - Invalid DB")
	}

	err = m.GetDB().Find(&data).Error
	return
}

func (m *TypeModel) GetByID(id string) (data []schema.Type, err error) {
	if m.GetDB() == nil {
		return nil, errors.New("[ AccountBook ] model.TypeModel::GetByID - Invalid DB")
	}

	err = m.GetDB().Where("id = ?", id).Find(&data).Error
	return
}

func (m *TypeModel) Add(data *schema.Type) (*schema.Type, error) {
	if m.GetDB() == nil {
		return nil, errors.New("[ AccountBook ] model.TypeModel::Add - Invalid DB")
	}

	err := m.GetDB().Create(data).Error
	return data, err
}

func (m *TypeModel) Update(data *schema.Type) error {
	if m.GetDB() == nil {
		return errors.New("[ AccountBook ] model.TypeModel::Update - Invalid DB")
	}

	err := m.GetDB().Where("id = ?", data.Id).Updates(data).Error
	return err
}

func (m *TypeModel) Delete(id string) error {
	if m.GetDB() == nil {
		return errors.New("[ AccountBook ] model.TypeModel::Delete - Invalid DB")
	}

	err := m.GetDB().Where("id = ?", id).Delete(&schema.Type{}).Error
	return err
}
