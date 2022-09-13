package rocketfuel

import "fmt"

const (
	baseResource = "/auth"
)

type AuthorizatonService service

type AuthorizatonRequest struct {
	email    string `json:"email,omitempty"`
	password string `json:"password,omitempty"`
}

func (s *AuthorizatonService) login() (Response, error) {
	u := fmt.Sprintf(baseResource + "/login")
	resp := Response{}
	cred := s.getMerchantCred()
	err := s.client.Call("POST", u, resp, cred)

	return resp, err
}
