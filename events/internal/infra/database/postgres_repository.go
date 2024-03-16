package database

import (
	"database/sql"

	"github.com/janapc/event-tickets/events/internal/domain"
)

type PostgresRepository struct {
	DB *sql.DB
}

func NewPostgresRepository(db *sql.DB) *PostgresRepository {
	return &PostgresRepository{
		DB: db,
	}
}

func (p *PostgresRepository) Register(event *domain.Event) error {
	stmt, err := p.DB.Prepare("INSERT INTO events(id, name, description, image_url, price, expirate_at, created_at, updated_at) VALUES($1,$2,$3,$4,$5,$6,$7,$8)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(event.ID, event.Name, event.Description, event.ImageUrl, event.Price, event.ExpirateAt, event.CreatedAt, event.UpdatedAt)
	if err != nil {
		return err
	}
	return nil
}
func (p *PostgresRepository) Update(event *domain.Event) error {
	stmt, err := p.DB.Prepare("UPDATE events SET name = $1, description = $2, image_url = $3, price = $4, expirate_at = $5, updated_at = $6 WHERE id = $7")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(event.Name, event.Description, event.ImageUrl, event.Price, event.ExpirateAt, event.UpdatedAt, event.ID)
	if err != nil {
		return err
	}
	return nil
}

func (p *PostgresRepository) Remove(id string) error {
	stmt, err := p.DB.Prepare("DELETE FROM events where id = $1")
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

func (p *PostgresRepository) List() ([]domain.Event, error) {
	rows, err := p.DB.Query("SELECT id, name, description, image_url, price, expirate_at, created_at, updated_at FROM events")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var events []domain.Event
	for rows.Next() {
		var event domain.Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.ImageUrl, &event.Price, &event.ExpirateAt, &event.CreatedAt, &event.UpdatedAt)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	return events, nil
}

func (p *PostgresRepository) FindById(id string) (*domain.Event, error) {
	stmt, err := p.DB.Prepare("SELECT id, name, description, image_url, price, expirate_at, created_at, updated_at FROM events WHERE id = $1")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	var event domain.Event
	err = stmt.QueryRow(id).Scan(&event.ID, &event.Name, &event.Description, &event.ImageUrl, &event.Price, &event.ExpirateAt, &event.CreatedAt, &event.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &event, nil
}
