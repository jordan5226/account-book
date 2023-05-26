package schema

type User struct {
	Id       string    `json:"id" gorm:"primaryKey;type:uuid;NOT NULL;default:uuid_generate_v4()"`
	Name     string    `json:"name"`
	Uid      string    `json:"uid"`
	Pwd      string    `json:"pwd"`
	Currency string    `json:"currency"`
	CreateAt LocalTime `json:"createAt"`
}

func (u *User) GetTableName() string {
	return "users"
}
