package main

import (
	"time"
	"fmt"
	"net/http"
	"log"
	"encoding/json"
)

type QResponse struct {
	ID        string `json:"id"`
	Timestamp time.Time `json:"timestamp"`
	Result struct {
		Source           string `json:"source"`
		Resolvedquery    string `json:"resolvedquery"`
		Action           string `json:"action"`
		Actionincomplete bool `json:"actionincomplete"`
		Parameter struct {
			Name string `json:"name"`
		} `json:"parameter"`
		Contexts []struct {
			Name string `json:"name"`
			Parameter struct {
				Name string `json:"name"`
			} `json:"parameter"`
			Lifespan int `json:"lifespan"`
		} `json:"contexts"`
		Metadata struct {
			IntentID   string `json:"intent_id"`
			IntentName string `json:"intent_name"`
		}`json:"metadata"`
		Fulfillment struct {
			Speech string `json:"speech"`
		}`json:"fulfillment"`
	}`json:"result"`
	Status struct {
		Code      int `json:"code"`
		ErrorType string `json:"error_type"`
	}`json:"status"`
}

func main() {

	url := fmt.Sprintf("https://api.api.ai/v1/query?v=20150910&query=hi&lang=en&sessionId=1234567890")
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("New Request ", err)
		return
	}

	req.Header.Add("Authorization", "Bearer 648ccd20a2344205910a6fd2fda0baf5")

	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		log.Fatal(" error", err)
		return
	}

	defer resp.Body.Close()

	// fill the record with data receive as json
	var record QResponse

	// Use json.Decode for reading streams of JSON data
	if err := json.NewDecoder(resp.Body).Decode(&record); err != nil {
		log.Println(err)
	}

	fmt.Println("Status =", record.Status.Code)
	fmt.Println("Response =", record.Result.Fulfillment.Speech)

}
