package database

import (
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/janapc/event-tickets/events/internal/domain"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

func TestRegisterEvent(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to open sqlmock: %v", err)
	}
	defer db.Close()
	query := regexp.QuoteMeta("INSERT INTO events(name, description, image_url, price, currency, event_date) VALUES($1,$2,$3,$4,$5,$6) RETURNING *")
	eventDate := time.Now().Add(48 * time.Hour).Format(time.RFC3339)
	params := domain.EventParams{
		Name:        "GoConf",
		Description: "Go conference",
		ImageUrl:    "http://image",
		Price:       99.99,
		EventDate:   eventDate,
		Currency:    "USD",
	}
	event, _ := domain.NewEvent(params)
	mock.ExpectPrepare(query).WillBeClosed().ExpectQuery().WithArgs(event.Name, event.Description, event.ImageUrl, event.Price, event.Currency, event.EventDate).WillReturnRows(sqlmock.NewRows([]string{
		"id", "name", "description", "image_url", "price", "currency", "event_date", "created_at", "update_at",
	}).AddRow(1, event.Name, event.Description, event.ImageUrl, event.Price, event.Currency, event.EventDate, time.Now(), time.Now()))

	assert.NotEmpty(t, event)
	repository := NewEventRepository(db)
	result, err := repository.Register(event)
	assert.NoError(t, err)
	assert.NotEmpty(t, result)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("unmet sqlmock expectations: %v", err)
	}
}

func TestUpdateEvent(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to open sqlmock: %v", err)
	}
	defer db.Close()
	query := regexp.QuoteMeta("UPDATE events SET name = $1, description = $2, image_url = $3, price = $4, currency = $5, event_date = $6, updated_at = $7 WHERE id = $8")
	eventDate := time.Now().Add(48 * time.Hour)
	event := &domain.Event{
		ID:          int64(1),
		Name:        "GoConf",
		Description: "Go conference",
		ImageUrl:    "http://image",
		Price:       99.99,
		EventDate:   eventDate,
		Currency:    "USD",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	mock.ExpectPrepare(query).ExpectExec().WithArgs(event.Name, event.Description, event.ImageUrl, event.Price, event.Currency, event.EventDate, event.UpdatedAt, event.ID).WillReturnResult((sqlmock.NewResult(0, 1)))
	repository := NewEventRepository(db)
	err = repository.Update(event)
	assert.NoError(t, err)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("unmet sqlmock expectations: %v", err)
	}
}

func TestRemoveEvent(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to open sqlmock: %v", err)
	}
	defer db.Close()
	query := regexp.QuoteMeta("DELETE FROM events where id = $1")
	id := int64(1)
	mock.ExpectPrepare(query).WillBeClosed().ExpectExec().WithArgs(id).WillReturnResult(sqlmock.NewResult(0, 1))
	repository := NewEventRepository(db)
	err = repository.Remove(id)
	assert.NoError(t, err)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("unmet sqlmock expectations: %v", err)
	}
}

func TestListEvent(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to open sqlmock: %v", err)
	}
	defer db.Close()
	query := regexp.QuoteMeta("SELECT id, name, description, image_url, price, currency, event_date, created_at, updated_at FROM events")
	now := time.Now()
	rows := sqlmock.NewRows([]string{
		"id", "name", "description", "image_url", "price", "currency", "event_date", "created_at", "updated_at",
	}).AddRow(
		1, "GoConf", "Go conference", "http://image", 99.99, "USD", now, now, now,
	)
	mock.ExpectQuery(query).WillReturnRows(rows)
	repository := NewEventRepository(db)
	events, err := repository.List()
	assert.NoError(t, err)
	assert.Len(t, events, 1)
	assert.Equal(t, events[0].Name, "GoConf")
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("unmet sqlmock expectations: %v", err)
	}
}

func TestGetOneEvent(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to open sqlmock: %v", err)
	}
	defer db.Close()
	id := int64(1)
	query := regexp.QuoteMeta("SELECT id, name, description, image_url, price, currency, event_date, created_at, updated_at FROM events WHERE id = $1")
	now := time.Now()
	rows := sqlmock.NewRows([]string{
		"id", "name", "description", "image_url", "price", "currency", "event_date", "created_at", "updated_at",
	}).AddRow(
		1, "GoConf", "Go conference", "http://image", 99.99, "USD", now, now, now,
	)
	mock.ExpectPrepare(query).ExpectQuery().WithArgs(id).WillReturnRows(rows)
	repository := NewEventRepository(db)
	result, err := repository.FindByID(id)
	assert.NoError(t, err)
	assert.NotEmpty(t, result)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("unmet sqlmock expectations: %v", err)
	}
}
