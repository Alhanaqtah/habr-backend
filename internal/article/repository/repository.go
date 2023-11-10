package repository

import (
	"context"
	"log/slog"

	"Alhanaqtah/habr-backend/internal/article"
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

func (r *repository) GetAll() *[]article.Article {
	const op = "article.repository.GetAll"

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	rows, err := r.storage.Query(ctx, `SELECT id, title, author, flow, creation_time, level_of_complexity, time_to_read, views, tags, hubs, rating, content FROM articles`)
	if err != nil {
		r.log.Error("failed to get all articles: ", slog.String("err", err.Error()), slog.Attr{
			Key:   "op",
			Value: slog.StringValue(op),
		})
		return nil
	}

	var articles []article.Article

	for rows.Next() {
		var article article.Article

		err = rows.Scan(
			&article.Id,
			&article.Title,
			&article.Author,
			&article.Flow,
			&article.CreationTime,
			&article.LevelOfComplexity,
			&article.TimeToRead,
			&article.Views,
			&article.Tags,
			&article.Hubs,
			&article.Rating,
			&article.Content,
		)
		if err != nil {
			r.log.Error("failed to get all articles", slog.String("err", err.Error()), slog.Attr{
				Key:   "op",
				Value: slog.StringValue(op),
			})
			return nil
		}

		articles = append(articles, article)
	}

	return &articles
}

func (r *repository) GetByID(id int) *article.Article {
	const op = "article.repository.GetByID"

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	row := r.storage.QueryRow(ctx, `SELECT id, title, author, flow, creation_time, level_of_complexity, time_to_read, views, tags, hubs, rating, content FROM articles WHERE id=$1`, id)

	var article article.Article

	err := row.Scan(
		&article.Id,
		&article.Title,
		&article.Author,
		&article.Flow,
		&article.CreationTime,
		&article.LevelOfComplexity,
		&article.TimeToRead,
		&article.Views,
		&article.Tags,
		&article.Hubs,
		&article.Rating,
		&article.Content,
	)
	if err != nil {
		r.log.Error("failed to get flow of articles", slog.String("err", err.Error()), slog.Attr{
			Key:   "op",
			Value: slog.StringValue(op),
		})
		return nil
	}

	return &article
}

func (r *repository) GetFlow(flow string) *[]article.Article {
	const op = "article.repository.GetFlow"

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	rows, err := r.storage.Query(ctx, `SELECT id, title, author, flow, creation_time, level_of_complexity, time_to_read, views, tags, hubs, rating, content FROM articles WHERE flow=$1`, flow)
	if err != nil {
		r.log.Error("failed to get flow of articles: ", slog.String("err", err.Error()), slog.Attr{
			Key:   "op",
			Value: slog.StringValue(op),
		})
		return nil
	}

	var articles []article.Article

	for rows.Next() {
		var article article.Article

		err = rows.Scan(
			&article.Id,
			&article.Title,
			&article.Author,
			&article.Flow,
			&article.CreationTime,
			&article.LevelOfComplexity,
			&article.TimeToRead,
			&article.Views,
			&article.Tags,
			&article.Hubs,
			&article.Rating,
			&article.Content,
		)
		if err != nil {
			r.log.Error("failed to get flow of articles", slog.String("err", err.Error()), slog.Attr{
				Key:   "op",
				Value: slog.StringValue(op),
			})
			return nil
		}

		articles = append(articles, article)
	}

	return &articles
}
