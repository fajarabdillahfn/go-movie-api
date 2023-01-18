package http

import (
	"encoding/json"
	"github.com/fajarabdillahfn/go-movie-api/internal/model"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"

	cWrapper "github.com/fajarabdillahfn/go-movie-api/common/wrapper"
)

func (d *Delivery) AddMovie(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	ctx := r.Context()

	var newMovie model.Movie

	err := json.NewDecoder(r.Body).Decode(&newMovie)
	if err != nil {
		cWrapper.ErrorJSON(w, err, "invalid parameter", http.StatusBadRequest)
		return
	}

	movie, err := d.movieUC.CreateMovie(ctx, &newMovie)
	if err != nil {
		log.Println("ERROR create movie: " + err.Error())
		cWrapper.ErrorJSON(w, err, "failed to add movie", http.StatusInternalServerError)
		return
	}

	err = cWrapper.WriteJSON(w, http.StatusCreated, movie, "")
	if err != nil {
		log.Println("ERROR send response: " + err.Error())
		return
	}
}
