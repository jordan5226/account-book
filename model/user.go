package model

import (
	"account-book/lib/pgdb/schema"
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type IModelUser interface {
	GetByID(id string, tx *gorm.DB) ([]schema.User, error)
	Get(uid string, pwd string, tx *gorm.DB) ([]schema.User, error)
	Add(data *schema.User, tx *gorm.DB) (*schema.User, error)
	Update(data *schema.User, tx *gorm.DB) error
	Delete(id string, uid string, pwd string, tx *gorm.DB) error
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
func (m *UserModel) GetByID(id string, tx *gorm.DB) (data []schema.User, err error) {
	if m.GetDB() == nil {
		return nil, errors.New("[ AccountBook ] model.UserModel::GetByID - Invalid DB")
	}

	if tx == nil {
		err = m.GetDB().Transaction(func(tx *gorm.DB) error {
			return tx.Clauses(clause.Locking{Strength: "NO KEY UPDATE"}).Where("id = ?", id).Limit(1).Find(&data).Error
		})
		return
	}

	err = tx.Clauses(clause.Locking{Strength: "NO KEY UPDATE"}).Where("id = ?", id).Limit(1).Find(&data).Error
	return
}

func (m *UserModel) Get(uid string, pwd string, tx *gorm.DB) (data []schema.User, err error) {
	if m.GetDB() == nil {
		return nil, errors.New("[ AccountBook ] model.UserModel::Get - Invalid DB")
	}

	if tx == nil {
		err = m.GetDB().Transaction(func(tx *gorm.DB) error {
			return tx.Clauses(clause.Locking{Strength: "NO KEY UPDATE"}).Where("uid = ? and pwd = ?", uid, pwd).Limit(1).Find(&data).Error
		})
		return
	}

	err = tx.Clauses(clause.Locking{Strength: "NO KEY UPDATE"}).Where("uid = ? and pwd = ?", uid, pwd).Limit(1).Find(&data).Error
	return
}

func (m *UserModel) Add(data *schema.User, tx *gorm.DB) (*schema.User, error) {
	if m.GetDB() == nil {
		return nil, errors.New("[ AccountBook ] model.UserModel::Add - Invalid DB")
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

func (m *UserModel) Update(data *schema.User, tx *gorm.DB) error {
	if m.GetDB() == nil {
		return errors.New("[ AccountBook ] model.UserModel::Update - Invalid DB")
	}

	if tx == nil {
		return m.GetDB().Transaction(func(tx *gorm.DB) error {
			return tx.Where("id = ?", data.Id).Updates(data).Error
		})
	}

	return tx.Where("id = ?", data.Id).Updates(data).Error
}

func (m *UserModel) Delete(id string, uid string, pwd string, tx *gorm.DB) error {
	if m.GetDB() == nil {
		return errors.New("[ AccountBook ] model.UserModel::Delete - Invalid DB")
	}

	if tx == nil {
		return m.GetDB().Transaction(func(tx *gorm.DB) error {
			return tx.Where("id = ? and uid = ? and pwd = ?", id, uid, pwd).Delete(&schema.User{}).Error
		})
	}

	return tx.Where("id = ? and uid = ? and pwd = ?", id, uid, pwd).Delete(&schema.User{}).Error
}
