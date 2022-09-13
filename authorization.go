package rocketfuel

const (
	baseResource = "/auth"
)

type AuthorizatonService service

type AuthorizatonRequest struct {
}

func (s *AuthorizatonService) login() (Response, error) {
	u := baseResource + "/login"
	resp := Response{}
}
