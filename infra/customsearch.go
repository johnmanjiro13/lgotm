package infra

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"math/rand"
	"net/http"

	"google.golang.org/api/customsearch/v1"
)

const imageNum = 5

type customSearchRepository struct {
	svc      *customsearch.Service
	engineID string
}

func NewCustomSearchRepository(svc *customsearch.Service, engineID string) *customSearchRepository {
	return &customSearchRepository{
		svc:      svc,
		engineID: engineID,
	}
}

func (r *customSearchRepository) FindImage(ctx context.Context, query string) (io.Reader, error) {
	search := r.svc.Cse.List().Context(ctx).Cx(r.engineID).
		SearchType("image").
		Num(imageNum).
		Q(query).
		Start(1)
	result, err := search.Do()
	if err != nil {
		return nil, fmt.Errorf("failed to search image: %w", err)
	}
	images := result.Items
	if len(images) <= 0 {
		return nil, fmt.Errorf("image was not found.")
	}
	res, err := http.Get(images[rand.Intn(len(images))].Link)
	if err != nil {
		return nil, fmt.Errorf("failed to get image: %w", err)
	}
	defer res.Body.Close()

	buf := new(bytes.Buffer)
	if _, err = buf.ReadFrom(res.Body); err != nil {
		return nil, err
	}
	return buf, nil
}
