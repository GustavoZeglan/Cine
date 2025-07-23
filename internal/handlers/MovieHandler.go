package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/GustavoZeglan/Cine/internal/dto"
	"github.com/GustavoZeglan/Cine/internal/usecase"
	adapter "github.com/GustavoZeglan/Cine/pkg/adapter/handler"
	"github.com/GustavoZeglan/Cine/pkg/helper"
)

type MovieHandler struct {
	GetMovies   *usecase.GetMovies
	CreateMovie *usecase.CreateMovie
}

type IMovieHandler interface {
	GetAll(ctx *adapter.HttpContext) error
	Create(ctx *adapter.HttpContext) error
}

func NewMovieHandler(getMovies *usecase.GetMovies, createMovie *usecase.CreateMovie) *MovieHandler {
	return &MovieHandler{GetMovies: getMovies, CreateMovie: createMovie}
}

var _ IMovieHandler = (*MovieHandler)(nil)
var ve helper.ValidationError

func (mh *MovieHandler) GetAll(ctx *adapter.HttpContext) error {
	movies, err := mh.GetMovies.Execute(ctx.Context)
	if err != nil {
		return err
	}
	return helper.WriteJSON(ctx.Writer, http.StatusOK, movies)
}

func (mh *MovieHandler) Create(ctx *adapter.HttpContext) error {
	var input dto.CreateMovieInput
	if err := json.NewDecoder(ctx.Request.Body).Decode(&input); err != nil {
		return helper.InvalidJSON()
	}
	if err := input.Validate(); err != nil {
		if errors.As(err, &ve) {
			return helper.WriteJSON(ctx.Writer, http.StatusBadRequest, ve.Errors)
		}
	}
	defer ctx.Request.Body.Close()
	output, err := mh.CreateMovie.Execute(ctx.Request.Context(), input)
	if err != nil {
		return err
	}
	return helper.WriteJSON(ctx.Writer, http.StatusCreated, output)
}
