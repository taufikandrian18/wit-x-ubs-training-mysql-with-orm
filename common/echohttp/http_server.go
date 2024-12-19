package echohttp

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"gitlab.com/wit-id/go-orm-mysql/common/constants"
	"gitlab.com/wit-id/go-orm-mysql/common/httpservice"
	roleApp "gitlab.com/wit-id/go-orm-mysql/src/role/application"
	"gitlab.com/wit-id/go-orm-mysql/toolkit/config"
	"gitlab.com/wit-id/go-orm-mysql/toolkit/echokit"
)

func RunEchoHTTPService(ctx context.Context, s *httpservice.Service, cfg config.KVStore) {
	e := echo.New()
	e.HTTPErrorHandler = handleEchoError(cfg)
	e.Use(handleLanguage)

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowCredentials: true,
		AllowMethods:     []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete, http.MethodOptions},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, constants.DefaultAllowHeaderToken, constants.DefaultAllowHeaderRefreshToken, constants.DefaultAllowHeaderAuthorization},
	}))

	runtimeCfg := echokit.NewRuntimeConfig(cfg, "restapi")
	runtimeCfg.HealthCheckFunc = s.GetServiceHealth

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "ok")
	})

	roleApp.AddRouteRole(s, cfg, e)

	// set route config
	httpservice.SetRouteConfig(ctx, s, cfg, e)

	// run actual server
	echokit.RunServerWithContext(ctx, e, runtimeCfg)
}
