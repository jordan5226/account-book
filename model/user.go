package model

import (
	"account-book/lib/pgdb/schema"
	"errors"

	"gorm.io/gorm"
)

type IModelUser interface {
	Get(uid string, pwd string) ([]schema.User, error)
	Add(data *schema.User) error
	Update(data *schema.User) error
	Delete(id string, uid string, pwd string) error
}

// Derived from Base Class: Model
type UserModel struct {
	Model
}

func (m *UserModel) InitModel(db *gorm.DB) {
	m.Model.InitModel(db)
}

func (m *UserModel) GetDB() *gorm.DB {
	return m.Model.GetDB()
}

func NewUserModel() *UserModel {
	m := new(UserModel)
	m.InitModel(db)
	return m
}

// Implement Interface
func (m *UserModel) Get(uid string, pwd string) (data []schema.User, err error) {
	if m.GetDB() == nil {
		return nil, errors.New("[ AccountBook ] model.UserModel::Get - Invalid DB")
	}

	err = m.GetDB().Where("uid = ? and pwd = ?", uid, pwd).Limit(1).Find(&data).Error
	return
}

func (m *UserModel) Add(data *schema.User) error {
	if m.GetDB() == nil {
		return errors.New("[ AccountBook ] model.UserModel::Add - Invalid DB")
	}

	err := m.GetDB().Create(data).Error
	return err
}

func (m *UserModel) Update(data *schema.User) error {
	if m.GetDB() == nil {
		return errors.New("[ AccountBook ] model.UserModel::Update - Invalid DB")
	}

	err := m.GetDB().Where("id = ?", data.Id).Updates(data).Error
	return err
}

func (m *UserModel) Delete(id string, uid string, pwd string) error {
	if m.GetDB() == nil {
		return errors.New("[ AccountBook ] model.UserModel::Delete - Invalid DB")
	}

	err := m.GetDB().Where("id = ? and uid = ? and pwd = ?", id, uid, pwd).Delete(&schema.User{}).Error
	return err
}
