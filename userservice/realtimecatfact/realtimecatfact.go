package realtimecatfact

import (
	"encoding/json"
	"net/http"

	"github.com/libaishwarya/myapp/userservice"
)

type RealTimeCatFact struct{}

func (r *RealTimeCatFact) GetCatFact() (*userservice.CatFact, error) {
	resp, err := http.Get("https://catfact.ninja/fact")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var factResponse struct {
		Fact   string `json:"fact"`
		Length int    `json:"length"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&factResponse); err != nil {
		return nil, err
	}

	return &userservice.CatFact{
		Fact:   factResponse.Fact,
		Length: factResponse.Length,
	}, nil
}
