package services

import "database/sql"

type UserService struct {
	db *sql.DB
}

func NewUserService() *UserService {
	return &UserService{}
}
