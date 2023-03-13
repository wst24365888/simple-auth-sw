package main

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/ServiceWeaver/weaver"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

type RegisterHandler interface {
	Register(ctx context.Context, user UserPayload) error
}

type registerHandler struct {
	weaver.Implements[RegisterHandler]
	weaver.WithConfig[config]

	db *sql.DB
}

func (r *registerHandler) Init(_ context.Context) error {
	cfg := r.Config()
	db, err := sql.Open(
		cfg.DRIVER,
		fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			cfg.HOST,
			cfg.PORT,
			cfg.USER,
			cfg.PASSWORD,
			cfg.DBNAME,
		),
	)

	r.db = db
	return err
}

func (r *registerHandler) Register(ctx context.Context, UserPayload UserPayload) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(UserPayload.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	_, err = r.db.ExecContext(ctx, "INSERT INTO users (username, password) VALUES ($1, $2)", UserPayload.Username, hash)
	return err
}
