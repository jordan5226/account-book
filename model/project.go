package model

import (
	"account-book/lib/pgdb/schema"
	"errors"

	"gorm.io/gorm"
)

type IModelProject interface {
	Get(uid string) ([]schema.Project, error)
	Add(data *schema.Project) error
	Update(data *schema.Project) error
	Delete(id string) error
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
func (m *ProjectModel) Get(uid string) (data []schema.Project, err error) {
	if m.GetDB() == nil {
		return nil, errors.New("[ AccountBook ] model.ProjectModel::Get - Invalid DB")
	}

	err = m.GetDB().Where("user_id = ?", uid).Find(&data).Error
	return
}

func (m *ProjectModel) Add(data *schema.Project) error {
	if m.GetDB() == nil {
		return errors.New("[ AccountBook ] model.ProjectModel::Add - Invalid DB")
	}

	err := m.GetDB().Create(data).Error
	return err
}

func (m *ProjectModel) Update(data *schema.Project) error {
	if m.GetDB() == nil {
		return errors.New("[ AccountBook ] model.ProjectModel::Update - Invalid DB")
	}

	err := m.GetDB().Where("id = ?", data.Id).Updates(data).Error
	return err
}

func (m *ProjectModel) Delete(id string) error {
	if m.GetDB() == nil {
		return errors.New("[ AccountBook ] model.ProjectModel::Delete - Invalid DB")
	}

	err := m.GetDB().Where("id = ?", id).Delete(&schema.Project{}).Error
	return err
}
