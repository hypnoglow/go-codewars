package codewars

import (
	"fmt"
	"net/http"
)

// SolutionResponse represents the Codewars API response on solution POSTing.
type SolutionResponse struct {
	Success bool
}

// SolutionAttemptResponse represents the Codewars API response on solution
// attempt, which contains a deferred message id (dmid) which will be used to
// poll for the response.
type SolutionAttemptResponse struct {
	SolutionResponse
	DMID string
}

// OutputFormat is the output format to be used by AttemptSolution().
type OutputFormat string

const (
	// OutputFormatHTML is a HTML format.
	OutputFormatHTML OutputFormat = "html"

	// OutputFormatRaw is a raw format.
	OutputFormatRaw OutputFormat = "raw"
)

// AttemptSolutionOptions represents the available AttemptSolution() options.
type AttemptSolutionOptions struct {
	// The code that you is being submitted
	Code string `url:"code"`

	// The output format to be used.
	// (Optional)
	OutputFormat OutputFormat `url:"output_format,omitempty"`
}

// AttemptSolution is used to submit a solution to be validated by the code
// challenge author’s test cases. It will return a deferred message id (dmid)
// which will be used to poll for the response. Polling must be used to retrieve
// the response.
func (s *KatasService) AttemptSolution(projectID, solutionID string, opts *AttemptSolutionOptions) (*SolutionAttemptResponse, *http.Response, error) {
	url := fmt.Sprintf(
		"%s/projects/%s/solutions/%s/attempt",
		katasResource,
		projectID,
		solutionID,
	)

	req, err := s.client.NewRequest("POST", url, opts)
	if err != nil {
		return nil, nil, err
	}

	solutionAttempt := new(SolutionAttemptResponse)
	res, err := s.client.Do(req, solutionAttempt)
	if err != nil {
		return nil, res, err
	}

	return solutionAttempt, res, err
}

// FinalizeSolution is used to finalize the previously submitted solution.
// This endpoint will only return a success message if there has been
// a previously successful solution.
func (s *KatasService) FinalizeSolution(projectID, solutionID string) (*SolutionResponse, *http.Response, error) {
	url := fmt.Sprintf(
		"%s/projects/%s/solutions/%s/finalize",
		katasResource,
		projectID,
		solutionID,
	)

	req, err := s.client.NewRequest("POST", url, nil)
	if err != nil {
		return nil, nil, err
	}

	solutionResponse := new(SolutionResponse)
	res, err := s.client.Do(req, solutionResponse)
	if err != nil {
		return nil, res, err
	}

	return solutionResponse, res, err
}
