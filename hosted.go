package rocketfuel

import (
	"encoding/json"
	"fmt"
)

const (
	baseHost = "/hosted-page"
)

type HostedPageService service
type Cart struct {
	Id       string `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Price    string `json:"price,omitempty"`
	Quantity string `json:"quantity,omitempty"`
}
type HostedPageRequest struct {
	Amount      string `json:"amount,omitempty"`
	Cart        []Cart `json:"cart,omitempty"`
	Currency    string `json:"currency,omitempty"`
	Order       string `json:"order,omitempty"`
	RedirectUrl string `json:"redirectUrl,omitempty"`
	Merchant_id string `json:"merchant_id,omitempty"`
}

func (s *HostedPageService) Create(body HostedPageRequest) (Response, error) {
	u := fmt.Sprintf(baseHost)
	resp := Response{}
	mapB, _ := json.Marshal(body)
	err := s.client.Call("POST", u, string(mapB), &resp)

	return resp, err
}
