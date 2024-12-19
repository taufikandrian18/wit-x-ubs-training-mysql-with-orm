package service

import (
	"context"

	"github.com/pkg/errors"
	"gitlab.com/wit-id/go-orm-mysql/common/httpservice"
	"gitlab.com/wit-id/go-orm-mysql/common/utility"
	sqlc "gitlab.com/wit-id/go-orm-mysql/src/repository/pgbo_sqlc"
	"gitlab.com/wit-id/go-orm-mysql/toolkit/log"
)

func (s *RoleService) CreateRole(
	ctx context.Context,
	params sqlc.InsertRoleParams) (
	Role sqlc.Role,
	err error) {

	_, err = utility.Transaction(ctx, s.mainDB, func(query *sqlc.Queries) (response interface{}, err error) {
		err = query.InsertRole(ctx, params)
		return
	})
	if err != nil {
		log.FromCtx(ctx).Error(err, "failed create Role")
		err = errors.WithStack(httpservice.ErrUnknownSource)

		return
	}

	roleID, errGetRole := s.GetLastInsertedID(ctx)
	if errGetRole != nil {
		err = errGetRole
		return
	}

	// Get relation data post
	Role, err = s.GetRoleByGUID(ctx, roleID)
	if err != nil {
		return
	}

	return
}
