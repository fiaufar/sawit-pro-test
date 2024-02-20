package repository

import (
	"database/sql"

	"github.com/fiaufar/sawit-pro-test/entity"
	"github.com/fiaufar/sawit-pro-test/infrastructure"
)

type UserCredentialRepositoryInterface interface {
	GetByUserId(userId int64) (*entity.UserCredential, error)
	Create(userCr *entity.UserCredential) error
}

type UserCredentialRepository struct {
	Db *sql.DB
}

type NewUserCredentialRepositoryOptions struct {
	DbConn *infrastructure.DbConnection
}

func NewUserCredentialRepository(opts NewUserCredentialRepositoryOptions) *UserCredentialRepository {
	return &UserCredentialRepository{
		Db: opts.DbConn.Db,
	}
}

func (r *UserCredentialRepository) GetByUserId(userId int64) (*entity.UserCredential, error) {
	var userCr entity.UserCredential

	err := r.Db.QueryRow("SELECT user_id, salt, password FROM user_credential WHERE user_id = $1", userId).Scan(&userCr.UserId, &userCr.Salt, &userCr.Password)
	if err != nil {
		return nil, err
	}

	return &userCr, nil
}

func (r *UserCredentialRepository) Create(userCr *entity.UserCredential) error {
	_, err := r.Db.Exec("INSERT INTO user_credential (user_id, salt, password) VALUES ($1, $2, $3)", userCr.UserId, userCr.Salt, userCr.Password)
	if err != nil {
		return err
	}

	return nil
}
