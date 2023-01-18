package http

import (
	"encoding/json"
	cWrapper "github.com/fajarabdillahfn/go-movie-api/common/wrapper"
	"github.com/fajarabdillahfn/go-movie-api/internal/model"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"strconv"
)

func (d *Delivery) UpdateMovieById(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := r.Context()

	movieId, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		cWrapper.ErrorJSON(w, err, "invalid parameter", http.StatusBadRequest)
		return
	}

	_, err = d.movieUC.GetMovieById(ctx, uint(movieId))
	if err != nil {
		log.Println("ERROR update a movie: " + err.Error())
		cWrapper.ErrorJSON(w, err, "movie is not valid", http.StatusBadRequest)
		return
	}

	var updateMovie model.InputMovie

	err = json.NewDecoder(r.Body).Decode(&updateMovie)
	if err != nil {
		cWrapper.ErrorJSON(w, err, "invalid parameter", http.StatusBadRequest)
		return
	}

	err = d.movieUC.UpdateMovie(ctx, uint(movieId), &updateMovie)
	if err != nil {
		log.Println("ERROR update a movie: " + err.Error())
		cWrapper.ErrorJSON(w, err, "failed to update movie", http.StatusInternalServerError)
		return
	}

	movie, err := d.movieUC.GetMovieById(ctx, uint(movieId))
	if err != nil {
		log.Println("ERROR get a movie after updated: " + err.Error())
		cWrapper.ErrorJSON(w, err, "failed to validate updated movie", http.StatusInternalServerError)
		return
	}

	err = cWrapper.WriteJSON(w, http.StatusOK, movie, "")
	if err != nil {
		log.Println("ERROR send response: " + err.Error())
		return
	}
}
