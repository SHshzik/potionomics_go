package handlers

import (
	"github.com/SHshzik/potionomics_go/pkg/logger"
	"github.com/SHshzik/potionomics_go/service"
	"github.com/go-playground/validator/v10"
)

type HTTPServer struct {
	app *service.App
	l   logger.Interface
	v   *validator.Validate
}

func NewHTTPServer(app *service.App, l logger.Interface) *HTTPServer {
	return &HTTPServer{app: app, l: l, v: validator.New()}
}
