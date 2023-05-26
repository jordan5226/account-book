package schema

type Entry struct {
	Id       string    `json:"id" gorm:"primaryKey;type:uuid;NOT NULL;default:uuid_generate_v4()"`
	UserId   string    `json:"userId" binding:"required"`
	Time     LocalTime `json:"time" time_format:"2023-04-01"`
	Behavior int       `json:"behavior" binding:"required,min=0,max=2"`
	Amount   int       `json:"amount" binding:"required,min=0"`
	Type     string    `json:"type" binding:"required"`
	Account  string    `json:"account" binding:"required"`
	Project  string    `json:"project" binding:"required"`
	Store    string    `json:"store" binding:"required"`
	Note     string    `json:"note" binding:"chkNote"`
}

func (e *Entry) GetTableName() string {
	return "entries"
}
