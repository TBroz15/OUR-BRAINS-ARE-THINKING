package yt

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/user"
	"path/filepath"

	"github.com/charmbracelet/log"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

// Creates a new YouTube Service by environment variables (YOUTUBE_CLIENT_ID, YOUTUBE_CLIENT_SECRET, YOUTUBE_REDIRECT_URL)
//
// YouTube token is stored in the ~~balls~~ ".ytCred" directory
//
// This function will log fatal if there is any errors.
func CreateYouTubeService() *youtube.Service {
	ctx := context.Background()

	clientID := os.Getenv("YOUTUBE_CLIENT_ID")
	clientSecret := os.Getenv("YOUTUBE_CLIENT_SECRET")
	redirectURL := os.Getenv("YOUTUBE_REDIRECT_URL")

	if clientID == "" || clientSecret == "" {
		log.Fatal("Missing YOUTUBE_CLIENT_ID or YOUTUBE_CLIENT_SECRET environment variables")
	}

	config := &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  redirectURL,
		Scopes:       []string{youtube.YoutubeForceSslScope},
		Endpoint:     google.Endpoint,
	}

	client := getClient(config)

	ytService, err := youtube.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Unable to create YouTube client: %v", err)
	}

	return ytService
}

func getClient(config *oauth2.Config) *http.Client {
	tokFile := tokenCacheFile()
	tok, err := tokenFromFile(tokFile)
	if err != nil {
		tok = getTokenFromWeb(config)
		saveToken(tokFile, tok)
	}
	return config.Client(context.Background(), tok)
}

func tokenCacheFile() string {
	usr, _ := user.Current()
	tokenCacheDir := filepath.Join(usr.HomeDir, ".ytCred")
	os.MkdirAll(tokenCacheDir, 0700)
	return filepath.Join(tokenCacheDir, "youtube-token.json")
}

func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	var token oauth2.Token
	err = json.NewDecoder(f).Decode(&token)
	return &token, err
}

func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Open the following link in your browser then type the authorization code:\n%v\n", authURL)

	var authCode string
	fmt.Print("Enter code: ")
	fmt.Scan(&authCode)

	tok, err := config.Exchange(context.TODO(), authCode)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web: %v", err)
	}
	return tok
}

func saveToken(path string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", path)
	f, _ := os.Create(path)
	defer f.Close()
	json.NewEncoder(f).Encode(token)
}
