package http

import (
	"fmt"
	cWrapper "github.com/fajarabdillahfn/go-movie-api/common/wrapper"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"strconv"
)

func (d *Delivery) DeleteMovie(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := r.Context()

	movieId, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		cWrapper.ErrorJSON(w, err, "invalid parameter", http.StatusBadRequest)
		return
	}

	err = d.movieUC.DeleteMovieById(ctx, uint(movieId))
	if err != nil {
		log.Println("ERROR delete a movie: " + err.Error())
		cWrapper.ErrorJSON(w, err, "failed to delete a movie", http.StatusInternalServerError)
		return
	}

	err = cWrapper.WriteJSON(w, http.StatusOK, fmt.Sprintf("movie with id %d deleted", movieId), "status")
	if err != nil {
		log.Println("ERROR send response: " + err.Error())
		return
	}
}
