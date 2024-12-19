package application

import (
	"math"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"gitlab.com/wit-id/go-orm-mysql/common/httpservice"
	"gitlab.com/wit-id/go-orm-mysql/src/repository/payload"
	"gitlab.com/wit-id/go-orm-mysql/src/role/service"
	"gitlab.com/wit-id/go-orm-mysql/toolkit/config"
	"gitlab.com/wit-id/go-orm-mysql/toolkit/log"
)

func AddRouteRole(s *httpservice.Service, cfg config.KVStore, e *echo.Echo) {
	svc := service.NewRoleService(s.GetDB(), cfg)

	roleApp := e.Group("/role")
	roleApp.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Role app ok")
	})

	roleApp.POST("", createRole(svc))
	roleApp.GET("/detail/:id", getRoleByGuid(svc))
	roleApp.PUT("/:id", updateRole(svc))
	roleApp.POST("/list", listRole(svc))
	roleApp.DELETE("/:id", deleteRole(svc))
}

func createRole(svc *service.RoleService) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var request payload.InsertRolePayload
		if err := ctx.Bind(&request); err != nil {
			log.FromCtx(ctx.Request().Context()).Error(err, "failed to parse request")
			return errors.WithStack(httpservice.ErrBadRequest)
		}

		if err := request.Validate(ctx.Request().Context()); err != nil {
			return err
		}

		Role, err := svc.CreateRole(ctx.Request().Context(), request.ToEntity())
		if err != nil {
			return err
		}

		return httpservice.ResponseData(ctx, payload.ToPayloadRole(Role), nil)
	}
}

func getRoleByGuid(svc *service.RoleService) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		id := ctx.Param("id")
		if id == "" {
			return errors.WithStack(httpservice.ErrBadRequest)
		}

		// Convert the string `idStr` to int64
		idRole, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			return errors.WithStack(httpservice.ErrBadRequest) // Handle conversion error
		}

		Role, err := svc.GetRoleByGUID(ctx.Request().Context(), idRole)
		if err != nil {
			return err
		}

		return httpservice.ResponseData(ctx, payload.ToPayloadRole(Role), nil)
	}
}

func updateRole(svc *service.RoleService) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		id := ctx.Param("id")
		if id == "" {
			return errors.WithStack(httpservice.ErrBadRequest)
		}
		var request payload.UpdateRolePayload
		if err := ctx.Bind(&request); err != nil {
			log.FromCtx(ctx.Request().Context()).Error(err, "failed to parse request")
			return errors.WithStack(httpservice.ErrBadRequest)
		}

		if err := request.Validate(ctx.Request().Context()); err != nil {
			return err
		}

		// Convert the string `idStr` to int64
		idRole, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			return errors.WithStack(httpservice.ErrBadRequest) // Handle conversion error
		}

		Role, err := svc.UpdateRole(ctx.Request().Context(),
			request.ToEntity(idRole))
		if err != nil {
			return err
		}

		return httpservice.ResponseData(ctx, payload.ToPayloadRole(Role), nil)
	}
}

func listRole(svc *service.RoleService) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var request payload.ListRolePayload

		if err := ctx.Bind(&request); err != nil {
			log.FromCtx(ctx.Request().Context()).Error(err, "failed to parse request")
			return errors.WithStack(httpservice.ErrBadRequest)
		}

		// Validate request
		if err := request.Validate(ctx.Request().Context()); err != nil {
			return err
		}

		listData, totalData, err := svc.ListRole(ctx.Request().Context(), request.ToEntity())
		if err != nil {
			return err
		}

		// TOTAL PAGE
		totalPage := math.Ceil(float64(totalData) / float64(request.Limit))

		return httpservice.ResponsePagination(ctx,
			payload.ToPayloadListRole(listData),
			nil, int(request.Offset),
			int(request.Limit),
			int(totalPage),
			int(totalData))
	}
}

func deleteRole(svc *service.RoleService) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		id := ctx.Param("id")
		if id == "" {
			return errors.WithStack(httpservice.ErrBadRequest)
		}

		// Convert the string `idStr` to int64
		idRole, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			return errors.WithStack(httpservice.ErrBadRequest) // Handle conversion error
		}

		err = svc.DeleteRole(ctx.Request().Context(), idRole)
		if err != nil {
			return err
		}

		return httpservice.ResponseData(ctx, nil, nil)
	}
}
