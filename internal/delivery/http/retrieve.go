package http

import (
	cWrapper "github.com/fajarabdillahfn/go-movie-api/common/wrapper"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"strconv"
	"strings"
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

func (d *Delivery) GetMovieById(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := r.Context()

	movieId, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		cWrapper.ErrorJSON(w, err, "invalid parameter", http.StatusBadRequest)
		return
	}

	movie, err := d.movieUC.GetMovieById(ctx, uint(movieId))
	if err != nil {
		log.Println("ERROR get a movie: " + err.Error())
		if strings.Contains(err.Error(), "not found") {
			cWrapper.ErrorJSON(w, err, "movie not found", http.StatusNotFound)
		} else {
			cWrapper.ErrorJSON(w, err, "failed to get a movie", http.StatusInternalServerError)
		}
		return
	}

	err = cWrapper.WriteJSON(w, http.StatusOK, movie, "")
	if err != nil {
		log.Println("ERROR send response: " + err.Error())
		return
	}
}
