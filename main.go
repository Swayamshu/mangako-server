package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Swayamshu/mangadex/router"
	"github.com/joho/godotenv"
)

// const BaseURL = "https://api.mangadex.org"

func main() {
	fmt.Println("Welcome to MangaDex Test Server!")

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed to load .env file.", err)
	}

	fmt.Println("Server is getting started ...")

	router := router.Router()
	fmt.Println("Listening on Port 4000.\nhttp://localhost:4000")
	log.Fatal(http.ListenAndServe(":4000", router))

	// controller.PerformGetMangaByTitle("One Piece")
	// controller.PerformGetMangaCoverImage("a1c7c817-4e59-43b7-9365-09675a149a6f")
	// controller.PerformGetMangaStatistics("0301208d-258a-444a-8ef7-66e433d801b1")
	// controller.PerformGetChapters("a1c7c817-4e59-43b7-9365-09675a149a6f")
	// controller.PerformGetChapterData("ebc530e1-e89a-4335-865a-4a450bf977dc")
	// str := controller.GetNewCodeVerifier()
	// fmt.Println(str)

	// controller.PerformMALAuthentication()
	// controller.PerformGetMangaRanking("bypopularity", "10")
	// controller.PerformGetMangaDetails("2")
	// controller.PerformGetUserInfo(utils.BearerToken)
}
