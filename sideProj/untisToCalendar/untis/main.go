package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"
)

// Creating variables that store user data, it's a temperory solution, would be a lot smarter to store the date as an .env values
const (
	// school   = "gotthard-kuehl-schule"
	username = "TorRas"
	password = "Behördestinkt1"
)

// Logging into user's Untis account and gets the timetable, panicking if there are any errors 
func main() {
	client := &http.Client{}

	baseURL := "https://gotthard-kuehl-schule.webuntis.com/WebUntis/jsonrpc.do"

	// setting start day to current day and end date two weeks later
	startDate := time.Now()
	endDate := startDate.AddDate(0, 0, 14)

	// logging into user's account
	session, err := authenticate(client, baseURL)
	if err != nil {
		panic(err)
	}

	fmt.Println("Session ID:", session)

	// converting both start and end date to integers, errors are not to be expected
	startDateConv, _ := strconv.Atoi(startDate.Format("20060102"))
	endDateConv, _ := strconv.Atoi(endDate.Format("20060102"))

	// getting the timetable
	err = getTimetable(client, baseURL, session, 949, startDateConv, endDateConv)
	if err != nil {
		panic(err)
	}
}

// Logging into user's account
func authenticate(client *http.Client, url string) (string, error) {

	// our POST request
	body := map[string]any{
		"jsonrpc": "2.0",
		"method":  "authenticate",
		"params": map[string]any{
			"user":     username,
			"password": password,
			"client":   "webuntis-sync",
		},
		"id": 1,
	}

	b, _ := json.Marshal(body)

	req, err := http.NewRequest("POST", url, bytes.NewReader(b))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	//fmt.Println("HTTP status:", resp.Status)

	raw, _ := io.ReadAll(resp.Body)
	//fmt.Println(string(raw))

	// saving and returning the session ID, now we know if user is succesfully logged into his account or not
	var result struct {
		Result struct {
			SessionID string `json:"sessionId"`
		} `json:"result"`
	}

	if err := json.Unmarshal(raw, &result); err != nil {
		return "", err
	}

	return result.Result.SessionID, nil
}

// Prints timetable from startDate to endDate
func getTimetable(client *http.Client, url, sessionID string, klasseID int, startDate int, endDate int) error {

	payload := map[string]any{
		"jsonrpc": "2.0",
		"id":      2,
		"method":  "getTimetable",
		"params": map[string]any{
			"id":        klasseID,
			"type":      1,
			"startDate": startDate,
			"endDate":   endDate,
		},
	}

	b, _ := json.Marshal(payload)

	req, err := http.NewRequest("POST", url, bytes.NewReader(b))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Cookie", "JSESSIONID="+sessionID)

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	fmt.Println("Timetable HTTP status:", resp.Status)

	raw, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	fmt.Println(string(raw))

	return nil
}
