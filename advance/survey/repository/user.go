package repository

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type userRepo struct {
	DB *sqlx.DB
}

type User struct {
	UserID    int64  `json:"user_id" db:"user_id"`
	Name      string `json:"name" db:"name"`
	HairColor string `json:"hair_color" db:"hair_color"`
	Age       int64  `json:"age" db:"age"`
}

func NewUserRepository(db *sqlx.DB) UserRepository {
	return userRepo{
		DB: db,
	}
}

func (u userRepo) InsertUserData(ctx context.Context, form User) error {
	_, err := u.DB.ExecContext(ctx, "insert into users(name, hair_color, age) values($1, $2, $3)", form.Name, form.HairColor, form.Age)
	if err != nil {
		return err
	}

	return nil
}

func (u userRepo) GetUserData(ctx context.Context) (res []User, err error) {
	rows, err := u.DB.QueryxContext(ctx, "select user_id, name, hair_color, age from users")
	if err != nil {
		return res, err
	}

	defer rows.Close()

	for rows.Next() {
		tmp := User{}
		err = rows.StructScan(&tmp)
		if err != nil {
			return res, err
		}

		res = append(res, tmp)
	}

	return res, err
}

type UserRepository interface {
	InsertUserData(ctx context.Context, form User) error
	GetUserData(ctx context.Context) (res []User, err error)
}
