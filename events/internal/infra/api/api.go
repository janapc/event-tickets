package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/janapc/event-tickets/events/internal/application"
	"github.com/janapc/event-tickets/events/internal/domain"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

type Api struct {
	Repository domain.EventRepository
}

func NewApi(repo domain.EventRepository) *Api {
	return &Api{
		Repository: repo,
	}
}

func (a *Api) Init(port string) {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))
	r.Use(Authorization)
	r.Route("/events", func(r chi.Router) {
		baseUrl := fmt.Sprintf("%s/events/docs/doc.json", "http://localhost"+port)
		r.Get("/docs/*", httpSwagger.Handler(httpSwagger.URL(baseUrl)))
		r.Get("/", a.GetEvents)
		r.Get("/{eventId}", a.GetEventById)
		r.Mount("/admin", a.adminRouter())
	})
	log.Printf("Server running in port %s", port)
	http.ListenAndServe(port, r)
}

func (a *Api) adminRouter() http.Handler {
	r := chi.NewRouter()
	r.Use(AdminOnly)
	r.Post("/", a.RegisterEvent)
	r.Put("/{eventId}", a.UpdateEvent)
	r.Delete("/{eventId}", a.RemoveEvent)
	return r
}

// GetEvents godoc
// @Description list events
// @Accept json
// @Produce json
// @Success 200 {array} application.OutputGetEventsDTO
// @Failure 500
// @Router / [get]
// @Security BearerAuth
func (a *Api) GetEvents(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	app := application.NewGetEvents(a.Repository)
	result, err := app.Execute()
	if err != nil {
		message, statusCode := HandlerErrors(err)
		w.WriteHeader(statusCode)
		w.Write(message)
		return
	}
	json, _ := json.Marshal(result)
	w.Write(json)
}

// GetEventById godoc
// @Description get a event by id
// @Accept json
// @Produce json
// @Param id path string true "event id" Format(uuid)
// @Success 200 {object} domain.Event
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /{id} [get]
// @Security BearerAuth
func (a *Api) GetEventById(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	id := chi.URLParam(r, "eventId")
	app := application.NewGetEventById(a.Repository)
	result, err := app.Execute(id)
	if err != nil {
		message, statusCode := HandlerErrors(err)
		w.WriteHeader(statusCode)
		w.Write(message)
		return
	}
	json, _ := json.Marshal(result)
	w.Write(json)
}

// RemoveEvent godoc
// @Description remove a event by id
// @Accept json
// @Produce json
// @Param id path string true "event id" Format(uuid)
// @Success 204
// @Failure 404
// @Failure 500
// @Router /admin/{id} [delete]
// @Security BearerAuth
func (a *Api) RemoveEvent(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	id := chi.URLParam(r, "eventId")
	app := application.NewRemoveEvent(a.Repository)
	err := app.Execute(id)
	if err != nil {
		message, statusCode := HandlerErrors(err)
		w.WriteHeader(statusCode)
		w.Write(message)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// RegisterEvent godoc
// @Description register a new event
// @Accept json
// @Produce json
// @Param request body application.InputRegisterEventDTO true "event request"
// @Success 201
// @Failure 400
// @Failure 500
// @Router /admin/ [post]
// @Security BearerAuth
func (a *Api) RegisterEvent(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	var input application.InputRegisterEventDTO
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, "Body is invalid", http.StatusBadRequest)
		return
	}
	app := application.NewRegisterEvent(a.Repository)
	result, err := app.Execute(input)
	if err != nil {
		message, statusCode := HandlerErrors(err)
		w.WriteHeader(statusCode)
		w.Write(message)
		return
	}
	json, _ := json.Marshal(result)
	w.Write(json)
}

// UpdateEvent godoc
// @Description update a event
// @Accept json
// @Produce json
// @Param id path string true "event id" Format(uuid)
// @Param request body application.InputUpdateEventDTO true "event request"
// @Success 204
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /admin/{id} [put]
// @Security BearerAuth
func (a *Api) UpdateEvent(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	var input application.InputUpdateEventDTO
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, "Body is invalid", http.StatusBadRequest)
		return
	}
	input.ID = chi.URLParam(r, "eventId")
	app := application.NewUpdateEvent(a.Repository)
	err = app.Execute(input)
	if err != nil {
		message, statusCode := HandlerErrors(err)
		w.WriteHeader(statusCode)
		w.Write(message)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
