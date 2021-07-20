package Models

type User struct {
	Id       uint   `json:"id"`
	Name     string `json:"name"`
	Age      string `json:"age"`
	Position string `json:"position"`
	Wage     string `json:"wage"`
}

func (b *User) TableName() string {
	return "user"
}
