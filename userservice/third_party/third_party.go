package thirdparty

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/libaishwarya/myapp/userservice"
)

type ThirdParty struct {
}

func (th *ThirdParty) GetUsers() ([]userservice.ThirdPartyUser, error) {

	resp, err := http.Get("https://jsonplaceholder.typicode.com/users")
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var users []userservice.ThirdPartyUser
	if err := json.Unmarshal(body, &users); err != nil {
		return nil, err
	}

	return users, nil

}
