package api

import (
	"github.com/nenecchuu/arcana/env"
	"github.com/nenecchuu/lizbeth-be-core/init/service"

	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	_ "github.com/nenecchuu/lizbeth-be-core/docs" // swagger docs
)

type Options struct {
	DefaultTimeout int
	Service        *service.RestService
}

type API struct {
	options *Options
	service *service.RestService
}

func New(o *Options) *API {
	return &API{
		options: o,
		service: o.Service,
	}
}

func (a *API) Register(srv *fiber.App, hdl *service.RestHandlers, mw *service.Middlewares) {
	a.initSwagger(srv)
	a.registerAuthAPI(srv, hdl, mw)
}

func (a *API) initSwagger(srv *fiber.App) {
	if a.service.Config.Rest.EnableSwagger {
		if env.GetEnvironmentName() != "production" {
			srv.Use("swagger", swagger.Handler)
		}
	}
}

func (a *API) registerAuthAPI(srv *fiber.App, hdl *service.RestHandlers, mw *service.Middlewares) {
	router := srv.Group("/auth")

	router.Get("/callback", hdl.AuthRestHandler.HandleLinkageCallback)
}
