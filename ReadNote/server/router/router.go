package router

import (
	"github.com/Chief-Rishab/mymodule/controller"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/articles", controller.GetAllArticles).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/article", controller.CreateArticle).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/article/{id}", controller.UpdateArticle).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/articleUnread/{id}", controller.UpdateArticleUnread).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/delArticle/{id}", controller.DeleteArticle).Methods("DELETE", "OPTIONS")
	router.HandleFunc("/api/deletearticles", controller.DeleteAllArticles).Methods("DELETE", "OPTIONS")
	

	return router
}
