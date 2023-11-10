package user

import (
	"log/slog"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

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
	r.Route("/users", func(r chi.Router) {
		r.Get("/", h.getAllUsers)
		//r.Get("/{username}", h.getUser)
		//r.Get("/{username}/publications", h.getUserPublications)
		//r.Get("/{username}/followers", h.getUserFollowers)
		//r.Get("/{username}/followings", h.getUserFollowings)
	})
}

func (h *handler) getAllUsers(w http.ResponseWriter, r *http.Request) {
	//const op = "user.handler.getAllUsers"

	users := h.storage.GetUsers()

	render.JSON(w, r, users)

}

func (h *handler) getUser(w http.ResponseWriter, r *http.Request) {
	const op = "user.handler.getUser"
}

func (h *handler) getUserPublications(w http.ResponseWriter, r *http.Request) {
	const op = "user.handler.getUserPublications"
}

func (h *handler) getUserFollowers(w http.ResponseWriter, r *http.Request) {
	const op = "user.handler.getUserFollowers"

}

func (h *handler) getUserFollowings(w http.ResponseWriter, r *http.Request) {
	const op = "user.handler.getUserFollowings"
}
