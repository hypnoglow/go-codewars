package codewars

import (
	"fmt"
	"net/http"
)

const deferredResource = "deferred"

// DeferredService handles communication with the deferred related methos
// of Codewars API.
type DeferredService struct {
	client *Client
}

// DeferredResponse represents a Codewars deferred response.
type DeferredResponse struct {
	Success  bool     `json:"success"`
	DMID     string   `json:"dmid"`
	Valid    bool     `json:"valid"`
	Reason   string   `json:"reason"`
	Output   []string `json:"output"`
	WallTime int      `json:"wall_time"`
}

// GetDeferredResponse is used for polling for a deferred response.
// This is to be used in conjunction with the attempt endpoint.
//
// Polling should not be performed more than twice a second.
// If a API consumer abuses the rate guidelines,
// its IP address will be temporarily suspended.
func (s *DeferredService) GetDeferredResponse(DMID string) (*DeferredResponse, *http.Response, error) {
	url := fmt.Sprintf("%s/%s", deferredResource, DMID)

	req, err := s.client.NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	dr := new(DeferredResponse)
	res, err := s.client.Do(req, dr)
	if err != nil {
		return nil, res, err
	}

	return dr, res, nil
}
