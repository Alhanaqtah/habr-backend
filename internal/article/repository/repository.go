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

	rows, err := r.storage.Query(ctx,
		`SELECT
    			articles.id,
    			articles.title,
    			(
        			SELECT
						ARRAY_AGG(username)
        			FROM users
        				JOIN bookmarks ON bookmarks.user_id = users.id
					WHERE bookmarks.article_id = articles.id
    			) AS Authors,
				flow.flow AS Flow,
				articles.creation_time,
				levels_of_complexity.level_of_complexity AS Level_of_complexity,
				tags.tag AS Tags,
				hubs.hub AS Hubs,
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
    			hubs ON articles.hub_id = hubs.id;
	`)
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

func (r *repository) GetByID(id int) *article.Article {
	const op = "article.repository.GetByID"

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	row := r.storage.QueryRow(ctx,
		`SELECT
			articles.id,
			articles.title,
			(
				SELECT
					ARRAY_AGG(username)
				FROM users
					JOIN bookmarks ON bookmarks.user_id = users.id
				WHERE bookmarks.article_id = articles.id
			) AS Authors,
			flow.flow AS Flow,
			articles.creation_time,
			levels_of_complexity.level_of_complexity AS Level_of_complexity,
			tags.tag AS Tags,
			hubs.hub AS Hubs,
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
		WHERE articles.id = $1;`,
		id,
	)

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

func (r *repository) GetFlow(flow string) *[]article.Article {
	const op = "article.repository.GetFlow"

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	rows, err := r.storage.Query(ctx,
		`SELECT
			articles.id,
			articles.title,
			(
				SELECT
					ARRAY_AGG(username)
				FROM users
					JOIN bookmarks ON bookmarks.user_id = users.id
				WHERE bookmarks.article_id = articles.id
			) AS Authors,
			flow.flow AS Flow,
			articles.creation_time,
			levels_of_complexity.level_of_complexity AS Level_of_complexity,
			tags.tag AS Tag,
			hubs.hub AS Hub,
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
		WHERE flow = $1;`,
		flow,
	)
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
