package untis

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"
)

// Client holds WebUntis API configuration
type Client struct {
	Username string
	Password string
	BaseURL  string
	ClassID  int
}

// NewClient creates a new WebUntis client with the given configuration
func NewClient(username, password, baseURL string, classID int) *Client {
	return &Client{
		Username: username,
		Password: password,
		BaseURL:  baseURL,
		ClassID:  classID,
	}
}

// Authenticate logs into user's Untis account and returns a session ID
func (c *Client) Authenticate(httpClient *http.Client) (string, error) {
	body := map[string]any{
		"jsonrpc": "2.0",
		"method":  "authenticate",
		"params": map[string]any{
			"user":     c.Username,
			"password": c.Password,
			"client":   "webuntis-sync",
		},
		"id": 1,
	}

	b, _ := json.Marshal(body)

	req, err := http.NewRequest("POST", c.BaseURL, bytes.NewReader(b))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := httpClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	raw, _ := io.ReadAll(resp.Body)

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

// GetTimetable fetches the timetable from WebUntis for a given date range and returns parsed lesson data
func (c *Client) GetTimetable(httpClient *http.Client, sessionID string, startDate, endDate time.Time) ([]Lesson, error) {
	startDateInt, _ := strconv.Atoi(startDate.Format("20060102"))
	endDateInt, _ := strconv.Atoi(endDate.Format("20060102"))

	payload := map[string]any{
		"jsonrpc": "2.0",
		"id":      2,
		"method":  "getTimetable",
		"params": map[string]any{
			"id":        c.ClassID,
			"type":      1,
			"startDate": startDateInt,
			"endDate":   endDateInt,
		},
	}

	b, _ := json.Marshal(payload)

	req, err := http.NewRequest("POST", c.BaseURL, bytes.NewReader(b))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Cookie", "JSESSIONID="+sessionID)

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("WebUntis returned status %d", resp.StatusCode)
	}

	raw, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result TimetableResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}

	return result.Result, nil
}

// GetSubjects fetches subject list from WebUntis
func (c *Client) GetSubjects(httpClient *http.Client, sessionID string) (map[int]Subject, error) {
	payload := map[string]any{
		"jsonrpc": "2.0",
		"id":      3,
		"method":  "getSubjects",
		"params":  map[string]any{},
	}

	b, _ := json.Marshal(payload)
	req, err := http.NewRequest("POST", c.BaseURL, bytes.NewReader(b))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Cookie", "JSESSIONID="+sessionID)

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	raw, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result struct {
		Result []Subject `json:"result"`
	}
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}

	m := make(map[int]Subject, len(result.Result))
	for _, s := range result.Result {
		m[s.ID] = s
	}
	return m, nil
}

// GetTeachers fetches teacher list from WebUntis
func (c *Client) GetTeachers(httpClient *http.Client, sessionID string) (map[int]Teacher, error) {
	payload := map[string]any{
		"jsonrpc": "2.0",
		"id":      4,
		"method":  "getTeachers",
		"params":  map[string]any{},
	}

	b, _ := json.Marshal(payload)
	req, err := http.NewRequest("POST", c.BaseURL, bytes.NewReader(b))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Cookie", "JSESSIONID="+sessionID)

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	raw, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result struct {
		Result []Teacher `json:"result"`
	}
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}

	m := make(map[int]Teacher, len(result.Result))
	for _, t := range result.Result {
		m[t.ID] = t
	}
	return m, nil
}
