package codewars

import (
	"fmt"
	"net/http"
)

// TrainingSession represents Codewars API Training Session, which includes
// Kata information and Session data.
type TrainingSession struct {
	Name string
	Slug string

	// The code challenge description is given in markdown format.
	Description string

	Author            string
	Rank              int
	AverageCompletion float64
	Tags              []string
	Session           *Session
}

// Session represents Codewars API Session.
type Session struct {
	// This ID value will be needed when submitting a solution.
	ProjectID string
	// This ID value will be needed when submitting a solution.
	SolutionID string
	// This is the initial solution code that is given to a user.
	Setup string
	// This is the example test cases that are initially given to a user.
	ExampleFixture string
	// If the user is continuing a previous started solution,
	// this value will represent their previous work.
	Code string
}

// StrategyType represents a value for selecting a code challenge selection
// strategy.
//
// All strategies will prefer incomplete and non-started code challenges,
// unless otherwise stated.
type StrategyType string

const (
	// StrategyDefault also referred to as the “Rank Up” workout.
	// Will select a challenge that is above your current level.
	StrategyDefault StrategyType = "default"

	// StrategyRandom randomly selected code challenges.
	StrategyRandom StrategyType = "random"

	// StrategyReferenceWorkout will select code challenges
	// that are tagged as reference.
	StrategyReferenceWorkout StrategyType = "reference_workout"

	// StrategyBetaWorkout will select beta code challenges.
	StrategyBetaWorkout StrategyType = "beta_workout"

	// StrategyRetrainWorkout will focus on code challenges
	// that you have already completed.
	StrategyRetrainWorkout StrategyType = "retrain_workout"

	// StrategyAlgorithmRetest will focus on algorithm code challenges
	// that you have already completed.
	StrategyAlgorithmRetest StrategyType = "algorithm_retest"

	// StrategyKyu8Workout will focus on 8 kyu code challenges.
	StrategyKyu8Workout StrategyType = "kyu_8_workout"

	// StrategyKyu7Workout will focus on 7 kyu code challenges.
	StrategyKyu7Workout StrategyType = "kyu_7_workout"

	// StrategyKyu6Workout will focus on 6 kyu code challenges.
	StrategyKyu6Workout StrategyType = "kyu_6_workout"

	// StrategyKyu5Workout will focus on 5 kyu code challenges.
	StrategyKyu5Workout StrategyType = "kyu_5_workout"

	// StrategyKyu4Workout will focus on 4 kyu code challenges.
	StrategyKyu4Workout StrategyType = "kyu_4_workout"

	// StrategyKyu3Workout will focus on 3 kyu code challenges.
	StrategyKyu3Workout StrategyType = "kyu_3_workout"

	// StrategyKyu2Workout will focus on 2 kyu code challenges.
	StrategyKyu2Workout StrategyType = "kyu_2_workout"

	// StrategyKyu1Workout will focus on 1 kyu code challenges.
	StrategyKyu1Workout StrategyType = "kyu_1_workout"
)

// TrainNextOptions represents the available TrainNext() options.
type TrainNextOptions struct {
	// The strategy to use for choosing what the next code challenge should be.
	// (Optional)
	Strategy StrategyType `url:"strategy,omitempty"`

	// True if you only want to peek at the next item in your queue,
	// without removing it from the queue or beginning a new training session.
	// (Optional)
	Peek bool `url:"peek,omitempty"`
}

// TrainNext begins a new training session for the next code challenge (kata)
// within user's training queue.
//
// If the next code challenge within your queue is one that
// you have not started yet, then a timer will begin as soon as
// the request is made. The timer is used solely for tracking
// average completion times and does not affect your honor in any way.
func (s *KatasService) TrainNext(lang string, opts *TrainNextOptions) (*TrainingSession, *http.Response, error) {
	url := fmt.Sprintf("%s/%s/train", katasResource, lang)

	req, err := s.client.NewRequest("POST", url, opts)
	if err != nil {
		return nil, nil, err
	}

	trainingSession := new(TrainingSession)
	res, err := s.client.Do(req, trainingSession)
	if err != nil {
		return nil, res, err
	}

	return trainingSession, res, err
}

// Train begins a new training session for the specified code challenge (kata).
//
// If the code challenge has not been started by yet by the user,
// then a timer will begin as soon as the request is made.
// The timer is used solely for tracking average completion times
// and does not affect the user’s honor in any way.
func (s *KatasService) Train(slug, lang string) (*TrainingSession, *http.Response, error) {
	url := fmt.Sprintf("%s/%s/%s/train", katasResource, slug, lang)

	req, err := s.client.NewRequest("POST", url, nil)
	if err != nil {
		return nil, nil, err
	}

	trainingSession := new(TrainingSession)
	res, err := s.client.Do(req, trainingSession)
	if err != nil {
		return nil, res, err
	}

	return trainingSession, res, err
}
