package mocks

import (
	"database/sql"

	"github.com/janapc/event-tickets/events/internal/domain"
	_ "github.com/mattn/go-sqlite3"
)

type DatabaseMockRepository struct {
	DB *sql.DB
}

func NewDatabaseMockRepository() (*DatabaseMockRepository, error) {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		return nil, err
	}
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS events(
		id CHAR(36) NOT NULL,
		name TEXT(150) NOT NULL,
		description TEXT(150) NOT NULL,
		currency TEXT(150) NOT NULL,
		image_url TEXT(150) NOT NULL,
		price DECIMAL(10, 2) NOT NULL,
		event_date TIMESTAMP NOT NULL,
		created_at TIMESTAMP NOT NULL, 
		updated_at TIMESTAMP NOT NULL,
		PRIMARY KEY(id)
		)`)
	if err != nil {
		return nil, err
	}
	return &DatabaseMockRepository{
		DB: db,
	}, nil
}
func (d *DatabaseMockRepository) Register(event *domain.Event) error {
	stmt, err := d.DB.Prepare("INSERT INTO events(id, name, description, image_url, price, currency, event_date, created_at, updated_at) VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(event.ID, event.Name, event.Description, event.ImageUrl, event.Price, event.Currency, event.EventDate, event.CreatedAt, event.UpdatedAt)
	if err != nil {
		return err
	}
	return nil
}
func (d *DatabaseMockRepository) Update(event *domain.Event) error {
	stmt, err := d.DB.Prepare("UPDATE events SET name = $1, description = $2, image_url = $3, price = $4, currency = $5, event_date = $6, updated_at = $7 WHERE id = $8")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(event.Name, event.Description, event.ImageUrl, event.Price, event.Currency, event.EventDate, event.UpdatedAt, event.ID)
	if err != nil {
		return err
	}
	return nil
}

func (d *DatabaseMockRepository) Remove(id string) error {
	stmt, err := d.DB.Prepare("DELETE FROM events where id = $1")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	return nil
}

func (d *DatabaseMockRepository) List() ([]domain.Event, error) {
	rows, err := d.DB.Query("SELECT id, name, description, image_url, price, currency, event_date, created_at, updated_at FROM events")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var events []domain.Event
	for rows.Next() {
		var event domain.Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.ImageUrl, &event.Price, &event.Currency, &event.EventDate, &event.CreatedAt, &event.UpdatedAt)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	return events, nil
}

func (d *DatabaseMockRepository) FindById(id string) (*domain.Event, error) {
	stmt, err := d.DB.Prepare("SELECT id, name, description, image_url, price, currency, event_date, created_at, updated_at FROM events WHERE id = $1")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	var event domain.Event
	err = stmt.QueryRow(id).Scan(&event.ID, &event.Name, &event.Description, &event.ImageUrl, &event.Price, &event.Currency, &event.EventDate, &event.CreatedAt, &event.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &event, nil
}
