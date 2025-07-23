package dto

import (
	"strings"

	"github.com/GustavoZeglan/Cine/pkg/helper"
)

type CreateMovieInput struct {
	Title string `json:"title"`
}

func (i *CreateMovieInput) Validate() error {
	errors := make(map[string]string)

	if strings.TrimSpace(i.Title) == "" {
		errors["title"] = "title is required"
	}

	if len(errors) > 0 {
		return helper.ValidationError{Errors: errors}
	}
	return nil
}
