package payload

import (
	"context"
	"database/sql"
	"time"

	"gitlab.com/wit-id/go-orm-mysql/common/constants"
	"gitlab.com/wit-id/go-orm-mysql/common/utility"
	sqlc "gitlab.com/wit-id/go-orm-mysql/src/repository/pgbo_sqlc"
)

type InsertRolePayload struct {
	RoleName string `json:"role_name" valid:"required"`
}

type UpdateRolePayload struct {
	RoleName string `json:"role_name" valid:"required"`
}

type ListRolePayload struct {
	Filter ListFilterRolePayload `json:"filter"`
	Limit  int32                 `json:"limit" valid:"required~limit is required field"`
	Offset int32                 `json:"page" valid:"required~page is required field"`
	Order  string                `json:"order" valid:"required~order is required field"`
	Sort   string                `json:"sort" valid:"required~sort is required field"` // ASC, DESC
}

type ListFilterRolePayload struct {
	SetRoleID   bool   `json:"set_role_id"`
	RoleID      int64  `json:"role_id"`
	SetRoleName bool   `json:"set_role_name"`
	RoleName    string `json:"role_name"`
}

type readRolePayload struct {
	RoleID    int64      `json:"role_id,omitempty"`
	RoleName  string     `json:"role_name"`
	CreatedAt time.Time  `json:"created_at"`
	CreatedBy string     `json:"created_by"`
	UpdatedAt *time.Time `json:"updated_at"`
	UpdatedBy *string    `json:"updated_by"`
}

func (payload *InsertRolePayload) Validate(ctx context.Context) (err error) {
	if err = utility.ValidateStruct(ctx, payload); err != nil {
		return
	}

	return
}

func (payload *UpdateRolePayload) Validate(ctx context.Context) (err error) {
	if err = utility.ValidateStruct(ctx, payload); err != nil {
		return
	}

	return
}

func (payload *ListRolePayload) Validate(ctx context.Context) (err error) {
	if err = utility.ValidateStruct(ctx, payload); err != nil {
		return
	}

	return
}

func (payload *InsertRolePayload) ToEntity() (data sqlc.InsertRoleParams) {
	data = sqlc.InsertRoleParams{
		RoleName:  payload.RoleName,
		Status:    constants.StatusActive,
		CreatedBy: constants.CreatedByTemporaryBySystem,
	}

	return
}

func (payload *UpdateRolePayload) ToEntity(key int64) (data sqlc.UpdateRoleParams) {
	data = sqlc.UpdateRoleParams{
		RoleName:  payload.RoleName,
		Status:    constants.StatusActive,
		UpdatedBy: sql.NullString{String: constants.CreatedByTemporaryBySystem, Valid: true},
		RoleID:    key,
	}

	return
}

func (payload *ListRolePayload) ToEntity() (data sqlc.ListRolesParams) {
	data = sqlc.ListRolesParams{
		Column1:  payload.Filter.SetRoleID,
		RoleID:   payload.Filter.RoleID,
		Column3:  payload.Filter.SetRoleID,
		Column4:  payload.Filter.SetRoleName,
		RoleName: payload.Filter.RoleName,
		Column6:  payload.Filter.SetRoleName,
		Column7:  makeOrderParam(payload.Order, payload.Sort),
		Column8:  makeOrderParam(payload.Order, payload.Sort),
		Column9:  makeOrderParam(payload.Order, payload.Sort),
		Column10: makeOrderParam(payload.Order, payload.Sort),
		Offset:   makeOffset(payload.Limit, payload.Offset),
		Limit:    limitWithDefault(payload.Limit),
	}

	return
}

func ToPayloadRole(data sqlc.Role) (payload readRolePayload) {
	payload = readRolePayload{
		RoleID:    data.RoleID,
		RoleName:  data.RoleName,
		CreatedAt: data.CreatedAt.Time,
		CreatedBy: data.CreatedBy,
	}

	if data.UpdatedAt.Valid {
		payload.UpdatedAt = &data.UpdatedAt.Time
	}

	if data.UpdatedBy.Valid {
		payload.UpdatedBy = &data.UpdatedBy.String
	}

	return
}

func ToPayloadGetRole(data sqlc.Role) (payload readRolePayload) {
	payload = readRolePayload{
		RoleID:    data.RoleID,
		RoleName:  data.RoleName,
		CreatedAt: data.CreatedAt.Time,
		CreatedBy: data.CreatedBy,
	}

	if data.UpdatedAt.Valid {
		payload.UpdatedAt = &data.UpdatedAt.Time
	}

	if data.UpdatedBy.Valid {
		payload.UpdatedBy = &data.UpdatedBy.String
	}

	return
}

func ToPayloadListRole(collection []sqlc.Role) (payload []*readRolePayload) {
	payload = make([]*readRolePayload, len(collection))

	for i := range collection {
		payload[i] = new(readRolePayload)
		data := readRolePayload{
			RoleID:    collection[i].RoleID,
			RoleName:  collection[i].RoleName,
			CreatedAt: collection[i].CreatedAt.Time,
			CreatedBy: collection[i].CreatedBy,
		}

		if collection[i].UpdatedAt.Valid {
			data.UpdatedAt = &collection[i].UpdatedAt.Time
		}

		if collection[i].UpdatedBy.Valid {
			data.UpdatedBy = &collection[i].UpdatedBy.String
		}

		payload[i] = &data
	}

	return
}
