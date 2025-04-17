package database

import (
	"database/sql"
	"testing"

	"github.com/janapc/event-tickets/clients/internal/domain"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

func connectionDatabase() *sql.DB {
	db, _ := sql.Open("sqlite3", ":memory:")
	_, _ = db.Exec(`CREATE TABLE clients(
		id CHAR(36) NOT NULL,
		name TEXT(150) NOT NULL,
		email TEXT(150) NOT NULL,
		created_at TIMESTAMP NOT NULL, 
		PRIMARY KEY(id)
		)`)
	return db
}

func TestSaveAClient(t *testing.T) {
	conn := connectionDatabase()
	repository := NewClientRepository(conn)
	c, _ := domain.NewClient("test", "test@test.com")
	err := repository.Save(c)
	assert.NoError(t, err)
}

func TestGetAClientByEmail(t *testing.T) {
	conn := connectionDatabase()
	repository := NewClientRepository(conn)
	c, _ := domain.NewClient("test", "test@test.com")
	_ = repository.Save(c)
	client, err := repository.GetByEmail(c.Email)
	assert.NoError(t, err)
	assert.NotEmpty(t, client)
	assert.Equal(t, client.ID, c.ID)
}
