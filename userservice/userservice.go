package userservice

type ThirdPartyUser struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type JsonPlaceholder interface {
	GetUsers() ([]ThirdPartyUser, error)
}
