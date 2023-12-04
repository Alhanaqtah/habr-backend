package article

import (
	lg "Alhanaqtah/habr-backend/pkg/logger/sl"
	"log/slog"
	"net/http"
	"slices"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

var flows = []string{
	"develop",
	"admin",
	"design",
	"management",
	"marketing",
	"popsci",
}

type Handler struct {
	storage storage
	log     *slog.Logger
}

func NewHandler(storage storage, log *slog.Logger) *Handler {
	return &Handler{
		storage: storage,
		log:     log,
	}
}

func (h *Handler) Register(r *chi.Mux) {
	r.Route("/articles", func(r chi.Router) {
		r.Get("/", h.getAllArticles)
		r.Get("/{id}", h.getArticleByID)
		r.Get("/flow/{flow}", h.getArticlesByFlow)
	})
}

func (h *Handler) getAllArticles(w http.ResponseWriter, r *http.Request) {
	//const op = "article.handler.getAllArticles"

	// Запрос к бд
	articles := h.storage.GetAll()

	// Формирование ответа
	render.JSON(w, r, articles)
}

func (h *Handler) getArticleByID(w http.ResponseWriter, r *http.Request) {
	const op = "article.handler.getArticlesByID"

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		lg.Err(h.log, op, "failed to get flow {id} param", err)

		render.JSON(w, r, "Статья не найдена")

		return
	}

	// Запрос к бд
	article := h.storage.GetByID(id)

	if article == nil {
		h.log.Debug("article not found", slog.Attr{
			Key:   "op",
			Value: slog.StringValue(op),
		})

		render.JSON(w, r, "Статья не найдена")

		return
	}

	// Формирование ответа
	render.JSON(w, r, article)
}

func (h *Handler) getArticlesByFlow(w http.ResponseWriter, r *http.Request) {
	const op = "article.handler.getArticlesByFlow"

	flow := chi.URLParam(r, "flow")

	slog.Debug("flow: ", flow, slog.Attr{
		Key:   "op",
		Value: slog.StringValue(op),
	})

	// Валидация
	if !slices.Contains(flows, flow) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Такого потока не существует"))
		return
	}

	// Запрос к бд
	articles := h.storage.GetFlow(flow)

	// Формирование ответа
	render.JSON(w, r, articles)
}
