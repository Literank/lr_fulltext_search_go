/*
Package search does all search engine related implementations.
*/
package search

import (
	"context"

	"github.com/elastic/go-elasticsearch/v8"

	"literank.com/fulltext-books/domain/model"
)

const INDEX_BOOK = "book_idx"

// ElasticSearchEngine runs all index/search operations
type ElasticSearchEngine struct {
	client   *elasticsearch.TypedClient
	pageSize int
}

// NewEngine constructs a new ElasticSearchEngine
func NewEngine(address string, pageSize int) (*ElasticSearchEngine, error) {
	cfg := elasticsearch.Config{
		Addresses: []string{address},
	}
	client, err := elasticsearch.NewTypedClient(cfg)
	if err != nil {
		return nil, err
	}
	// Create the index
	return &ElasticSearchEngine{client, pageSize}, nil
}

// IndexBook indexes a new book
func (s *ElasticSearchEngine) IndexBook(ctx context.Context, b *model.Book) (string, error) {
	resp, err := s.client.Index(INDEX_BOOK).
		Request(b).
		Do(ctx)
	if err != nil {
		return "", err
	}
	return resp.Id_, nil
}
