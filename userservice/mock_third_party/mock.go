package mockthirdparty

import (
	"fmt"

	"github.com/libaishwarya/myapp/userservice"
)

type MockThirdParty struct {
	Fail  bool
	Users []userservice.ThirdPartyUser
}

func (m MockThirdParty) GetUsers() ([]userservice.ThirdPartyUser, error) {
	if m.Fail {
		return nil, fmt.Errorf("error returned")
	}

	if m.Users == nil {
		return []userservice.ThirdPartyUser{
			{
				Name:  "test",
				Email: "test",
			},
		}, nil
	}

	return m.Users, nil
}
