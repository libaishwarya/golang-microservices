package userservice

type CatFact struct {
	Fact   string `json:"fact"`
	Length int    `json:"length"`
}

type CatFactService interface {
	GetCatFact() (*CatFact, error)
}
