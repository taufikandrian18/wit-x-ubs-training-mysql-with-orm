package service

import (
	"context"

	"github.com/pkg/errors"
	"gitlab.com/wit-id/go-orm-mysql/common/httpservice"
	"gitlab.com/wit-id/go-orm-mysql/common/utility"
	sqlc "gitlab.com/wit-id/go-orm-mysql/src/repository/pgbo_sqlc"
	"gitlab.com/wit-id/go-orm-mysql/toolkit/log"
)

func (s *RoleService) UpdateRole(ctx context.Context, data sqlc.UpdateRoleParams) (
	Role sqlc.Role, err error) {

	_, err = utility.Transaction(ctx, s.mainDB, func(query *sqlc.Queries) (response interface{}, err error) {
		err = query.UpdateRole(ctx, data)
		return
	})
	if err != nil {
		log.FromCtx(ctx).Error(err, "failed update Role")
		err = errors.WithStack(httpservice.ErrUnknownSource)

		return
	}

	// Get relation data post
	Role, err = s.GetRoleByGUID(ctx, data.RoleID)
	if err != nil {
		return
	}

	return
}
