package model

import (
	"errors"
	"time"

	"account-book/lib/pgdb/schema"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type IModelEntry interface {
	Get(_time time.Time, uid string, tx *gorm.DB) ([]schema.Entry, error)
	GetByID(id string, tx *gorm.DB) ([]schema.Entry, error)
	Add(data *schema.Entry, tx *gorm.DB) (*schema.Entry, error)
	Update(data *schema.Entry, tx *gorm.DB) error
	Delete(uid string, id string, tx *gorm.DB) error
}

// Derived from Base Class: Model
type EntryModel struct {
	Model
}

func (m *EntryModel) InitModel(db *gorm.DB) {
	m.Model.InitModel(db)
}

func (m *EntryModel) GetDB() *gorm.DB {
	return m.Model.GetDB()
}

func NewEntryModel() *EntryModel {
	m := new(EntryModel)
	m.InitModel(db)
	return m
}

// Implement Interface
func (m *EntryModel) Get(_time time.Time, uid string, tx *gorm.DB) (data []schema.Entry, err error) {
	if m.GetDB() == nil {
		return nil, errors.New("[ AccountBook ] model.EntryModel::Get - Invalid DB")
	}

	if tx == nil {
		err = m.GetDB().Transaction(func(tx *gorm.DB) error {
			return tx.Clauses(clause.Locking{Strength: "NO KEY UPDATE"}).Where("user_id = ? and ? <= time and time <= ?",
				uid,
				time.Date(_time.Year(), _time.Month(), _time.Day(), 0, 0, 0, 0, _time.Location()),
				time.Date(_time.Year(), _time.Month(), _time.Day(), 23, 59, 59, 999000000*int(time.Nanosecond), _time.Location())).Find(&data).Error
		})
		return
	}

	err = tx.Clauses(clause.Locking{Strength: "NO KEY UPDATE"}).Where("user_id = ? and ? <= time and time <= ?",
		uid,
		time.Date(_time.Year(), _time.Month(), _time.Day(), 0, 0, 0, 0, _time.Location()),
		time.Date(_time.Year(), _time.Month(), _time.Day(), 23, 59, 59, 999000000*int(time.Nanosecond), _time.Location())).Find(&data).Error
	return
}

func (m *EntryModel) GetByID(id string, tx *gorm.DB) (data []schema.Entry, err error) {
	if m.GetDB() == nil {
		return nil, errors.New("[ AccountBook ] model.EntryModel::GetByID - Invalid DB")
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

func (m *EntryModel) Add(data *schema.Entry, tx *gorm.DB) (*schema.Entry, error) {
	if m.GetDB() == nil {
		return nil, errors.New("[ AccountBook ] model.EntryModel::Add - Invalid DB")
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

func (m *EntryModel) Update(data *schema.Entry, tx *gorm.DB) error {
	if m.GetDB() == nil {
		return errors.New("[ AccountBook ] model.EntryModel::Update - Invalid DB")
	}

	if tx == nil {
		return m.GetDB().Transaction(func(tx *gorm.DB) error {
			return tx.Where("id = ?", data.Id).Updates(data).Error
		})
	}

	return tx.Where("id = ?", data.Id).Updates(data).Error
}

func (m *EntryModel) Delete(uid string, id string, tx *gorm.DB) error {
	if m.GetDB() == nil {
		return errors.New("[ AccountBook ] model.EntryModel::Delete - Invalid DB")
	}

	if tx == nil {
		return m.GetDB().Transaction(func(tx *gorm.DB) error {
			return tx.Where("user_id = ? and id = ?", uid, id).Delete(&schema.Entry{}).Error
		})
	}

	return tx.Where("user_id = ? and id = ?", uid, id).Delete(&schema.Entry{}).Error
}
