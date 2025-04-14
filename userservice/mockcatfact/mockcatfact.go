package mockcatfact

import (
	"fmt"

	"github.com/libaishwarya/myapp/userservice"
)

type MockCatFact struct {
	Fail bool
}

func (m *MockCatFact) GetCatFact() (*userservice.CatFact, error) {
	if m.Fail {
		return nil, fmt.Errorf("error fetching cat fact")
	}

	return &userservice.CatFact{
		Fact:   "Cats have five toes on their front paws, but only four on their back paws.",
		Length: 72,
	}, nil
}
