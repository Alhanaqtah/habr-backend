package repository

import (
	"Alhanaqtah/habr-backend/pkg/logger"
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
		logger.Err(r.log, op, "failed to get all users", err)
		return nil
	}
	rows.Close()

	var users []user.User

	for rows.Next() {
		var user user.User

		err := rows.Scan(
			&user.Username,
			&user.Rating,
			&user.Karma,
			&user.RegistrationDate,
		)
		if err != nil {
			logger.Err(r.log, op, "failed to sc users", err)
			return nil
		}

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
		logger.Err(r.log, op, "failed to scan user", err)
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

//func (r *repository) GetUserFollowers(username string) *user.User {
//const op = "user.repository.GetUserFollowers"
//
//ctx, cancel := context.WithCancel(context.Background())
//defer cancel()
//
//row := r.storage.QueryRow(ctx, `SELECT followers FROM users WHERE username=$1`, username)
//
//var user user.User
//
//err := row.Scan(
//	&user.Followers,
//)
//if err != nil {
//	r.log.Error("failed to scan user followers", slog.String("err", err.Error()), slog.Attr{
//		Key:   "op",
//		Value: slog.StringValue(op),
//	})
//
//	return nil
//}
//
//return &user
//}

//func (r *repository) GetUserFollowings(username string) *user.User {
//const op = "user.repository.GetUserFollowings"
//
//ctx, cancel := context.WithCancel(context.Background())
//defer cancel()
//
//userID := r.storage.QueryRow(ctx, `SELECT id FROM users WHERE username=$1`, username)
//
//var user user.User
//
//err := userID.Scan(&user.Id)
//if err != nil {
//	logger.Err(r.log, op, "failed to get user followings", err)
//	return nil
//}
//
//r.log.Debug("got userID", slog.Int("userID", user.Id))
//
//rows, err := r.storage.Query(ctx, `SELECT * FROM followings WHERE user_id=$1`, user.Id)
//if err != nil {
//	logger.Err(r.log, op, "failed to get user followings", err)
//	return nil
//}
//
//for rows.Next() {
//	var follower int
//	err = rows.Scan(
//		&follower,
//	)
//	if err != nil {
//		logger.Err(r.log, op, "failed to scan user followings", err)
//		return nil
//	}
//
//	user.Followings = append(user.Followings, follower)
//}
//
//return &user
//}
