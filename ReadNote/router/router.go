package router

import (
	"github.com/Chief-Rishab/mymodule/controller"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/articles", controller.GetAllArticles).Methods("GET")
	router.HandleFunc("/api/article", controller.CreateArticle).Methods("POST")
	router.HandleFunc("/api/article/{id}", controller.UpdateArticle).Methods("PUT")
	router.HandleFunc("/api/article/{id}", controller.DeleteArticle).Methods("DELETE")
	router.HandleFunc("/api/deletearticles", controller.DeleteAllArticles).Methods("DELETE")

	return router
}
