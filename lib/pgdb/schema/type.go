package schema

type Type struct {
	Id   string `json:"id" gorm:"primaryKey;type:uuid;NOT NULL;default:uuid_generate_v4()"`
	Name string `json:"name"`
	Icon string `json:"icon"`
}

func (t *Type) GetTableName() string {
	return "entry_types"
}
