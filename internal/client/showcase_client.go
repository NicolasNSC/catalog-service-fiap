package client

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/NicolasNSC/catalog-service-fiap/internal/dto"
)

//go:generate mockgen -source=showcase_client.go -destination=./mocks/showcase_client_mock.go -package=mocks
type ShowcaseClientInterface interface {
	CreateListing(ctx context.Context, data dto.CreateListingDTO) error
}

type httpShowcaseClient struct {
	client  *http.Client
	baseURL string
}

func NewShowcaseClient(baseURL string) ShowcaseClientInterface {
	return &httpShowcaseClient{
		client:  &http.Client{},
		baseURL: baseURL,
	}
}

func (c *httpShowcaseClient) CreateListing(ctx context.Context, data dto.CreateListingDTO) error {
	payload, err := json.Marshal(data)
	if err != nil {
		return err
	}

	url := c.baseURL + "/listings"
	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(payload))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return errors.New("showcase service returned non-success status: " + resp.Status)
	}

	return nil
}
