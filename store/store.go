package store

type User struct {
	ID       int    `json:"id"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type UserStore interface {
	CreateUser(user *User) error
	GetUserByEmail(email string) (*User, error)
	StoreRes(ExternalUser *ExternalUser) error
	StoreCatFact
}
type StoreCatFact interface {
	StoreCatFact(fact *CatFact) error
}
type ExternalUser struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type CatFact struct {
	Fact   string `json:"fact"`
	Length int    `json:"length"`
}
