package handlers

import (
	"context"
	"errors"
	"io"
	"log"
	"net/http"

	"github.com/natalliakoita/gocourse/lesson_10/datasource"
)

type ContactHandler struct {
	ds  *datasource.DS
	ctx context.Context
}

func NewContactHandler(d *datasource.DS, ctx context.Context) ContactHandler {
	cont := ContactHandler{
		ds:  d,
		ctx: ctx,
	}
	return cont
}

func ParseBodyRequest(r *http.Request) ([]byte, error) {
	body := r.Body
	defer body.Close()
	bodyData, err := io.ReadAll(body)
	if err != nil {
		log.Print(err.Error())
		return nil, errors.New("failed to read request body")
	}
	return bodyData, nil
}
