package methods

import (
	"encoding/json"
	"errors"
	"fmt"
)

type registerParams struct {
	Name string `json:"name"`
}

func Register(params json.RawMessage, id int) (interface{}, error) {
	p := registerParams{}
	if err := json.Unmarshal(params, &p); err != nil {
		return "", errors.New("wrong params sent, expected {name:string}")
	}
	_, ok := Storage[p.Name]
	if ok {
		return nil, errors.New(fmt.Sprintf("User %s exist, come up with a new one", p.Name))
	}
	Storage[p.Name] = id
	return fmt.Sprintf("User, %s! Successfully added", p.Name), nil
}