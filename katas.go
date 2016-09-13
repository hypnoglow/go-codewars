package codewars

import (
	"fmt"
	"net/http"
	"time"
)

const katasResource = "code-challenges"

// KatasService handles communication with the code challenges (katas) related
// methos of Codewars API.
type KatasService struct {
	client *Client
}

// Kata represents a Codewars code challenge (kata).
type Kata struct {
	ID          string     `json:"id"`
	Name        string     `json:"name"`
	Slg         string     `json:"slug"`
	Category    string     `json:"category"`
	PublishedAt *time.Time `json:"publishedAt"`
	ApprovedAt  *time.Time `json:"approvedAt"`
	Languages   []string   `json:"languages"`
	URL         string     `json:"url"`
	Rank        struct {
		ID    int    `json:"id"`
		Name  string `json:"name"`
		Color string `json:"color"`
	} `json:"rank"`
	CreatedAt          *time.Time    `json:"createdAt"`
	CreatedBy          *KataUserData `json:"createdBy"`
	ApprovedBy         *KataUserData `json:"approvedBy"`
	Description        string        `json:"description"`
	TotalAttempts      int           `json:"totalAttempts"`
	TotalCompleted     int           `json:"totalCompleted"`
	TotalStars         int           `json:"totalStars"`
	VoteScore          int           `json:"voteScore"`
	Tags               []string      `json:"tags"`
	ContributorsWanted bool          `json:"contributorsWanted"`
	Unresolved         struct {
		Issues      int `json:"issues"`
		Suggestions int `json:"suggestions"`
	} `json:"unresolved"`
}

// KataUserData represents user info inside kata structure.
type KataUserData struct {
	Username string `json:"username"`
	URL      string `json:"url"`
}

// GetKata gets a single code challenge (kata).
func (s *KatasService) GetKata(slug string) (*Kata, *http.Response, error) {
	url := fmt.Sprintf("%s/%s", katasResource, slug)

	req, err := s.client.NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	kata := new(Kata)
	res, err := s.client.Do(req, kata)
	if err != nil {
		return nil, res, err
	}

	return kata, res, nil
}
