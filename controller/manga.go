package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/Swayamshu/mangadex/model"
	"github.com/Swayamshu/mangadex/utils"
	"github.com/gorilla/mux"
)

func getMangaByTitle(title string) model.MangaList {
	apiURL, _ := url.Parse(utils.BaseURL)
	apiURL.Path += "/manga"

	// adding GET parameter for request
	params := url.Values{}
	params.Add("title", title)
	apiURL.RawQuery = params.Encode()

	response, err := http.Get(apiURL.String())
	if err != nil {
		log.Fatal("Error in performing GET request.", err)
		return model.MangaList{}
	}

	defer response.Body.Close()

	var mangaResponse model.MangaList
	err = json.NewDecoder(response.Body).Decode(&mangaResponse)

	if err != nil {
		log.Fatal("Error in parsing JSON response.", err)
		return model.MangaList{}
	}

	for _, manga := range mangaResponse.Data {
		fmt.Println(manga.Attributes.Title.En + ": " + manga.ID)
	}

	return mangaResponse
}

func PerformGetMangaByTitle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "GET")

	mangaTitle := r.FormValue("title")
	if mangaTitle == "" {
		http.Error(w, "Title parameter is missing", http.StatusBadRequest)
		return
	}

	mangaList := getMangaByTitle(mangaTitle)
	json.NewEncoder(w).Encode(mangaList)
}

func GetMangaCoverImage(mangaID string) string {
	// Requirements:
	// 1. id of cover art
	// 2. file name of cover art

	apiURL, _ := url.Parse(utils.BaseURL)
	apiURL.Path += "/manga/" + mangaID

	response, err := http.Get(apiURL.String())
	if err != nil {
		log.Fatal("Error in performing GET request.", err)
		return ""
	}

	defer response.Body.Close()

	var mangaResponse model.MangaResponse
	err = json.NewDecoder(response.Body).Decode(&mangaResponse)

	if err != nil {
		log.Fatal("Error in parsing JSON response.", err)
		return ""
	}

	var coverImageID string
	for _, mangaRelationship := range mangaResponse.Data.Relationships {
		if mangaRelationship.Type == "cover_art" {
			coverImageID = mangaRelationship.ID
			break
		}
	}

	fileURL, _ := url.Parse(utils.BaseURL)
	fileURL.Path += "/cover/" + coverImageID

	res, err := http.Get(fileURL.String())
	if err != nil {
		log.Fatal("Error in getting file name for cover image.", err)
		return ""
	}

	defer res.Body.Close()

	var coverArtResponse model.CoverArtResponse
	err = json.NewDecoder(res.Body).Decode(&coverArtResponse)

	if err != nil {
		log.Fatal("Error in parsing JSON response.", err)
		return ""
	}

	var fileName string = coverArtResponse.Data.Attributes.FileName

	coverArtURL, _ := url.Parse(utils.CoverURL)
	coverArtURL.Path += fmt.Sprintf("/%s/%s", mangaID, fileName)

	return coverArtURL.String()
}

func getMangaDetails(mangaID string) model.MangaDetails {
	apiURL, _ := url.Parse(utils.MALBaseURL)
	apiURL.Path += fmt.Sprintf("/manga/%s", mangaID)

	params := url.Values{}
	params.Add("fields", utils.MALMangaDetailsField)
	apiURL.RawQuery = params.Encode()

	request, err := http.NewRequest("GET", apiURL.String(), nil)
	if err != nil {
		log.Fatal("Error creating request for manga details.", err)
	}

	request.Header.Set("X-MAL-CLIENT-ID", CLIENT_ID)

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		log.Fatal("Error fetching manga details response.", err)
	}

	defer response.Body.Close()

	var mangaDetailsResponse model.MangaDetails
	err = json.NewDecoder(response.Body).Decode(&mangaDetailsResponse)
	if err != nil {
		log.Fatal("Error parsing manga details JSON.", err)
	}

	// use mangadex id here
	// id := "#"
	// mangaDetailsResponse.CoverImage = getMangaCoverImage(id)

	return mangaDetailsResponse
}

func PerformGetMangaDetails(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "GET")

	params := mux.Vars(r)
	mangaID := params["id"]
	manga := getMangaDetails(mangaID)
	json.NewEncoder(w).Encode(manga)
}

func getMangaRanking(filter string, limit string) model.TrendingMangaList {
	apiURL, _ := url.Parse(utils.MALBaseURL)
	apiURL.Path += "/manga/ranking"

	fmt.Println("fetching manga ranking, type:", filter)

	// correct way to ensure proper encoding of params in api url
	params := url.Values{}
	params.Add("ranking_type", filter)
	params.Add("limit", limit)
	apiURL.RawQuery = params.Encode()

	request, err := http.NewRequest("GET", apiURL.String(), nil)
	if err != nil {
		log.Fatal("Error creating request for manga ranking", err)
	}

	// request.Header.Set("Authorization", "Bearer "+utils.BearerToken)
	// we don't need Bearer Token when user login is not required
	request.Header.Set("X-MAL-CLIENT-ID", CLIENT_ID)

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		log.Fatal("Error fetching manga rankings.", err)
	}

	defer response.Body.Close()

	var trendingMangaResponse model.TrendingMangaList
	err = json.NewDecoder(response.Body).Decode(&trendingMangaResponse)
	if err != nil {
		log.Fatal("Error parsing trending manga JSON.", err)
	}
	// for _, manga := range trendingMangaResponse.Data {
	// 	fmt.Println("Rank:", manga.Ranking.Rank)
	// 	fmt.Printf("Manga: %s\n\n", manga.Node.Title)
	// }

	return trendingMangaResponse
}

func PerformGetTrendingManga(w http.ResponseWriter, r *http.Request) {
	fmt.Println("fetching trending manga ...")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "GET")

	mangaList := getMangaRanking("manga", "10")
	json.NewEncoder(w).Encode(mangaList)
}

func PerformGetPopularManga(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "GET")

	mangaList := getMangaRanking("bypopularity", "10")
	json.NewEncoder(w).Encode(mangaList)
}

func PerformGetMostFavoritedManga(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "GET")

	mangaList := getMangaRanking("favorite", "10")
	json.NewEncoder(w).Encode(mangaList)
}
