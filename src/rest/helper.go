package rest

import (
	"github.com/go-chi/render"
	"net/http"
)

type ErrResponse struct {
	Err            error
	HTTPStatusCode int
	StatusText     string
	ErrorText      string
}

func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}
func ErrInvalidRequest(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: 400,
		StatusText:     "Invalid request.",
		ErrorText:      err.Error(),
	}
}
func ErrUnauthorized(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: http.StatusUnauthorized,
		StatusText:     "Unauthozed",
		ErrorText:      err.Error(),
	}
}

func ErrNotFound(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: http.StatusNotFound,
		StatusText:     "Not Found",
		ErrorText:      err.Error(),
	}
}

func ErrNotAlreadyRegisterd(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: http.StatusConflict,
		StatusText:     "Already Registered",
		ErrorText:      err.Error(),
	}
}
