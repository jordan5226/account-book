package model

import (
	"errors"
	"time"

	"account-book/lib/pgdb/schema"

	"gorm.io/gorm"
)

type IModelEntry interface {
	Get(_time time.Time, uid string) ([]schema.Entry, error)
	Add(data *schema.Entry) error
	Update(data *schema.Entry) error
	Delete(uid string, id string) error
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
func (m *EntryModel) Get(_time time.Time, uid string) (data []schema.Entry, err error) {
	if m.GetDB() == nil {
		return nil, errors.New("[ AccountBook ] model.EntryModel::Get - Invalid DB")
	}

	err = m.GetDB().Where("user_id = ? and ? <= entry_time and entry_time <= ?",
		uid,
		time.Date(_time.Year(), _time.Month(), _time.Day(), 0, 0, 0, 0, _time.Location()),
		time.Date(_time.Year(), _time.Month(), _time.Day(), 23, 59, 59, 999000000*int(time.Nanosecond), _time.Location())).Find(&data).Error
	return
}

func (m *EntryModel) Add(data *schema.Entry) error {
	if m.GetDB() == nil {
		return errors.New("[ AccountBook ] model.EntryModel::Add - Invalid DB")
	}

	err := m.GetDB().Create(data).Error
	return err
}

func (m *EntryModel) Update(data *schema.Entry) error {
	if m.GetDB() == nil {
		return errors.New("[ AccountBook ] model.EntryModel::Update - Invalid DB")
	}

	err := m.GetDB().Where("id = ?", data.Id).Updates(data).Error
	return err
}

func (m *EntryModel) Delete(uid string, id string) error {
	if m.GetDB() == nil {
		return errors.New("[ AccountBook ] model.EntryModel::Delete - Invalid DB")
	}

	err := m.GetDB().Where("user_id = ? and id = ?", uid, id).Delete(&schema.Entry{}).Error
	return err
}
