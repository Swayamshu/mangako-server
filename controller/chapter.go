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

func PerformGetChapters(mangaID string) {
	apiURL, _ := url.Parse(utils.BaseURL)
	apiURL.Path += fmt.Sprintf("/manga/%s/feed", mangaID)

	response, err := http.Get(apiURL.String())
	fmt.Println(apiURL.String())
	if err != nil {
		log.Fatal("Error getting chapter feed.", err)
	}

	defer response.Body.Close()

	var chaptersResponse model.ChapterResponse
	err = json.NewDecoder(response.Body).Decode(&chaptersResponse)

	if err != nil {
		log.Fatal("Error in paresing chapter feed JSON.", err)
	}

	for _, chapter := range chaptersResponse.Data {
		fmt.Printf("Chapter: %s, ID: %s\n", chapter.Attributes.Chapter, chapter.ID)
	}
}

func PerformGetChapterData(chapterID string) {
	apiURL, _ := url.Parse(utils.BaseURL)
	apiURL.Path += fmt.Sprintf("/at-home/server/%s", chapterID)

	fmt.Println(apiURL.String())
	response, err := http.Get(apiURL.String())

	if err != nil {
		log.Fatal("Error in fetching chapter data.", err)
	}

	defer response.Body.Close()

	var chapterData model.ChapterData
	err = json.NewDecoder(response.Body).Decode(&chapterData)

	if err != nil {
		log.Fatal("Error in parsing chapter data.", err)
	}

	var endpoint = chapterData.BaseURL + "/data-saver/" + chapterData.Chapter.Hash
	for i, page := range chapterData.Chapter.DataSaver {
		imageURL := endpoint + "/" + page
		fmt.Printf("Page %d: %s\n", i, imageURL)
	}
}
