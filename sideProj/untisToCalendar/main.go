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
	"time"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
)

func main() {
	// Create a context for API calls
	ctx := context.Background()

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

	// Get an authenticated HTTP client
	client := getClient(ctx, config)

	// Create a new Calendar service using the authenticated client
	srv, err := calendar.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Error creating Calendar service: %v", err)
	}

	// Define a test event to create
	event := &calendar.Event{
		Summary:     "WebUntis test event automatic",
		Description: "Third event created from Go but this time automatic",
		Start: &calendar.EventDateTime{
			DateTime: time.Now().Add(1 * time.Hour).Format(time.RFC3339),
		},
		End: &calendar.EventDateTime{
			DateTime: time.Now().Add(2 * time.Hour).Format(time.RFC3339),
		},
	}

	// Insert the event into the primary calendar
	created, err := srv.Events.Insert("primary", event).Do()
	if err != nil {
		log.Fatalf("Error creating event: %v", err)
	}

	fmt.Println("Event created:", created.HtmlLink)
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
	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request)  {
		// readind the "code..." part of the URL
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
