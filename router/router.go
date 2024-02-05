package router

import (
	"github.com/Swayamshu/mangadex/controller"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/manga", controller.PerformGetMangaByTitle).Methods("GET")
	router.HandleFunc("/manga/trending", controller.PerformGetTrendingManga).Methods("GET")
	router.HandleFunc("/manga/popular", controller.PerformGetPopularManga).Methods("GET")
	router.HandleFunc("/manga/favorite", controller.PerformGetMostFavoritedManga).Methods("GET")
	router.HandleFunc("/manga/{id}", controller.PerformGetMangaDetails).Methods("GET")

	return router
}
