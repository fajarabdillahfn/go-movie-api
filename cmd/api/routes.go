package api

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func routes(router *httprouter.Router) {
	router.GET("/", func(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
		_, err := w.Write([]byte("Welcome to Movie API"))
		if err != nil {
			return
		}
	})

	router.POST("/movies", MovieHTTPDelivery.AddMovie)

	router.GET("/movies", MovieHTTPDelivery.GetAllMovies)
	router.GET("/movies/:id", MovieHTTPDelivery.GetMovieById)

	router.PATCH("/movies/:id", MovieHTTPDelivery.UpdateMovieById)

	router.DELETE("/movies/:id", MovieHTTPDelivery.DeleteMovie)
}
