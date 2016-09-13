package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/hypnoglow/go-codewars"
)

// Insert your token here.
const token = ""

func main() {
	cw := codewars.NewClient(token)

	slug := "printing-array-elements-with-comma-delimiters"

	// Fetching kata just to overview it's data.
	// There is no need to do this to proceed the training session.
	fmt.Printf("Fetching kata data for slug `%s`...\n", slug)
	kata, _, err := cw.Katas.GetKata(slug)
	if err != nil {
		fmt.Printf("Error: %v\n\n", err)
		return
	}

	kataJSON, _ := json.MarshalIndent(kata, "", "  ")
	fmt.Printf("Kata:\n%s\n\n", kataJSON)

	// This creates a "ProjectID" and a "SolutionID", so you can submit your
	// solution later.
	trainingSession, _, err := cw.Katas.Train(slug, "javascript")
	if err != nil {
		fmt.Printf("Error: %v\n\n", err)
		return
	}

	tsJSON, _ := json.MarshalIndent(trainingSession, "", "  ")
	fmt.Printf("Training Session:\n%s\n\n", tsJSON)

	// Prepare our solution.
	opts := &codewars.AttemptSolutionOptions{
		Code: "const printArray = (a) => a.join(',');",
	}

	attempt, _, err := cw.Katas.AttemptSolution(
		trainingSession.Session.ProjectID,
		trainingSession.Session.SolutionID,
		opts,
	)
	if err != nil {
		fmt.Printf("Error: %v\n\n", err)
		return
	}
	if attempt.Success == false {
		fmt.Printf("Attempt failed.\n")
		return
	}

	fmt.Println("Sleeping 5 seconds before checking attempt result...")
	time.Sleep(time.Second * 5)

	deferredResponse, _, err := cw.Deferred.GetDeferredResponse(attempt.DMID)
	if err != nil {
		fmt.Printf("Error: %v\n\n", err)
		return
	}

	if deferredResponse.Valid == false {
		fmt.Printf("Attempt failed. Reason: %s\n", deferredResponse.Reason)
		for _, out := range deferredResponse.Output {
			fmt.Println(out)
		}
		return
	}

	fmt.Println("You have passed all the tests! :)")

	final, _, err := cw.Katas.FinalizeSolution(
		trainingSession.Session.ProjectID,
		trainingSession.Session.SolutionID,
	)
	if err != nil {
		fmt.Printf("Error: %v\n\n", err)
		return
	}

	if final.Success == false {
		fmt.Printf("Finalization failed: no attempt found or incorrect request.\n")
		return
	}

	fmt.Println("The solution was successfully finalized!")
}
