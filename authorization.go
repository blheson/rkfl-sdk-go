package rocketfuel

import (
	"fmt"
)

const (
	baseAuth = "/auth"
)

type AuthorizationService service

type AuthorizationRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (s *AuthorizationService) LoginWithPassword() (Response, error) {
	u := fmt.Sprintf(baseAuth + "/login")
	resp := Response{}

	cred := s.client.getMerchantCred()

	err := s.client.Call("POST", u, cred, &resp)

	return resp, err
}
func (s *AuthorizationService) Login() (Response, error) {
	u := fmt.Sprintf(baseAuth + "/generate-auth-token")
	resp := Response{}
	
	cred := s.client.getMerchantCred()

fmt.Println(cred,"Cred: Testing init")

	err := s.client.Call("POST", u, cred, &resp)

	return resp, err
}