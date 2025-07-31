package model

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Age  string `json:"age"`
}

func (User) TableName() string {
	return "t_user"
}
