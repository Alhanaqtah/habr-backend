package article

import (
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

type handler struct {
	storage storage
	log     *slog.Logger
}

func NewHandler(storage storage, log *slog.Logger) *handler {
	return &handler{
		storage: storage,
		log:     log,
	}
}

func (h *handler) Register(r *chi.Mux) {
	r.Route("/articles", func(r chi.Router) {
		r.Get("/", h.getAllArticles)
		r.Get("/{id}", h.getArticleByID)
		r.Get("/flow/{flow}", h.getArticlesByFlow)
	})
}

func (h *handler) getAllArticles(w http.ResponseWriter, r *http.Request) {
	//const op = "article.handler.getAllArticles"

	// Запрос к бд
	articles := h.storage.GetAll()

	// Формирование ответа
	render.JSON(w, r, articles)
}

func (h *handler) getArticleByID(w http.ResponseWriter, r *http.Request) {
	const op = "article.handler.getArticlesByID"

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		h.log.Error("failed to get flow {id} param", slog.String("err", err.Error()), slog.Attr{
			Key:   "op",
			Value: slog.StringValue(op),
		})

		render.JSON(w, r, "Статья не найдена")

		return
	}

	// Запрос к бд
	article := h.storage.GetByID(id)

	if article == nil {
		h.log.Debug("article not found", slog.String("err", err.Error()), slog.Attr{
			Key:   "op",
			Value: slog.StringValue(op),
		})

		render.JSON(w, r, "Статья не найдена")

		return
	}

	// Формирование ответа
	render.JSON(w, r, article)
}

func (h *handler) getArticlesByFlow(w http.ResponseWriter, r *http.Request) {
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
