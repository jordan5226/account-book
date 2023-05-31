package model

import (
	"account-book/lib/pgdb/schema"
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type IModelProject interface {
	Get(uid string, tx *gorm.DB) ([]schema.Project, error)
	Add(data *schema.Project, tx *gorm.DB) (*schema.Project, error)
	Update(data *schema.Project, tx *gorm.DB) error
	Delete(id string, tx *gorm.DB) error
}

// Derived from Base Class: Model
type ProjectModel struct {
	Model
}

func (m *ProjectModel) InitModel(db *gorm.DB) {
	m.Model.InitModel(db)
}

func (m *ProjectModel) GetDB() *gorm.DB {
	return m.Model.GetDB()
}

func NewProjectModel() *ProjectModel {
	m := new(ProjectModel)
	m.InitModel(db)
	return m
}

// Implement Interface
func (m *ProjectModel) Get(uid string, tx *gorm.DB) (data []schema.Project, err error) {
	if m.GetDB() == nil {
		return nil, errors.New("[ AccountBook ] model.ProjectModel::Get - Invalid DB")
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

func (m *ProjectModel) Add(data *schema.Project, tx *gorm.DB) (*schema.Project, error) {
	if m.GetDB() == nil {
		return nil, errors.New("[ AccountBook ] model.ProjectModel::Add - Invalid DB")
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

func (m *ProjectModel) Update(data *schema.Project, tx *gorm.DB) error {
	if m.GetDB() == nil {
		return errors.New("[ AccountBook ] model.ProjectModel::Update - Invalid DB")
	}

	if tx == nil {
		return m.GetDB().Transaction(func(tx *gorm.DB) error {
			return tx.Where("id = ?", data.Id).Updates(data).Error
		})
	}

	return tx.Where("id = ?", data.Id).Updates(data).Error
}

func (m *ProjectModel) Delete(id string, tx *gorm.DB) error {
	if m.GetDB() == nil {
		return errors.New("[ AccountBook ] model.ProjectModel::Delete - Invalid DB")
	}

	if tx == nil {
		return m.GetDB().Transaction(func(tx *gorm.DB) error {
			return tx.Where("id = ?", id).Delete(&schema.Project{}).Error
		})
	}

	return tx.Where("id = ?", id).Delete(&schema.Project{}).Error
}
