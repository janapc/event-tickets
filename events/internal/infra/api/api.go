package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/jwtauth/v5"
	"github.com/janapc/event-tickets/events/internal/application"
	"github.com/janapc/event-tickets/events/internal/domain"
	"github.com/riandyrn/otelchi"

	httpSwagger "github.com/swaggo/http-swagger/v2"
)

const ROUTE_PREFIX = "/events"

type Api struct {
	Repository domain.IEventRepository
}

func NewApi(repo domain.IEventRepository) *Api {
	return &Api{
		Repository: repo,
	}
}

var tokenAuth *jwtauth.JWTAuth

func (a *Api) Init(port string) {
	r := chi.NewRouter()
	jwtSecret := []byte(os.Getenv("JWT_SECRET"))
	tokenAuth = jwtauth.New("HS256", jwtSecret, nil)
	r.Use(otelchi.Middleware("go-otel-postgres", otelchi.WithChiRoutes(r)))
	r.Use(middleware.RequestID)
	r.Use(middleware.Heartbeat("/healthcheck"))
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))
	baseUrlDocs := fmt.Sprintf("%s/docs/doc.json", os.Getenv("BASE_API_URL"))
	r.Get("/docs/*", httpSwagger.Handler(httpSwagger.URL(baseUrlDocs)))
	r.Route(ROUTE_PREFIX, func(r chi.Router) {
		r.Use(jwtauth.Verifier(tokenAuth))
		r.Use(jwtauth.Authenticator(tokenAuth))
		r.Use(WithJWTAuth(tokenAuth))

		r.Get("/", a.GetEvents)
		r.Get("/{eventId}", a.GetEventById)
		r.Mount("/admin", a.adminRouter())
	})
	log.Printf("Server running in port %s", port)
	if err := http.ListenAndServe(port, r); err != nil {
		log.Panicln(err)
	}
}

func (a *Api) adminRouter() http.Handler {
	r := chi.NewRouter()
	r.Use(OnlyAdmin)
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
	ctx := r.Context()
	app := application.NewGetEvents(a.Repository)
	events, err := app.Execute(ctx)
	if err != nil {
		message, statusCode := HandlerErrors(err)
		w.WriteHeader(statusCode)
		if _, err := w.Write(message); err != nil {
			log.Panicln(err)
		}
		return
	}
	response, _ := json.Marshal(events)
	if _, err := w.Write(response); err != nil {
		log.Panicln(err)
	}
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
	app := application.NewGetEventById(a.Repository)
	id, _ := strconv.Atoi(chi.URLParam(r, "eventId"))
	ctx := r.Context()
	event, err := app.Execute(ctx, int64(id))
	if err != nil {
		message, statusCode := HandlerErrors(err)
		w.WriteHeader(statusCode)
		if _, err := w.Write(message); err != nil {
			log.Panicln(err)
		}
		return
	}
	response, _ := json.Marshal(event)
	if _, err := w.Write(response); err != nil {
		log.Panicln(err)
	}
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
	app := application.NewRemoveEvent(a.Repository)
	id, _ := strconv.Atoi(chi.URLParam(r, "eventId"))
	ctx := r.Context()
	err := app.Execute(ctx, int64(id))
	if err != nil {
		message, statusCode := HandlerErrors(err)
		w.WriteHeader(statusCode)
		if _, err := w.Write(message); err != nil {
			log.Panicln(err)
		}
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
	ctx := r.Context()
	event, err := app.Execute(ctx, input)
	if err != nil {
		message, statusCode := HandlerErrors(err)
		w.WriteHeader(statusCode)
		if _, err := w.Write(message); err != nil {
			log.Panicln(err)
		}
		return
	}
	response, _ := json.Marshal(event)
	if _, err := w.Write(response); err != nil {
		log.Panicln(err)
	}
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
	ctx := r.Context()
	id, _ := strconv.Atoi(chi.URLParam(r, "eventId"))
	app := application.NewUpdateEvent(a.Repository)
	err = app.Execute(ctx, int64(id), input)
	if err != nil {
		message, statusCode := HandlerErrors(err)
		w.WriteHeader(statusCode)
		if _, err := w.Write(message); err != nil {
			log.Panicln(err)
		}
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
