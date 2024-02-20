package repository

import (
	"database/sql"

	"github.com/fiaufar/sawit-pro-test/entity"
	"github.com/fiaufar/sawit-pro-test/infrastructure"
)

type UserRepositoryInterface interface {
	GetById(id int64) (*entity.User, error)
	GetByPhoneNumber(phoneNumber string) (*entity.User, error)
	Create(user *entity.User) (int64, error)
	Update(user *entity.User) error
}

type UserRepository struct {
	Db *sql.DB
}

type NewUserRepositoryOptions struct {
	DbConn *infrastructure.DbConnection
}

func NewUserRepository(opts NewUserRepositoryOptions) *UserRepository {
	return &UserRepository{
		Db: opts.DbConn.Db,
	}
}

func (r *UserRepository) GetById(id int64) (*entity.User, error) {
	var user entity.User

	err := r.Db.QueryRow("SELECT id, fullname, phone_number, successful_logged_in FROM \"user\" WHERE id = $1", id).Scan(&user.Id, &user.Fullname, &user.PhoneNumber, &user.SuccessfulLoggedIn)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) GetByPhoneNumber(phoneNumber string) (*entity.User, error) {
	var user entity.User

	err := r.Db.QueryRow("SELECT id, fullname, phone_number, successful_logged_in FROM \"user\" WHERE phone_number = $1", phoneNumber).Scan(&user.Id, &user.Fullname, &user.PhoneNumber, &user.SuccessfulLoggedIn)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) Create(user *entity.User) (int64, error) {
	var insertedId int64
	err := r.Db.QueryRow("INSERT INTO \"user\" (fullname, phone_number) VALUES ($1, $2) RETURNING id", user.Fullname, user.PhoneNumber).Scan(&insertedId)

	if err != nil {
		return 0, err
	}

	return insertedId, nil
}

func (r *UserRepository) Update(user *entity.User) error {
	return r.Db.QueryRow("UPDATE \"user\" SET fullname = $1, phone_number = $2, successful_logged_in = $3 WHERE id = $4", user.Fullname, user.PhoneNumber, user.SuccessfulLoggedIn, user.Id).Err()
}
