package domain

import (
	"testing"
)

func TestNewClient_ValidParams(t *testing.T) {
	params := ClientParams{
		Name:  "John Doe",
		Email: "john.doe@example.com",
	}

	client, err := NewClient(params)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if client.Name != params.Name {
		t.Errorf("expected name %v, got %v", params.Name, client.Name)
	}

	if client.Email != params.Email {
		t.Errorf("expected email %v, got %v", params.Email, client.Email)
	}
}

func TestNewClient_MissingName(t *testing.T) {
	params := ClientParams{
		Name:  "",
		Email: "john.doe@example.com",
	}

	_, err := NewClient(params)
	if err == nil {
		t.Fatal("expected an error, got nil")
	}

	expectedErr := "the name field is mandatory"
	if err.Error() != expectedErr {
		t.Errorf("expected error %v, got %v", expectedErr, err.Error())
	}
}

func TestNewClient_MissingEmail(t *testing.T) {
	params := ClientParams{
		Name:  "John Doe",
		Email: "",
	}

	_, err := NewClient(params)
	if err == nil {
		t.Fatal("expected an error, got nil")
	}

	expectedErr := "the email field is mandatory"
	if err.Error() != expectedErr {
		t.Errorf("expected error %v, got %v", expectedErr, err.Error())
	}
}
