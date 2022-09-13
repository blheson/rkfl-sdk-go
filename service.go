package rocketfuel

import (
	"encoding/json"
	"fmt"
)

type Service service

func (s *Service) GetUUID(body HostedPageRequest) (Response, error) {
	u := fmt.Sprintf(baseHost)
	resp := Response{}
	mapB, _ := json.Marshal(body)
	err := s.client.Call("POST", u, string(mapB), &resp)

	return resp, err
}
