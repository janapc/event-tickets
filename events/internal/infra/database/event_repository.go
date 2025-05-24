package database

import (
	"context"
	"database/sql"

	"github.com/janapc/event-tickets/events/internal/domain"
)

type EventRepository struct {
	DB *sql.DB
}

func NewEventRepository(db *sql.DB) *EventRepository {
	return &EventRepository{
		DB: db,
	}
}

func (e *EventRepository) Register(ctx context.Context, event *domain.Event) (*domain.Event, error) {
	stmt, err := e.DB.PrepareContext(ctx, "INSERT INTO events(name, description, image_url, price, currency, event_date) VALUES($1,$2,$3,$4,$5,$6) RETURNING *")
	if err != nil {
		return &domain.Event{}, err
	}
	defer stmt.Close()
	var newEvent domain.Event
	err = stmt.QueryRowContext(ctx, event.Name, event.Description, event.ImageUrl, event.Price, event.Currency, event.EventDate).Scan(&newEvent.ID, &newEvent.Name, &newEvent.Description, &newEvent.ImageUrl, &newEvent.Price, &newEvent.Currency, &newEvent.EventDate, &newEvent.CreatedAt, &newEvent.UpdatedAt)
	if err != nil {
		return &domain.Event{}, err
	}

	return &newEvent, nil
}
func (e *EventRepository) Update(ctx context.Context, event *domain.Event) error {
	stmt, err := e.DB.PrepareContext(ctx, "UPDATE events SET name = $1, description = $2, image_url = $3, price = $4, currency = $5, event_date = $6, updated_at = $7 WHERE id = $8")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.ExecContext(ctx, event.Name, event.Description, event.ImageUrl, event.Price, event.Currency, event.EventDate, event.UpdatedAt, event.ID)
	if err != nil {
		return err
	}
	return nil
}

func (e *EventRepository) Remove(ctx context.Context, id int64) error {
	stmt, err := e.DB.PrepareContext(ctx, "DELETE FROM events where id = $1")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.ExecContext(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (e *EventRepository) List(ctx context.Context) ([]*domain.Event, error) {
	rows, err := e.DB.QueryContext(ctx, "SELECT id, name, description, image_url, price, currency, event_date, created_at, updated_at FROM events")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var events []*domain.Event
	for rows.Next() {
		event := &domain.Event{}
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.ImageUrl, &event.Price, &event.Currency, &event.EventDate, &event.CreatedAt, &event.UpdatedAt)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	return events, nil
}

func (e *EventRepository) FindByID(ctx context.Context, id int64) (*domain.Event, error) {
	stmt, err := e.DB.PrepareContext(ctx, "SELECT id, name, description, image_url, price, currency, event_date, created_at, updated_at FROM events WHERE id = $1")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	var event domain.Event
	err = stmt.QueryRowContext(ctx, id).Scan(&event.ID, &event.Name, &event.Description, &event.ImageUrl, &event.Price, &event.Currency, &event.EventDate, &event.CreatedAt, &event.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &event, nil
}
