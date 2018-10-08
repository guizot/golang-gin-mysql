package user

type User struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Address   string `json:"address"`
	Age       int    `json:"age"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

//Users
type Users []User
