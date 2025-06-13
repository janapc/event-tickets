package database

import (
	"context"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/janapc/event-tickets/clients/internal/domain"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

func TestSaveClient(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to open sqlmock: %v", err)
	}
	defer db.Close()
	query := regexp.QuoteMeta("INSERT INTO clients(name, email) VALUES($1,$2) RETURNING *")
	input, _ := domain.NewClient(domain.ClientParams{
		Name:  "Test Client",
		Email: "test@test.com",
	})
	mock.ExpectPrepare(query).WillBeClosed().ExpectQuery().WithArgs(input.Name, input.Email).WillReturnRows(sqlmock.NewRows([]string{
		"id", "name", "email", "created_at",
	}).AddRow(1, input.Name, input.Email, time.Now()))
	repository := NewClientRepository(db)
	ctx := context.Background()
	result, err := repository.Save(ctx, input)
	assert.NoError(t, err)
	assert.NotEmpty(t, result)
	assert.NotEmpty(t, result.ID)
	assert.NotEmpty(t, result.CreatedAt)
	assert.Equal(t, input.Name, result.Name)
	assert.Equal(t, input.Email, result.Email)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("unmet sqlmock expectations: %v", err)
	}
}

func TestGetOneClient(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to open sqlmock: %v", err)
	}
	email := "test@test.com"
	defer db.Close()
	query := regexp.QuoteMeta("SELECT id, name, email, created_at FROM clients WHERE email = $1")
	now := time.Now()
	rows := sqlmock.NewRows([]string{
		"id", "name", "email", "created_at",
	}).AddRow(
		1, "Test Client", email, now,
	)
	mock.ExpectPrepare(query).ExpectQuery().WithArgs(email).WillReturnRows(rows)
	repository := NewClientRepository(db)
	ctx := context.Background()
	result, err := repository.GetByEmail(ctx, email)
	assert.NoError(t, err)
	assert.NotEmpty(t, result)
	assert.NotEmpty(t, result.ID)
	assert.NotEmpty(t, result.CreatedAt)
	assert.Equal(t, "Test Client", result.Name)
	assert.Equal(t, email, result.Email)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("unmet sqlmock expectations: %v", err)
	}
}
