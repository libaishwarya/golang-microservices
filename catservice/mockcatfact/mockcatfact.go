package mockcatfact

import (
	"fmt"

	"github.com/libaishwarya/myapp/catservice"
)

type MockCatFact struct {
	Fail bool
}

func (m *MockCatFact) GetCatFact() (catservice.CatFact, error) {
	if m.Fail {
		return catservice.CatFact{}, fmt.Errorf("error fetching cat fact")
	}

	return catservice.CatFact{
		Fact:   "Cats have five toes on their front paws, but only four on their back paws.",
		Length: 72,
	}, nil
}
