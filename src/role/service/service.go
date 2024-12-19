package service

import (
	"database/sql"

	"gitlab.com/wit-id/go-orm-mysql/toolkit/config"
)

type RoleService struct {
	mainDB *sql.DB
	cfg    config.KVStore
}

func NewRoleService(
	mainDB *sql.DB,
	cfg config.KVStore,
) *RoleService {
	return &RoleService{
		mainDB: mainDB,
		cfg:    cfg,
	}
}
