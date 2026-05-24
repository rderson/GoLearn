// Package provides a program to create events on Google Calendar using OAuth2 authentication.
// It handles authentication flow, token management, and event creation.
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"

	"untisToCalendar/untis"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
)

func main() {
	// Create a context for API calls
	ctx := context.Background()

	// Authenticate with WebUntis and fetch timetable
	httpClient := &http.Client{}

	untisClient := untis.NewClient(
		"TorRas",         // TODO: Move to .env
		"Behördestinkt1", // TODO: Move to .env
		"https://gotthard-kuehl-schule.webuntis.com/WebUntis/jsonrpc.do",
		949, // TODO: Move to .env
	)

	// Authenticate with WebUntis
	sessionID, err := untisClient.Authenticate(httpClient)
	if err != nil {
		log.Fatalf("Error authenticating with WebUntis: %v", err)
	}
	fmt.Println("Authenticated with WebUntis, Session ID:", sessionID)

	// Fetch timetable for the next 14 days
	startDate := time.Now()
	endDate := startDate.AddDate(0, 0, 14)

	lessons, err := untisClient.GetTimetable(httpClient, sessionID, startDate, endDate)
	if err != nil {
		log.Fatalf("Error fetching timetable: %v", err)
	}
	fmt.Printf(" Fetched %d lessons from WebUntis\n", len(lessons))

	// Fetch subject and teacher mappings
	subjects, err := untisClient.GetSubjects(httpClient, sessionID)
	if err != nil {
		log.Fatalf("Error fetching subjects: %v", err)
	}
	teachersMap, err := untisClient.GetTeachers(httpClient, sessionID)
	if err != nil {
		log.Fatalf("Error fetching teachers: %v", err)
	}

	// Set up Google Calendar

	// Read Google API credentials from file
	b, err := os.ReadFile("credentials.json")
	if err != nil {
		log.Fatalf("Error reading credentials.json: %v", err)
	}

	// Parse the credentials and configure OAuth2 for Calendar scope
	config, err := google.ConfigFromJSON(b, calendar.CalendarScope)
	if err != nil {
		log.Fatalf("Error parsing credentials.json: %v", err)
	}

	// Get an authenticated HTTP client for Google
	googleClient := getClient(ctx, config)

	// Create a new Calendar service using the authenticated client
	srv, err := calendar.NewService(ctx, option.WithHTTPClient(googleClient))
	if err != nil {
		log.Fatalf("Error creating Calendar service: %v", err)
	}

	// Find or create the WebUntis calendar
	calendarID, err := getOrCreateCalendar(srv)
	if err != nil {
		log.Fatalf("Error getting calendar: %v", err)
	}
	fmt.Printf("Using calendar ID: %s\n", calendarID)

	// Remove existing events in this calendar for the same date range (recreate)
	fmt.Println("Removing existing events in WebUntis calendar for the date range...")
	timeMin := startDate.Format(time.RFC3339)
	timeMax := endDate.Format(time.RFC3339)
	evList, err := srv.Events.List(calendarID).
		ShowDeleted(false).
		SingleEvents(true).
		TimeMin(timeMin).
		TimeMax(timeMax).
		Do()
	if err != nil {
		log.Printf("Warning: could not list existing events: %v", err)
	} else {
		for _, ev := range evList.Items {
			if err := srv.Events.Delete(calendarID, ev.Id).Do(); err != nil {
				log.Printf("Warning: could not delete event %s: %v", ev.Id, err)
			}
		}
	}

	// Convert lessons to calendar events

	createdCount := 0
	for _, lesson := range lessons {
		// Skip cancelled lessons
		if lesson.Code == "cancelled" {
			continue
		}

		// Convert lesson dates and times to datetime
		eventStart, eventEnd := lessonToDateTime(lesson)

		// Resolve subject name (prefer LongName then Name) and teacher abbreviations
		subjectName := "Untis Lesson"
		if len(lesson.Subjects) > 0 {
			if s, ok := subjects[lesson.Subjects[0].ID]; ok {
				if s.LongName != "" {
					subjectName = s.LongName
				} else if s.Name != "" {
					subjectName = s.Name
				}
			}
		}

		var teacherAbbrs []string
		for _, t := range lesson.Teachers {
			if te, ok := teachersMap[t.ID]; ok {
				if te.Abbreviation != "" {
					teacherAbbrs = append(teacherAbbrs, te.Abbreviation)
				} else if te.Name != "" {
					teacherAbbrs = append(teacherAbbrs, te.Name)
				}
			}
		}

		// Create calendar event with subject as title and teacher list in description
		event := &calendar.Event{
			Summary:     subjectName,
			Description: fmt.Sprintf("Teachers: %s\nWebUntis ID: %d", strings.Join(teacherAbbrs, ", "), lesson.ID),
			Start: &calendar.EventDateTime{
				DateTime: eventStart.Format(time.RFC3339),
			},
			End: &calendar.EventDateTime{
				DateTime: eventEnd.Format(time.RFC3339),
			},
		}

		// Insert the event into the WebUntis calendar
		created, err := srv.Events.Insert(calendarID, event).Do()
		if err != nil {
			log.Printf("Warning: Error creating event: %v\n", err)
			continue
		}
		createdCount++
		fmt.Printf("Created event: %s\n", created.HtmlLink)
	}

	fmt.Printf("\nSuccessfully created %d calendar events\n", createdCount)
}

