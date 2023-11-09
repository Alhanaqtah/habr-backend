package article

import (
	"log/slog"
	"net/http"
	"slices"

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
		r.Get("/{flow}", h.getArticlesByFlow)
	})
}

func (h *handler) getAllArticles(w http.ResponseWriter, r *http.Request) {
	//const op = "article.handler.getAllArticles"

	// Запрос к бд
	articles := h.storage.GetAll()

	// Формирование ответа
	render.JSON(w, r, articles)
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
