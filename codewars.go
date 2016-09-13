package codewars

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/google/go-querystring/query"
)

const (
	libraryVersion = "0.0.0"
	defaultBaseURL = "https://www.codewars.com/api/v1/"
	userAgent      = "go-codewars/" + libraryVersion
)

// Client represents a Codewars API client.
type Client struct {
	// HTTP client used to communicate with the API.
	client *http.Client

	// Base URL for API requests.
	baseURL *url.URL

	// Token used to make authenticated API calls.
	token string

	// User agent used when communicating with the Codewars API.
	UserAgent string

	// Services used for talking to different parts of the Codewars API.
	Users    *UsersService
	Katas    *KatasService
	Deferred *DeferredService
}

// NewClient returns a new Codewars API client.
func NewClient(token string) *Client {
	u, err := url.Parse(defaultBaseURL)
	if err != nil {
		panic(err)
	}

	c := &Client{
		client:    http.DefaultClient,
		baseURL:   u,
		token:     token,
		UserAgent: userAgent,
	}

	c.Users = &UsersService{client: c}
	c.Katas = &KatasService{client: c}
	c.Deferred = &DeferredService{client: c}

	return c
}

// NewRequest returns an API request.
func (c *Client) NewRequest(method, path string, opts interface{}) (*http.Request, error) {
	u := *c.baseURL
	u.Opaque = c.baseURL.Path + path

	req := &http.Request{
		Method:     method,
		URL:        &u,
		Proto:      "HTTP/1.1",
		ProtoMajor: 1,
		ProtoMinor: 1,
		Header:     make(http.Header),
		Host:       u.Host,
	}

	if method == "POST" {
		data, err := query.Values(opts)
		if err != nil {
			return nil, err
		}

		encodedData := data.Encode()
		bodyReader := strings.NewReader(encodedData)
		req.Body = ioutil.NopCloser(bodyReader)
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		req.Header.Set("Content-Length", strconv.Itoa(len(encodedData)))
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", c.token)
	req.Header.Set("User-Agent", c.UserAgent)

	return req, nil
}

// Do executes the HTTP request and returns the response.
func (c *Client) Do(req *http.Request, v interface{}) (*http.Response, error) {
	res, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	err = CheckResponse(res)
	if err != nil {
		return res, err
	}

	if v != nil {
		if w, ok := v.(io.Writer); ok {
			_, err = io.Copy(w, res.Body)
		} else {
			err = json.NewDecoder(res.Body).Decode(v)
		}
	}

	return res, err
}

// ErrorResponse represents a Codewars API error.
type ErrorResponse struct {
	Response *http.Response
	Success  bool   `json:"success"`
	Reason   string `json:"reason"`
}

func (er *ErrorResponse) Error() string {
	path, _ := url.QueryUnescape(er.Response.Request.URL.Opaque)
	requestURL := fmt.Sprintf(
		"%s://%s%s",
		er.Response.Request.URL.Scheme,
		er.Response.Request.URL.Host,
		path,
	)

	return fmt.Sprintf(
		"%v %s : %d %s",
		er.Response.Request.Method,
		requestURL,
		er.Response.StatusCode,
		er.Reason,
	)
}

// CheckResponse checks the Codewars API response for errors and returns them
// if present. A response is considered an error if it has a status code outside
// the 200 range.
func CheckResponse(res *http.Response) error {
	if res.StatusCode >= 200 && res.StatusCode <= 299 {
		return nil
	}

	errorResponse := &ErrorResponse{Response: res}
	data, err := ioutil.ReadAll(res.Body)
	if err == nil && data != nil {
		json.Unmarshal(data, errorResponse)
	}

	return errorResponse
}
