package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/Swayamshu/mangadex/model"
	"github.com/Swayamshu/mangadex/utils"
)

func PerformGetUserInfo(accessToken string) {
	apiURL, _ := url.Parse(utils.MALBaseURL)
	apiURL.Path += "/users/@me"

	params := url.Values{}
	params.Add("fields", utils.MALUserDetailsField)
	apiURL.RawQuery = params.Encode()

	request, err := http.NewRequest("GET", apiURL.String(), nil)
	if err != nil {
		log.Fatal("Error creating request", err)
	}

	request.Header.Set("Authorization", "Bearer "+accessToken)

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		log.Fatal("Error in performing GET request.", err)
	}

	defer response.Body.Close()

	var userResponse model.User
	err = json.NewDecoder(response.Body).Decode(&userResponse)
	if err != nil {
		log.Fatal("Error in parsing user response JSON.", err)
	}

	fmt.Printf("You are successfully authenticated, %s!\n", userResponse.Name)
	fmt.Println(userResponse)
}
