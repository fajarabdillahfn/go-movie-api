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
}
