package delete

import (
	"fmt"
	"net/http"

	"github.com/arynskiii/url-shortener/internal/lib/api/response"
	"github.com/arynskiii/url-shortener/internal/lib/logger/sl"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"

	"golang.org/x/exp/slog"
)

const aliasLength = 6

type Request struct {
	URL   string `json:"url" validate:"required,url"`
	Alias string `json:"alias,omitempty"`
}

//go:generate go run  github.com/vektra/mockery/v2@v2.20.2 --name=URLSaver

type Response struct {
	Alias string `json:"alias,omitempty"`
}
type DeleteUrl interface {
	DeleteURL(alias string) error
}

func New(log *slog.Logger, DeleteUrl DeleteUrl) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.url.save.New"
		log = log.With(slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)
		var req Request
		err := render.DecodeJSON(r.Body, &req)
		if err != nil {
			log.Error("failed to decode request body", sl.Err(err))
			render.JSON(w, r, response.Error("failed to decode request"))
			return
		}
		log.Info("request body decoded", slog.Any("request", req))

		fmt.Println("---------------------------------------")
		alias := req.Alias

		fmt.Println(alias)
		err = DeleteUrl.DeleteURL(alias)

		if err != nil {
			log.Error("failed to add url", sl.Err(err))
			render.JSON(w, r, response.Error("failed to add url"))
			return
		}

	}
}
