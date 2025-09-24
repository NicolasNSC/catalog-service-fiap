package client_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/NicolasNSC/catalog-service-fiap/internal/client"
	"github.com/NicolasNSC/catalog-service-fiap/internal/dto"
)

func TestCreateListing_Success(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/listings" {
			t.Errorf("expected path /listings, got %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusCreated)
	}))
	defer server.Close()
	showcaseClient := client.NewShowcaseClient(server.URL)

	data := dto.CreateListingDTO{
		VehicleID: "123",
		Brand:     "Toyota",
		Model:     "Corolla",
		Price:     100000,
	}

	err := showcaseClient.CreateListing(context.Background(), data)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}

func TestCreateListing_NonSuccessStatus(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
	}))
	defer server.Close()

	showcaseClient := client.NewShowcaseClient(server.URL)

	data := dto.CreateListingDTO{}

	err := showcaseClient.CreateListing(context.Background(), data)
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

func TestCreateListing_HTTPError(t *testing.T) {
	showcaseClient := client.NewShowcaseClient("http://invalid-host")

	data := dto.CreateListingDTO{}

	err := showcaseClient.CreateListing(context.Background(), data)
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}
