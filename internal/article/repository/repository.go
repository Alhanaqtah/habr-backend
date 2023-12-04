package repository

import (
	lg "Alhanaqtah/habr-backend/pkg/logger/sl"
	"context"
	"database/sql"
	"github.com/lib/pq"
	"log/slog"
	"strings"

	"Alhanaqtah/habr-backend/internal/article"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Repository struct {
	storage *pgxpool.Pool
	log     *slog.Logger
}

func New(pgClient *pgxpool.Pool, log *slog.Logger) *Repository {
	return &Repository{
		storage: pgClient,
		log:     log,
	}
}

func (r *Repository) GetAll() *[]article.Article {
	const op = "article.repository.GetAll"

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	query := `
		SELECT
			articles.id,
			articles.title,
			ARRAY_AGG(users.username) AS Authors,
			flow.flow AS Flow,
			articles.creation_time,
			levels_of_complexity.level_of_complexity AS Level_of_complexity,
			ARRAY_AGG(tags.tag) AS Tags,
			ARRAY_AGG(hubs.hub) AS Hubs,
			articles.rating,
			articles.content
		FROM
			articles
		JOIN
			flow ON articles.flow_id = flow.id
		JOIN
			levels_of_complexity ON articles.level_of_complexity_id = levels_of_complexity.id
		JOIN
			tags ON articles.tag_id = tags.id
		JOIN
			hubs ON articles.hub_id = hubs.id
		LEFT JOIN
			bookmarks ON bookmarks.article_id = articles.id
		LEFT JOIN
			users ON users.id = bookmarks.user_id
		GROUP BY
			articles.id, flow.flow, levels_of_complexity.level_of_complexity;
	`

	rows, err := r.storage.Query(ctx, query)
	if err != nil {
		lg.Err(r.log, op, "failed to get all articles", err)
		return nil
	}

	var articles []article.Article

	for rows.Next() {
		var article article.Article

		var authors []string
		var tags sql.NullString
		var hubs sql.NullString

		err = rows.Scan(
			&article.Id,
			&article.Title,
			pq.Array(&authors),
			&article.Flow,
			&article.CreationTime,
			&article.LevelOfComplexity,
			&tags,
			&hubs,
			&article.Rating,
			&article.Content,
		)
		if err != nil {
			lg.Err(r.log, op, "failed to get all articles", err)
			return nil
		}

		if tags.Valid {
			article.Tags = strings.Split(tags.String, ",")
		}

		if hubs.Valid {
			article.Hubs = strings.Split(hubs.String, ",")
		}

		article.Authors = authors
		articles = append(articles, article)
	}

	return &articles
}

func (r *Repository) GetFlow(flow string) *[]article.Article {
	const op = "article.repository.GetFlow"

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	query := `
		SELECT
			articles.id,
			articles.title,
			ARRAY_AGG(users.username) AS Authors,
			flow.flow AS Flow,
			articles.creation_time,
			levels_of_complexity.level_of_complexity AS Level_of_complexity,
			ARRAY_AGG(tags.tag) AS Tags,
			ARRAY_AGG(hubs.hub) AS Hubs,
			articles.rating,
			articles.content
		FROM
			articles
		JOIN
			flow ON articles.flow_id = flow.id
		JOIN
			levels_of_complexity ON articles.level_of_complexity_id = levels_of_complexity.id
		JOIN
			tags ON articles.tag_id = tags.id
		JOIN
			hubs ON articles.hub_id = hubs.id
		LEFT JOIN
			bookmarks ON bookmarks.article_id = articles.id
		LEFT JOIN
			users ON users.id = bookmarks.user_id
		WHERE flow.flow = $1
		GROUP BY
			articles.id, flow.flow, levels_of_complexity.level_of_complexity;
	`

	rows, err := r.storage.Query(ctx, query, flow)
	if err != nil {
		lg.Err(r.log, op, "failed to get flow of articles", err)
		return nil
	}

	var articles []article.Article

	for rows.Next() {
		var article article.Article

		var authors []string
		var tags sql.NullString
		var hubs sql.NullString

		err = rows.Scan(
			&article.Id,
			&article.Title,
			pq.Array(&authors),
			&article.Flow,
			&article.CreationTime,
			&article.LevelOfComplexity,
			&tags,
			&hubs,
			&article.Rating,
			&article.Content,
		)
		if err != nil {
			lg.Err(r.log, op, "failed to get all articles", err)
			return nil
		}

		if tags.Valid {
			article.Tags = strings.Split(tags.String, ",")
		}

		if hubs.Valid {
			article.Hubs = strings.Split(hubs.String, ",")
		}

		article.Authors = authors
		articles = append(articles, article)
	}

	return &articles
}

func (r *Repository) GetByID(id int) *article.Article {
	const op = "article.repository.GetByID"

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	query := `
		SELECT
			articles.id,
			articles.title,
			ARRAY_AGG(users.username) AS Authors,
			flow.flow AS Flow,
			articles.creation_time,
			levels_of_complexity.level_of_complexity AS Level_of_complexity,
			ARRAY_AGG(tags.tag) AS Tags,
			ARRAY_AGG(hubs.hub) AS Hubs,
			articles.rating,
			articles.content
		FROM
			articles
		JOIN
			flow ON articles.flow_id = flow.id
		JOIN
			levels_of_complexity ON articles.level_of_complexity_id = levels_of_complexity.id
		JOIN
			tags ON articles.tag_id = tags.id
		JOIN
			hubs ON articles.hub_id = hubs.id
		LEFT JOIN
			bookmarks ON bookmarks.article_id = articles.id
		LEFT JOIN
			users ON users.id = bookmarks.user_id
		WHERE articles.id = $1
		GROUP BY
			articles.id, flow.flow, levels_of_complexity.level_of_complexity;
	`

	row := r.storage.QueryRow(ctx, query, id)

	var article article.Article

	var authors []string
	var tags sql.NullString
	var hubs sql.NullString

	err := row.Scan(
		&article.Id,
		&article.Title,
		pq.Array(&authors),
		&article.Flow,
		&article.CreationTime,
		&article.LevelOfComplexity,
		&tags,
		&hubs,
		&article.Rating,
		&article.Content,
	)
	if err != nil {
		lg.Err(r.log, op, "failed to get all articles", err)
		return nil
	}

	if tags.Valid {
		article.Tags = strings.Split(tags.String, ",")
	}

	if hubs.Valid {
		article.Hubs = strings.Split(hubs.String, ",")
	}

	article.Authors = authors

	return &article
}
