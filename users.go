package codewars

import (
	"fmt"
	"net/http"
)

const usersResource = "users"

// UsersService handles communication with the user related methos
// of Codewars API.
type UsersService struct {
	client *Client
}

// User represents a Codewars user.
type User struct {
	Username            string   `json:"username"`
	Name                string   `json:"name"`
	Honor               int      `json:"honor"`
	Clan                string   `json:"clan"`
	LeaderboardPosition int      `json:"leaderboardPosition"`
	Skills              []string `json:"skills"`
	Ranks               struct {
		Overall   *UserRank            `json:"overall"`
		Languages map[string]*UserRank `json:"languages"`
	} `json:"ranks"`
	CodeChallenges struct {
		TotalAuthored  int `json:"totalAuthored"`
		TotalCompleted int `json:"totalCompleted"`
	} `json:"codeChallenges"`
}

// UserRank represents user's overall ranking or his ranking in a particular language.
type UserRank struct {
	Rank  int    `json:"rank"`
	Name  string `json:"name"`
	Color string `json:"color"`
	Score int    `json:"score"`
}

// GetUser gets a single user.
func (s *UsersService) GetUser(username string) (*User, *http.Response, error) {
	url := fmt.Sprintf("%s/%s", usersResource, username)

	req, err := s.client.NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	user := new(User)
	res, err := s.client.Do(req, user)
	if err != nil {
		return nil, res, err
	}

	return user, res, nil
}
