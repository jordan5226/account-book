package schema

type Project struct {
	Id     string `json:"id" gorm:"primaryKey;type:uuid;NOT NULL;default:uuid_generate_v4()"`
	UserId string `json:"userId" gorm:"foreignKey;type:uuid;NOT NULL" binding:"required"`
	Name   string `json:"name"`
	Icon   string `json:"icon"`
}

func (p *Project) GetTableName() string {
	return "projects"
}
