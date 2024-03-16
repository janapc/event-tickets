package database

import (
	"database/sql"
	"testing"
	"time"

	"github.com/janapc/event-tickets/events/internal/domain"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

const DDMMYYYY = "02/01/2006"

func connectionDatabase() *sql.DB {
	db, _ := sql.Open("sqlite3", ":memory:")
	_, _ = db.Exec(`CREATE TABLE events(
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
	return db
}

func TestRegisterEvent(t *testing.T) {
	expirateAt := time.Now().Add(48 * time.Hour).Format(DDMMYYYY)
	event, _ := domain.NewEvent("show banana", "description", "http:test.png", 600.40, expirateAt)
	assert.NotEmpty(t, event)
	connection := connectionDatabase()
	assert.NotEmpty(t, connection)
	database := NewPostgresRepository(connection)
	err := database.Register(event)
	assert.NoError(t, err)
	p, err := database.FindById(event.ID)
	assert.NoError(t, err)
	assert.NotEmpty(t, p)
	assert.Equal(t, p.Name, event.Name)
}

func TestUpdateEvent(t *testing.T) {
	expirateAt := time.Now().Add(48 * time.Hour).Format(DDMMYYYY)
	event, _ := domain.NewEvent("show banana", "description", "http:test.png", 600.40, expirateAt)
	assert.NotEmpty(t, event)
	connection := connectionDatabase()
	assert.NotEmpty(t, connection)
	database := NewPostgresRepository(connection)
	_ = database.Register(event)
	event.Name = "teste"
	event.UpdatedAt = time.Now().Add(30 * time.Minute)
	err := database.Update(event)
	assert.NoError(t, err)
	p, err := database.FindById(event.ID)
	assert.NoError(t, err)
	assert.NotEmpty(t, p)
	assert.Equal(t, p.Name, event.Name)
	assert.NotEqual(t, p.UpdatedAt, p.CreatedAt)
}

func TestRemoveEvent(t *testing.T) {
	expirateAt := time.Now().Add(48 * time.Hour).Format(DDMMYYYY)
	event, _ := domain.NewEvent("show banana", "description", "http:test.png", 600.40, expirateAt)
	assert.NotEmpty(t, event)
	connection := connectionDatabase()
	assert.NotEmpty(t, connection)
	database := NewPostgresRepository(connection)
	_ = database.Register(event)
	err := database.Remove(event.ID)
	assert.NoError(t, err)
	_, err = database.FindById(event.ID)
	assert.Error(t, err, "sql: no rows in result set")
}

func TestListEvent(t *testing.T) {
	expirateAt := time.Now().Add(48 * time.Hour).Format(DDMMYYYY)
	event, _ := domain.NewEvent("show banana", "description", "http:test.png", 600.40, expirateAt)
	assert.NotEmpty(t, event)
	connection := connectionDatabase()
	assert.NotEmpty(t, connection)
	database := NewPostgresRepository(connection)
	_ = database.Register(event)
	events, err := database.List()
	assert.NoError(t, err)
	assert.Len(t, events, 1)
	assert.Equal(t, events[0].Name, "show banana")
}

func TestGetOneEvent(t *testing.T) {
	expirateAt := time.Now().Add(48 * time.Hour).Format(DDMMYYYY)
	event, _ := domain.NewEvent("show banana", "description", "http:test.png", 600.40, expirateAt)
	assert.NotEmpty(t, event)
	connection := connectionDatabase()
	assert.NotEmpty(t, connection)
	database := NewPostgresRepository(connection)
	_ = database.Register(event)
	result, err := database.FindById(event.ID)
	assert.NoError(t, err)
	assert.Equal(t, result.ExpirateAt, event.ExpirateAt)
	assert.NotEmpty(t, result)
}
