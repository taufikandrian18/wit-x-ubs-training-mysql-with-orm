package service

import (
	"context"
	"database/sql"

	"github.com/pkg/errors"
	"gitlab.com/wit-id/go-orm-mysql/common/constants"
	"gitlab.com/wit-id/go-orm-mysql/common/httpservice"
	"gitlab.com/wit-id/go-orm-mysql/common/utility"
	sqlc "gitlab.com/wit-id/go-orm-mysql/src/repository/pgbo_sqlc"
	"gitlab.com/wit-id/go-orm-mysql/toolkit/log"
)

func (s *RoleService) DeleteRole(ctx context.Context, id int64) (err error) {
	data := sqlc.UpdateRoleStatusParams{
		Status: constants.StatusDeleted,
		UpdatedBy: sql.NullString{
			String: constants.CreatedByTemporaryBySystem,
			Valid:  true,
		},
		RoleID: id,
	}
	_, err = utility.Transaction(ctx, s.mainDB, func(query *sqlc.Queries) (response interface{}, err error) {
		err = query.UpdateRoleStatus(ctx, data)

		return
	})
	if err != nil {
		log.FromCtx(ctx).Error(err, "failed delete Role")
		err = errors.WithStack(httpservice.ErrUnknownSource)

		return
	}

	return
}
