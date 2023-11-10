package repository

import (
	"context"
	"log/slog"

	"Alhanaqtah/habr-backend/internal/user"
	"github.com/jackc/pgx/v4/pgxpool"
)

type repository struct {
	storage *pgxpool.Pool
	log     *slog.Logger
}

func New(pgClient *pgxpool.Pool, log *slog.Logger) *repository {
	return &repository{
		storage: pgClient,
		log:     log,
	}
}

func (r *repository) GetUsers() *[]user.User {
	const op = "user.repository.GetUsers"

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	rows, err := r.storage.Query(ctx, `SELECT username, rating, karma, registration_date FROM users`)
	if err != nil {
		r.log.Error("failed to get all users", slog.String("err", err.Error()), slog.Attr{
			Key:   "op",
			Value: slog.StringValue(op),
		})

		return nil
	}

	var users []user.User

	for rows.Next() {
		var user user.User

		rows.Scan(
			&user.Username,
			&user.Rating,
			&user.Karma,
			&user.RegistrationDate,
		)

		users = append(users, user)
	}

	return &users

}
