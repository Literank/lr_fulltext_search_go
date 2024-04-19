/*
Package search does all search engine related implementations.
*/
package search

import (
	"context"
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/typedapi/core/search"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"

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

// SearchBooks search from ES and return a list of books
func (s *ElasticSearchEngine) SearchBooks(ctx context.Context, query string) ([]*model.Book, error) {
	resp, err := s.client.Search().Index(INDEX_BOOK).
		Request(&search.Request{
			Query: &types.Query{
				MultiMatch: &types.MultiMatchQuery{
					Query:  query,
					Fields: []string{"title", "author", "content"},
				},
			},
		}).
		Do(ctx)
	if err != nil {
		return nil, err
	}
	books := make([]*model.Book, 0)
	for _, hit := range resp.Hits.Hits {
		var b model.Book
		if err := json.Unmarshal(hit.Source_, &b); err != nil {
			return nil, err
		}
		books = append(books, &b)
	}
	return books, nil
}
