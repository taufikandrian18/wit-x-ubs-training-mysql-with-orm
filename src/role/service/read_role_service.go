package service

import (
	"context"

	"github.com/pkg/errors"
	"gitlab.com/wit-id/go-orm-mysql/common/httpservice"
	sqlc "gitlab.com/wit-id/go-orm-mysql/src/repository/pgbo_sqlc"
	"gitlab.com/wit-id/go-orm-mysql/toolkit/log"
)

func (s *RoleService) GetRoleByGUID(ctx context.Context, RoleID int64) (
	role sqlc.Role, err error) {
	query := sqlc.New(s.mainDB)

	role, err = query.GetRole(ctx, RoleID)
	if err != nil {
		log.FromCtx(ctx).Error(err, "failed get Role")
		err = errors.WithStack(httpservice.ErrDataNotFound)

		return
	}

	return
}

func (s *RoleService) GetLastInsertedID(ctx context.Context) (
	roleID int64, err error) {
	query := sqlc.New(s.mainDB)

	roleID, err = query.GetLatsInsertedID(ctx)
	if err != nil {
		log.FromCtx(ctx).Error(err, "failed get last inserted Role")
		err = errors.WithStack(httpservice.ErrDataNotFound)

		return
	}

	return
}

func (s *RoleService) ListRole(ctx context.Context, request sqlc.ListRolesParams) (
	listRole []sqlc.Role, totalData int64, err error) {
	query := sqlc.New(s.mainDB)

	// Get Total Data
	totalData, err = s.getTotalDataRole(ctx, query, request)
	if err != nil {
		return
	}

	listRole, err = query.ListRoles(ctx, request)
	if err != nil {
		log.FromCtx(ctx).Error(err, "failed read list Role")
		err = errors.WithStack(httpservice.ErrUnknownSource)

		return
	}

	return
}

func (s *RoleService) getTotalDataRole(
	ctx context.Context,
	query *sqlc.Queries,
	request sqlc.ListRolesParams) (totalData int64, err error) {
	requestParam := sqlc.CountRoleParams{
		Column1:  request.Column1,
		RoleID:   request.RoleID,
		Column3:  request.Column3,
		Column4:  request.Column4,
		RoleName: request.RoleName,
		Column6:  request.Column6,
	}

	totalData, err = query.CountRole(ctx, requestParam)
	if err != nil {
		log.FromCtx(ctx).Error(err, "Failed get total data Role")
		err = errors.WithStack(httpservice.ErrUnknownSource)

		return
	}

	return
}
