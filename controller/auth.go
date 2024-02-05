package controller

import (
	"bufio"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/Swayamshu/mangadex/model"
)

// TODO: move CLIENT_ID and CLIENT_SECRET to .env file
const CLIENT_ID = "253ee8dc9687c8d608e29abc28328748"
const CLIENT_SECRET = "653443bec496593d4e82cec4aa62b7f1ad5e0d19ae91c7a7166d20ed1fe67c55"
const REDIRECT_URI = "https://myanimelist.net"
const BaseAuthorizeURL = "https://myanimelist.net/v1/oauth2/authorize"
const BaseTokenURL = "https://myanimelist.net/v1/oauth2/token"

func GetNewCodeVerifier() string {
	randomBytes := make([]byte, 100)
	_, err := rand.Read(randomBytes)

	if err != nil {
		log.Fatal("Error getting code verifier.", err)
	}

	codeVerifier := base64.URLEncoding.EncodeToString(randomBytes)

	if len(codeVerifier) > 128 {
		codeVerifier = codeVerifier[:128]
	}

	return codeVerifier
}

func GetAuthorisationURL(codeChallenge string) string {
	authorizeURL := BaseAuthorizeURL + fmt.Sprintf("?response_type=code&client_id=%s&code_challenge=%s", CLIENT_ID, codeChallenge)
	return authorizeURL
}

func GetCodeFromRedirectURI(redirectURI string) string {
	redirectURL, _ := url.Parse(redirectURI)
	authCode := redirectURL.Query().Get("code")

	if authCode == "" {
		log.Fatal("Code parameter not found in the redirection URI!")
	}

	return authCode
}

func GenerateNewToken(authorisationCode string, codeChallenge string) model.Token {
	data := url.Values{}
	data.Set("client_id", CLIENT_ID)
	data.Set("client_secret", CLIENT_SECRET)
	data.Set("code", authorisationCode)
	data.Set("code_verifier", codeChallenge)
	data.Set("grant_type", "authorization_code")
	// data.Set("redirect_uri", REDIRECT_URI)

	response, err := http.PostForm(BaseTokenURL, data)
	if err != nil {
		log.Fatal("Error in fetching authorisation token:", err)
	}

	defer response.Body.Close()

	fmt.Println("Response Status:", response.StatusCode)

	if response.StatusCode != http.StatusOK {
		body, _ := json.MarshalIndent(response.Body, "", "  ")
		log.Fatalf("Error: %s", body)
	}

	var tokenResponse model.Token
	err = json.NewDecoder(response.Body).Decode(&tokenResponse)
	if err != nil {
		log.Fatal("Error in parsing authorisation token:", err)
	}

	fmt.Println("Token generated successfully!")

	fileData, _ := json.MarshalIndent(tokenResponse, "", "")
	err = ioutil.WriteFile("token.json", fileData, 0644)
	if err != nil {
		log.Fatal("Error in writing token file.", err)
	}

	return tokenResponse
}

func PerformMALAuthentication() {
	codeVerifier := GetNewCodeVerifier()
	authURL := GetAuthorisationURL(codeVerifier)

	fmt.Println("Please navigate to the following URL in your browser and authorize the application:", authURL)

	fmt.Print("Enter authorization code: ")
	reader := bufio.NewReader(os.Stdin)
	authCode, _ := reader.ReadString('\n')
	authCode = authCode[:len(authCode)-1]

	token := GenerateNewToken(authCode, codeVerifier)
	PerformGetUserInfo(token.AccessToken)

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	authCode := r.URL.Query().Get("code")
	// 	if authCode == "" {
	// 		log.Fatal("Code parameter missing from the redirect URI!")
	// 		return
	// 	}
	// 	token := GenerateNewToken(authCode, codeVerifier)
	// 	PerformGetUserInfo(token.AccessToken)
	// 	w.Write([]byte("Authorisation completed successfully! You may close this window now."))
	// })
}
