package main

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	"github.com/ServiceWeaver/weaver"
	"github.com/golang-jwt/jwt/v5"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

type LoginHandler interface {
	Login(ctx context.Context, user UserPayload) (string, error)
}

type loginHandler struct {
	weaver.Implements[LoginHandler]
	weaver.WithConfig[config]

	db *sql.DB
}

func (l *loginHandler) Init(_ context.Context) error {
	cfg := l.Config()
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

	l.db = db
	return err
}

func (l *loginHandler) Login(ctx context.Context, userPayload UserPayload) (string, error) {
	var dbUser User
	err := l.db.QueryRowContext(ctx, "SELECT id, username, password FROM users WHERE username = $1", userPayload.Username).Scan(&dbUser.ID, &dbUser.Username, &dbUser.Password)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(userPayload.Password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": userPayload.Username,
		"password": userPayload.Password,
	})

	// sign token
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
