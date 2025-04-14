package realtimecatfact

import (
	"encoding/json"
	"net/http"

	"github.com/libaishwarya/myapp/catservice"
)

type RealTimeCatFact struct{}

func (r *RealTimeCatFact) GetCatFact() (catservice.CatFact, error) {
	resp, err := http.Get("https://catfact.ninja/fact")
	if err != nil {
		return catservice.CatFact{}, err
	}
	defer resp.Body.Close()

	var factResponse struct {
		Fact   string `json:"fact"`
		Length int    `json:"length"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&factResponse); err != nil {
		return catservice.CatFact{}, err
	}

	return catservice.CatFact{
		Fact:   factResponse.Fact,
		Length: factResponse.Length,
	}, nil
}
