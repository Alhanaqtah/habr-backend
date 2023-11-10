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

func (r *repository) GetUser(username string) *user.User {
	const op = "user.repository.GetUser"

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	row := r.storage.QueryRow(ctx, `SELECT username, registration_date, rating, karma, publications, bookmarks, followers, followings FROM users WHERE username=$1`, username)

	var user user.User

	err := row.Scan(
		&user.Username,
		&user.RegistrationDate,
		&user.Rating,
		&user.Karma,
		&user.Publications,
		&user.Bookmarks,
		&user.Followers,
		&user.Followings,
	)
	if err != nil {
		r.log.Error("failed to scan user", slog.String("err", err.Error()), slog.Attr{
			Key:   "op",
			Value: slog.StringValue(op),
		})

		return nil
	}

	return &user

}

func (r *repository) GetUserPublications(username string) *user.User {
	const op = "user.repository.GetUserPublications"

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	row := r.storage.QueryRow(ctx, `SELECT publications FROM users WHERE username=$1`, username)

	var user user.User

	err := row.Scan(
		&user.Publications,
	)
	if err != nil {
		r.log.Error("failed to scan user publications", slog.String("err", err.Error()), slog.Attr{
			Key:   "op",
			Value: slog.StringValue(op),
		})

		return nil
	}

	return &user
}

func (r *repository) GetUserFollowers(username string) *user.User {
	const op = "user.repository.GetUserFollowers"

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	row := r.storage.QueryRow(ctx, `SELECT followers FROM users WHERE username=$1`, username)

	var user user.User

	err := row.Scan(
		&user.Followers,
	)
	if err != nil {
		r.log.Error("failed to scan user followers", slog.String("err", err.Error()), slog.Attr{
			Key:   "op",
			Value: slog.StringValue(op),
		})

		return nil
	}

	return &user
}

func (r *repository) GetUserFollowings(username string) *user.User {
	const op = "user.repository.GetUserFollowings"

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	row := r.storage.QueryRow(ctx, `SELECT followings FROM users WHERE username=$1`, username)

	var user user.User

	err := row.Scan(
		&user.Followers,
	)
	if err != nil {
		r.log.Error("failed to scan user followings", slog.String("err", err.Error()), slog.Attr{
			Key:   "op",
			Value: slog.StringValue(op),
		})

		return nil
	}

	return &user
}
