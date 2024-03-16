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
		image_url TEXT(150) NOT NULL,
		price DECIMAL(10, 2) NOT NULL,
		expirate_at TIMESTAMP NOT NULL,
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

func (p *DatabaseMockRepository) Register(event *domain.Event) error {
	stmt, err := p.DB.Prepare("INSERT INTO events(id, name, description, image_url, price, expirate_at, created_at, updated_at) VALUES(?,?,?,?,?,?,?,?)")
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
func (p *DatabaseMockRepository) Update(event *domain.Event) error {
	stmt, err := p.DB.Prepare("UPDATE events SET name = ?, description = ?, image_url = ?, price = ?, expirate_at = ?, updated_at = ? WHERE id = ?")
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

func (p *DatabaseMockRepository) Remove(id string) error {
	stmt, err := p.DB.Prepare("DELETE FROM events where id = ?")
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

func (p *DatabaseMockRepository) List() ([]domain.Event, error) {
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

func (p *DatabaseMockRepository) FindById(id string) (*domain.Event, error) {
	stmt, err := p.DB.Prepare("SELECT id, name, description, image_url, price, expirate_at, created_at, updated_at FROM events WHERE id = ?")
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
