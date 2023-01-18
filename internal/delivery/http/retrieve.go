package http

import (
	cWrapper "github.com/fajarabdillahfn/go-movie-api/common/wrapper"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func (d *Delivery) GetAllMovies(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	ctx := r.Context()

	movies, err := d.movieUC.GetAllMovies(ctx)
	if err != nil {
		log.Println("ERROR get all movies: " + err.Error())
		cWrapper.ErrorJSON(w, err, "failed to get all movies", http.StatusInternalServerError)
		return
	}

	err = cWrapper.WriteJSON(w, http.StatusOK, movies, "movies")
	if err != nil {
		log.Println("ERROR send response: " + err.Error())
		return
	}
}