// getOrCreateCalendar finds or creates a calendar named "Untis" and returns its ID
func getOrCreateCalendar(srv *calendar.Service) (string, error) {
	// List all calendars to find if WebUntis already exists
	calendarList, err := srv.CalendarList.List().Do()
	if err != nil {
		return "", fmt.Errorf("error listing calendars: %v", err)
	}

	// Search for existing "Untis" calendar
	for _, cal := range calendarList.Items {
		if cal.Summary == "Untis" {
			fmt.Println("Found existing Untis calendar")
			return cal.Id, nil
		}
	}

	// Create new calendar if it doesn't exist
	fmt.Println("Creating new Untis calendar...")
	newCal := &calendar.Calendar{
		Summary:     "Untis",
		Description: "Automatic synced timetable from WebUntis",
	}

	created, err := srv.Calendars.Insert(newCal).Do()
	if err != nil {
		return "", fmt.Errorf("error creating calendar: %v", err)
	}

	fmt.Println("Created new WebUntis calendar")
	return created.Id, nil
}

// lessonToDateTime converts a WebUntis lesson to start and end datetime.Time objects
func lessonToDateTime(lesson untis.Lesson) (startTime, endTime time.Time) {
	// Parse date from YYYYMMDD format
	year := lesson.Date / 10000
	month := (lesson.Date % 10000) / 100
	day := lesson.Date % 100

	// Parse time from HHMM format
	startHour := lesson.StartTime / 100
	startMin := lesson.StartTime % 100
	endHour := lesson.EndTime / 100
	endMin := lesson.EndTime % 100

	// Create time objects with Berlin timezone
	loc, _ := time.LoadLocation("Europe/Berlin")
	startTime = time.Date(year, time.Month(month), day, startHour, startMin, 0, 0, loc)
	endTime = time.Date(year, time.Month(month), day, endHour, endMin, 0, 0, loc)

	return startTime, endTime
}

// getClient retrieves an authenticated HTTP client for Google APIs.
// It first tries to load a saved token from file, and if that fails,
// it initiates the OAuth2 flow to get a new token.
func getClient(ctx context.Context, config *oauth2.Config) *http.Client {
	tokFile := "token.json"
	tok, err := tokenFromFile(tokFile)
	if err != nil {
		// No saved token, need to get a new one
		tok = getTokenFromWeb(config)
		saveToken(tokFile, tok)
	}
	return config.Client(ctx, tok)
}

// getTokenFromWeb performs the OAuth2 authorization flow
// it opens the authorization URL in the browser, prompts the user to enter the authorization code,
// and exchanges it for an access token
func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	// channel to recieve the authorization code
	codeChan := make(chan string)

	// set the redirect URI to our local server where Google will send the authorization code
	config.RedirectURL = "http://localhost:8080/"

	// creating a handle function for our local server to copy the authorization code from the url and send it to the channel that was created earlier
	// so the user doesn't need  to paste the authorization code manually
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// reading the "code..." part of the URL
		code := r.URL.Query().Get("code")
		if code == "" {
			http.Error(w, "Authorization code not found.", 400)
			return
		}

		// sending the authorization code to the channel
		codeChan <- code

		// showing a user-friendly message on the page
		w.Header().Set("Content-Type", "text/html")
		fmt.Fprintf(w, "<h1>Authorization successful!</h1><p>You can close this window and return to the application.</p>")
	})

	// creating an adressable http server
	server := &http.Server{
		Addr: ":8080",
	}

	// starting the http server in a goroutine
	go func() {
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Println("Server error: ", err)
		}
	}()

	// generate the authorization URL
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)

	// open the URL in the default browser
	fmt.Println("Opening browser for authorization...")
	openBrowser(authURL)

	// wait for the authorization code from the callback
	fmt.Println("Waiting for authorization...")
	code := <-codeChan

	// shutting down the server
	if err := server.Shutdown(context.Background()); err != nil {
		log.Println("Error shutting down the server:", err)
	}

	// exchange the code for a token
	tok, err := config.Exchange(context.Background(), code)
	if err != nil {
		log.Fatalf("Error retrieving token: %v", err)
	}
	return tok
}

// tokenFromFile loads an OAuth2 token from a JSON file.
// Returns the token and any error encountered during loading.
func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var tok oauth2.Token
	err = json.NewDecoder(f).Decode(&tok)
	return &tok, err
}

// saveToken saves an OAuth2 token to a JSON file for future use.
func saveToken(path string, token *oauth2.Token) {
	f, err := os.Create(path)
	if err != nil {
		log.Fatalf("Error saving token: %v", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
}

// openBrowser opens the given URL in the default web browser.
// It uses platform-specific commands to launch the browser.
func openBrowser(url string) {
	var cmd string
	var args []string

	// Determine the command based on the operating system
	switch runtime.GOOS {
	case "windows":
		cmd = "rundll32"
		args = []string{"url.dll,FileProtocolHandler", url}
	case "darwin":
		cmd = "open"
		args = []string{url}
	default: // Linux and other Unix-like systems
		cmd = "xdg-open"
		args = []string{url}
	}

	// Execute the command to open the browser
	exec.Command(cmd, args...).Start()
}
