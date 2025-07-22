package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/GustavoZeglan/Cine/internal/usecase"
	"github.com/GustavoZeglan/Cine/pkg/contracts"
)

type MovieHandler struct {
	GetMovies   *usecase.GetMovies
	CreateMovie *usecase.CreateMovie
}

type IMovieHandler interface {
	GetAll(ctx *contracts.HttpContext)
	Create(ctx *contracts.HttpContext)
}

func NewMovieHandler(getMovies *usecase.GetMovies, createMovie *usecase.CreateMovie) *MovieHandler {
	return &MovieHandler{GetMovies: getMovies, CreateMovie: createMovie}
}

var _ IMovieHandler = (*MovieHandler)(nil)

func (mh *MovieHandler) GetAll(ctx *contracts.HttpContext) {
	movies, err := mh.GetMovies.Execute(ctx.Context)
	if err != nil {
		ctx.Writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(ctx.Writer).Encode("Failed to fetch movies")
	}
	ctx.Writer.WriteHeader(http.StatusOK)
	json.NewEncoder(ctx.Writer).Encode(movies)
}

func (mh *MovieHandler) Create(ctx *contracts.HttpContext) {
	var input struct {
		Title string `json:"title"`
	}
	if err := json.NewDecoder(ctx.Request.Body).Decode(&input); err != nil {
		ctx.Writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(ctx.Writer).Encode("Failed to parse json")
		return
	}
	output, err := mh.CreateMovie.Execute(ctx.Request.Context(), input)
	if err != nil {
		ctx.Writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(ctx.Writer).Encode("Failed to create movie")
		return
	}
	ctx.Writer.WriteHeader(http.StatusCreated)
	json.NewEncoder(ctx.Writer).Encode(output)
}
